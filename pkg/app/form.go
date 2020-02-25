/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-24 18:59:59
 * @LastEditors: leo
 * @LastEditTime: 2020-02-25 20:35:00
 */
package app

import (
	"net/http"

	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid 绑定并且验证数据
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}
	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}
	return http.StatusOK, e.SUCCESS
}
