/*
 * @Description: Admin
 * @Author: leo
 * @Date: 2020-02-21 13:45:17
 * @LastEditors: leo
 * @LastEditTime: 2020-02-21 18:52:29
 */
package admin

import (
	"log"

	"github.com/XcXerxes/go-blog-server/models"
	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/XcXerxes/go-blog-server/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type AuthBody struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// PostAuth 提交登录
// @Summary 登录
// @Description 登录
// @Accept json
// @produce json
// @param username body AuthBody true "登录参数"
// @Success 200
// @Router /signin [post]
func PostAuth(c *gin.Context) {
	code := e.INVALID_PARAMS
	var auth AuthBody
	if err := c.ShouldBind(&auth); err != nil {
		models.SendResponse(c, code, nil, nil)
		return
	}
	username := auth.Username
	password := auth.Password

	valid := validation.Validation{}
	ok, _ := valid.Valid(&auth)
	data := make(map[string]interface{})
	if ok {
		if models.CheckAuth(username, password) {
			if token, err := util.GenerateToken(username, password); err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key:%s, err.message: %s", err.Key, err.Message)
		}
	}

	models.SendResponse(c, code, nil, data)
}
