/*
 * @Description: 分页获取方法
 * @Author: leo
 * @Date: 2020-02-19 16:46:14
 * @LastEditors: leo
 * @LastEditTime: 2020-02-19 18:55:17
 */

package util

import (
	"github.com/XcXerxes/go-blog-server/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetPage 根据获取到的 当前页数 返回 数据库中的offset
func GetPage(c *gin.Context) int {
	result := 0
	if page, _ := com.StrTo(c.Query("page")).Int(); page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}
	return result
}
