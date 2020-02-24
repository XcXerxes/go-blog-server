/*
 * @Description: 定制返回方法
 * @Author: leo
 * @Date: 2020-02-24 14:16:43
 * @LastEditors: leo
 * @LastEditTime: 2020-02-24 20:16:50
 */
package app

import (
	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

// Response 返回
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
}
