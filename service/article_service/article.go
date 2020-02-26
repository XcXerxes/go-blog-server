/*
 * @Description: 文章管理
 * @Author: leo
 * @Date: 2020-02-24 15:04:30
 * @LastEditors: leo
 * @LastEditTime: 2020-02-26 19:00:41
 */

package article_service

import (
	"encoding/json"
	"fmt"

	"github.com/XcXerxes/go-blog-server/models"
	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/XcXerxes/go-blog-server/pkg/gredis"
	"github.com/XcXerxes/go-blog-server/service/cache_service"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

// Add 添加文章方法
func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"created_by":      a.CreatedBy,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
	}
	if err := models.AddArticle(article); err != nil {
		return err
	}
	gredis.LikeDeletes(e.CACHE_ARTICLE)
	return nil
}

// Edit 修改文章
func (a *Article) Edit() error {
	if err := models.EditArticle(a.ID, map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"created_by":      a.CreatedBy,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
	}); err != nil {
		return err
	}
	gredis.LikeDeletes(e.CACHE_ARTICLE)
	return nil
}

// Get 获取单个文章
func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article
	// 初始化redis 的结构体
	cache := cache_service.Article{ID: a.ID}
	// 获取得到当前 key
	key := cache.GetArticleKey()
	// 如果key 存在
	if gredis.Exists(key) {
		// 通过 key 得到 value
		data, err := gredis.Get(key)
		if err != nil {
			fmt.Errorf("%v", err)
		} else {
			// 解析数据
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}
	// 如果没有缓存，直接读取数据库
	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}
	// 同时将数据 存储到 redis 中
	gredis.Set(key, article, 3600)
	return article, nil
}

// GetAll 获取文章列表
func (a *Article) GetAll() ([]*models.Article, error) {
	var (
		articles, cacheArticles []*models.Article
	)
	cache := cache_service.Article{
		TagID:    a.TagID,
		State:    a.State,
		Title:    a.Title,
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cache.GetArticlesKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			fmt.Errorf("%v", err)
		} else {
			json.Unmarshal(data, &cacheArticles)
			return cacheArticles, nil
		}
	}
	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}
	gredis.Set(key, articles, 3600)
	return articles, nil
}

// Delete 删除文章
func (a *Article) Delete() error {
	if err := models.DeleteArticle(a.ID); err != nil {
		return err
	}
	gredis.LikeDeletes(e.CACHE_ARTICLE)
	return nil
}

// ExistByID
func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleById(a.ID)
}

// Count 总条数
func (a *Article) Count() (int, error) {
	return models.GetArticleTotal(a.getMaps())
}

// getMaps 查询条件
func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	// maps["delete_on"] = 0
	if a.State != -1 {
		maps["state"] = a.State
	}
	if a.TagID != -1 {
		maps["tag_id"] = a.TagID
	}
	return maps
}
