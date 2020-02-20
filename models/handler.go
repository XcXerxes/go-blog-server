/*
 * @Description: handler
 * @Author: leo
 * @Date: 2020-02-20 18:38:05
 * @LastEditors: leo
 * @LastEditTime: 2020-02-20 19:07:06
 */
package models

import (
	"net/http"

	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:code`
	Msg  string      `json:msg`
	Data interface{} `json:data`
}

// SendResponse 统一处理返回值
func SendResponse(c *gin.Context, code int, err error, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  e.MsgFlags[code],
		Data: data,
	})
}
