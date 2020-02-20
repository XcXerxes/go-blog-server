/*
 * @Description: 文章
 * @Author: leo
 * @Date: 2020-02-20 18:53:54
 * @LastEditors: leo
 * @LastEditTime: 2020-02-20 21:02:00
 */
package v1

import (
	"log"

	"github.com/XcXerxes/go-blog-server/models"
	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/XcXerxes/go-blog-server/pkg/setting"
	"github.com/XcXerxes/go-blog-server/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetArticle 获取单个文章
// @Summary 获取单个文章
// @Description 获取单个文章
// @Accept json
// @produce json
// @param id query int true "id"
// @Success 200
// @Router /articles/{id} [get]
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()

	valid := validation.Validation

	valid.Min(id, 1, "id").Message("Id 必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleById(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key:%s, err.message: %s", err.Key, err.Message)
		}
	}
	models.SendResponse(c, code, nil, data)
}

// GetArticles 获取文章列表
// @Summary 获取文章列表
// @Description 获取文章列表
// @Accept json
// @produce json
// @param page query int true "页数"
// @param title query string true "标题"
// @Success 200
// @Router /articles/{id} [get]
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{}, 2)
	maps := make(map[string]interface{}, 4)

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		valid.Min(tagId, 1, "tag_id").Message("ID必须大于0")
	}
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key:%s, err.message: %s", err.Key, err.Message)
		}
	}
	models.SendResponse(c, code, nil, data)
}

// AddArticle 新增文章
func AddArticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBind(&article); err != nil {
		return
	}
	title := article.Title
	tagId := com.StrTo(article.TagID).MustInt()
	content := article.Content
	state := com.StrTo(article.State).MustInt()
	desc := article.Desc
	createdBy := article.CreatedBy

	valid := validation.Validation{}

	valid.Required("title", title).Message("标题不能为空")
	valid.Required("content", content).Message("内容不能为空")
	valid.Required("tag_id", tagId).Message("tag_id 不能为空")
	valid.Required("desc", desc).Message("desc 不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistTagById(tagId) {
			data := make(map[string]interface{}, 0)
		}
	}

}

// EditArticle 修改文章
func EditArticle(c *gin.Context) {

}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {

}
