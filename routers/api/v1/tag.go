/*
 * @Description: 标签api
 * @Author: leo
 * @Date: 2020-02-19 19:35:07
 * @LastEditors: leo
 * @LastEditTime: 2020-02-20 19:08:55
 */

package v1

import (
	"fmt"
	"log"

	"github.com/XcXerxes/go-blog-server/models"
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
// @Success 200
// @Router /tags [get]
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	// 如果有 name 字段
	if name != "" {
		maps["name"] = name
	}
	state := -1
	if arg := c.Query("state"); arg != "" {
		// 如果传过来的是 string 转为 int
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS
	data["lists"] = models.GetTags(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	// c.JSON(http.StatusOK, gin.H{
	// 	"code": code,
	// 	"msg":  e.GetMsg(code),
	// 	"data": data,
	// })
	models.SendResponse(c, code, nil, data)
	return
}

// AddTag 新增文章标签
// @Summary 新增文章标签
// @Description 新增文章标签
// @Accept json
// @produce json
// @param state name created_by body models.Tag true "新增文章标签"
// @Success 200
// @Router /tags [post]
func AddTag(c *gin.Context) {
	var tag models.Tag
	fmt.Println("==============")
	if err := c.ShouldBind(&tag); err != nil {
		fmt.Println("err========", err)
	}
	name := tag.Name
	state := tag.State
	createdBy := tag.CreatedBy
	// name := c.PostForm("name")
	// // 如果不存在 state 就默认赋值为0 同时转为 int
	// state := com.StrTo(c.DefaultPostForm("state", "0")).MustInt()

	fmt.Println(name, state, createdBy)
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key:%s, err.message: %s", err.Key, err.Message)
		}
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"code": code,
	// 	"msg":  e.GetMsg(code),
	// 	"data": make(map[string]string),
	// })
	models.SendResponse(c, code, nil, make(map[string]string))
}

// EditTag 修改文章标签
func EditTag(c *gin.Context) {
	var tag models.Tag
	id := com.StrTo(c.Param("id")).MustInt()
	if err := c.ShouldBind(&tag); err != nil {
		return
	}
	name := tag.Name
	modifiedBy := tag.ModifiedBy
	var state int = -1
	// if arg := tag.State; arg != "" {
	// 	state = com.StrTo(tag.State).MustInt()
	// 	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	// }

	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.Required(name, "name").Message("name不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagById(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key:%s, err.message: %s", err.Key, err.Message)
		}
	}
	models.SendResponse(c, code, nil, make(map[string]string))
}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}

	valid.Required(id, "id").Message("ID不能为空")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagById(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	models.SendResponse(c, code, nil, make(map[string]string))
}
