/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-24 16:13:09
 * @LastEditors: leo
 * @LastEditTime: 2020-02-26 15:06:38
 */
package tag_service

import (
	"encoding/json"
	"fmt"

	"github.com/XcXerxes/go-blog-server/models"
	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/XcXerxes/go-blog-server/pkg/gredis"
	"github.com/XcXerxes/go-blog-server/service/cache_service"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

// ExistByName 名称是否存在
func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

// ExistByID id是否存在
func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagById(t.ID)
}

// Add 添加标签
func (t *Tag) Add() error {
	if err := models.AddTag(t.Name, t.State, t.CreatedBy); err != nil {
		return err
	}
	gredis.LikeDeletes(e.CACHE_TAG)
	return nil
}

// Edit 编辑标签.
func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name
	if t.State >= 0 {
		data["state"] = t.State
	}
	if err := models.EditTag(t.ID, data); err != nil {
		return err
	}
	gredis.LikeDeletes(e.CACHE_TAG)
	return nil
}

// Delete 删除标签
func (t *Tag) Delete() error {
	if err := models.DeleteTag(t.ID); err != nil {
		return err
	}
	gredis.LikeDeletes(e.CACHE_TAG)
	return nil
}

// Count 获取条数
func (t *Tag) Count() (int, error) {
	return models.GetArticleTotal(t.getMaps())
}

// GetAll 获取标签列表
func (t *Tag) GetAll() ([]*models.Tag, error) {
	var (
		tags, cacheTags []*models.Tag
	)
	cache := cache_service.Tag{
		State:    t.State,
		PageNum:  t.PageNum,
		PageSize: t.PageSize,
		Name:     t.Name,
	}
	key := cache.GetTagsKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			fmt.Errorf("%v", err)
		} else {
			json.Unmarshal(data, &cacheTags)
			return cacheTags, nil
		}
	}
	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}
	gredis.Set(key, tags, 3600)
	return tags, nil
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}
	return maps
}
