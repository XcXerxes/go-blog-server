/*
 * @Description: 文章
 * @Author: leo
 * @Date: 2020-02-20 18:53:54
 * @LastEditors: leo
 * @LastEditTime: 2020-02-27 11:19:42
 */
package admin

import (
	"net/http"

	"github.com/XcXerxes/go-blog-server/service/article_service"
	"github.com/XcXerxes/go-blog-server/service/tag_service"

	"github.com/XcXerxes/go-blog-server/pkg/app"
	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/XcXerxes/go-blog-server/pkg/setting"
	"github.com/XcXerxes/go-blog-server/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type ArticleBody struct {
	Title     string `json:"title"`      // 标题
	TagID     string `json:"tag_id"`     // 标签id
	State     int    `json:"state"`      // 禁用或启用
	Desc      string `json:"desc"`       // 描述信息
	Content   string `json:"content"`    // 内容
	CreatedBy string `json:"created_by"` // 创建人
}

// GetArticle 获取单个文章
// @Summary 获取单个文章
// @Description 获取单个文章
// @Accept json
// @produce json
// @param id query int true "id"
// @Success 200
// @Router /articles/{id} [get]
func GetArticle(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("Id 必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	// 获取到 文章的服务
	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}
	article, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, article)
}

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

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		//maps["state"] = state
		valid.Range(state, -1, 1, "state").Message("状态只允许0或1")
	}
	tagId := -1
	if arg := c.Query("tag_id"); arg != "" {
		state = com.StrTo(arg).MustInt()
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
		State:    state,
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

type AddArticleForm struct {
	TagID         int    `json:"tag_id" valid:"Required;Min(1)"`                // 分类id
	Title         string `json:"title" valid:"Required;MaxSize(100)"`           // 标题
	Desc          string `json:"desc" valid:"Required;MaxSize(100)"`            // 描述
	Content       string `json:"content" valid:"Required"`                      // 内容
	CreatedBy     string `json:"created_by" valid:"Required;MaxSize(100)"`      // 创建人
	CoverImageUrl string `json:"cover_image_url" valid:"Required;MaxSize(255)"` // 封面图片
	State         int    `json:"state" valid:"Range(0, 1)"`                     // 启用or 禁用
}

// AddArticle 新增文章
// @Summary 新增文章
// @Description 新增文章
// @Accept json
// @produce json
// @param tag_id title desc content, created_by cover_image_url state body AddArticleForm true "新增文章"
// @Success 200
// @Router /articles [post]
func AddArticle(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form = AddArticleForm{CreatedBy: c.GetString("username")}
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	tagService := tag_service.Tag{ID: form.TagID}
	// 判断tabId 是否在 tag表中
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	articleService := article_service.Article{
		TagID:         form.TagID,
		Title:         form.Title,
		Desc:          form.Desc,
		Content:       form.Content,
		CoverImageUrl: form.CoverImageUrl,
		State:         form.State,
		CreatedBy:     form.CreatedBy,
	}

	if err := articleService.Add(); err != nil {
		appG.Response(http.StatusOK, e.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditArticleForm struct {
	ID            int    `json:"id" valid:"Required;Min(1)"`                    // id
	TagID         int    `json:"tag_id" valid:"Required;Min(1)"`                // 分类id
	Title         string `json:"title" valid:"Required;MaxSize(100)"`           // 标题
	Desc          string `json:"desc" valid:"Required;MaxSize(100)"`            // 描述
	Content       string `json:"content" valid:"Required"`                      // 内容
	ModifiedBy    string `json:"modified_by" valid:"Required;MaxSize(100)"`     // 修改人
	CoverImageUrl string `json:"cover_image_url" valid:"Required;MaxSize(255)"` // 封面图片
	State         int    `json:"state" valid:"Range(0, 1)"`                     // 启用or 禁用
}

// EditArticle 修改文章
// @Summary 修改文章
// @Description 修改文章
// @Accept json
// @produce json
// @param id path int true "唯一id"
// @param tag_id title desc content, created_by cover_image_url state body EditArticleForm true "新增文章"
// @Success 200
// @Router /articles/{id} [put]
func EditArticle(c *gin.Context) {
	var (
		appG = app.Gin{c}
		// 定义 form 同时将 得到的 id 参数传给 form
		form = EditArticleForm{ID: com.StrTo(c.Param("id")).MustInt(), ModifiedBy: c.GetString("username")}
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	articleService := article_service.Article{
		ID:            form.ID,
		TagID:         form.TagID,
		Title:         form.Title,
		Desc:          form.Desc,
		Content:       form.Content,
		CoverImageUrl: form.CoverImageUrl,
		State:         form.State,
		ModifiedBy:    form.ModifiedBy,
	}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	err = articleService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// DeleteArticle 删除文章
// @Summary 删除文章
// @Description 删除文章
// @Accept json
// @produce json
// @param id path int true "唯一id"
// @Success 200
// @Router /articles/{id} [delete]
func DeleteArticle(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("Id必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}
	err = articleService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
