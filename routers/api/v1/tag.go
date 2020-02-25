/*
 * @Description: 标签api
 * @Author: leo
 * @Date: 2020-02-19 19:35:07
 * @LastEditors: leo
 * @LastEditTime: 2020-02-25 20:37:19
 */

package v1

import (
	"net/http"

	"github.com/XcXerxes/go-blog-server/pkg/app"
	"github.com/XcXerxes/go-blog-server/service/tag_service"

	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/XcXerxes/go-blog-server/pkg/setting"
	"github.com/XcXerxes/go-blog-server/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetTags 获取多个文章标签 带字段筛选
// @Summary 获取文章标签列表
// @Description 获取文章标签列表 带字段筛选 带分页
// @Accept json
// @produce json
// @param page query int true "当前页数"
// @param name query string false "标签名称"
// @param state query int false "标签名称"
// @Success 200
// @Router /tags [get]
func GetTags(c *gin.Context) {
	appG := app.Gin{c}
	valid := validation.Validation{}
	name := com.StrTo(c.Query("name")).String()
	state := -1
	if arg := c.Query("state"); arg != "" {
		// 如果传过来的是 string 转为 int
		state = com.StrTo(arg).MustInt()
		valid.Range(state, -1, 1, "state")
	}
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	tagService := tag_service.Tag{
		Name:     name,
		State:    state,
		PageSize: setting.AppSetting.PageSize,
		PageNum:  util.GetPage(c),
	}
	tags, err := tagService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
		return
	}

	count, err := tagService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_TAG_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": tags,
		"total": count,
	})
}

type AddTagForm struct {
	Name      string `json:"name" valid:"Required;MaxSize(100)"`       // 名称
	State     int    `json:"state" valid:"Required;Range(0, 1)"`       // 禁用 or 启用
	CreatedBy string `json:"created_by" valid:"Required;MaxSize(100)"` // 创建人
}

// AddTag 新增文章标签
// @Summary 新增文章标签
// @Description 新增文章标签
// @Accept json
// @produce json
// @param name created_by state body AddTagForm true "新增文章标签"
// @Success 200
// @Failure 500
// @Router /tags [post]
func AddTag(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form AddTagForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	tagService := tag_service.Tag{
		Name:      form.Name,
		CreatedBy: form.CreatedBy,
		State:     form.State,
	}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
		return
	}
	err = tagService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditTagForm struct {
	ID         int    `json:"id" valid:"Required;Min(1)"`                // id
	Name       string `json:"name" valid:"Required;MaxSize(100)"`        // 名称
	State      int    `json:"state" valid:"Requred;Range(0, 1)"`         // 禁用 or 启用
	ModifiedBy string `json:"modified_by" valid:"Required;MaxSize(100)"` // 修改人
}

// EditTag 修改文章标签
// @Summary 修改文章标签
// @Description 修改文章标签
// @Accept json
// @produce json
// @Param id path int true "ID"
// @Param name created_by state body EditTagForm true "新增文章标签"
// @Success 200
// @Router /tags/{id} [put]
func EditTag(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form = EditTagForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	tagService := tag_service.Tag{
		ID:         form.ID,
		Name:       form.Name,
		ModifiedBy: form.ModifiedBy,
		State:      form.State,
	}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = tagService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// DeleteTag 删除文章标签
// @Summary 删除文章标签
// @Description 删除文章标签
// @Accept json
// @produce json
// @param id path int true "ID"
// @Success 200
// @Router /tags/{id} [delete]
func DeleteTag(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	tagService := tag_service.Tag{ID: id}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	if err := tagService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
