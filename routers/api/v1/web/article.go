/*
 * @Description: web
 * @Author: leo
 * @Date: 2020-02-27 11:57:45
 * @LastEditors: leo
 * @LastEditTime: 2020-02-27 11:59:57
 */
package web

import (
	"net/http"

	"github.com/XcXerxes/go-blog-server/pkg/app"
	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/XcXerxes/go-blog-server/pkg/setting"
	"github.com/XcXerxes/go-blog-server/pkg/util"
	"github.com/XcXerxes/go-blog-server/service/article_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetArticles 获取文章列表
// @Summary 获取文章列表
// @Description 获取文章列表
// @Accept json
// @produce json
// @param page query int true "页数"
// @param title query string false "标题"
// @Success 200
// @Router /articles [get]
func GetArticles(c *gin.Context) {
	appG := app.Gin{c}

	valid := validation.Validation{}

	tagId := -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		//maps["tag_id"] = tagId
		valid.Min(tagId, 1, "tag_id").Message("ID必须大于0")
	}

	title := c.Query("title")
	//code := e.INVALID_PARAMS
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	articleService := article_service.Article{
		State:    0,
		TagID:    tagId,
		Title:    title,
		PageSize: setting.AppSetting.PageSize,
		PageNum:  util.GetPage(c),
	}
	total, err := articleService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_ARTICLE_FAIL, nil)
		return
	}
	articles, err := articleService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": articles,
		"total": total,
	})
}
