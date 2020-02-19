/*
 * @Description: 标签api
 * @Author: leo
 * @Date: 2020-02-19 19:35:07
 * @LastEditors: leo
 * @LastEditTime: 2020-02-19 19:42:56
 */

package v1

import (
	"net/http"

	"github.com/XcXerxes/go-blog-server/models"
	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/XcXerxes/go-blog-server/pkg/setting"
	"github.com/XcXerxes/go-blog-server/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetTags 获取多个文章标签 带字段筛选
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
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// AddTag 新增文章标签
func AddTag(c *gin.Context) {

}

// EditTag 修改文章标签
func EditTag(c *gin.Context) {

}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {

}
