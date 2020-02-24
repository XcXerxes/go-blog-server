/*
 * @Description: Admin
 * @Author: leo
 * @Date: 2020-02-21 13:45:17
 * @LastEditors: leo
 * @LastEditTime: 2020-02-24 20:21:02
 */
package admin

import (
	"net/http"

	"github.com/XcXerxes/go-blog-server/pkg/app"

	"github.com/XcXerxes/go-blog-server/models"
	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/XcXerxes/go-blog-server/pkg/util"
	"github.com/gin-gonic/gin"
)

type AuthForm struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// PostAuth 提交登录
// @Summary 登录
// @Description 登录
// @Accept json
// @produce json
// @param username, password string body AuthForm true "登录参数"
// @Success 200
// @Failure 500
// @Router /signin [post]
func PostAuth(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form AuthForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	username := form.Username
	password := form.Password
	ok, err := models.CheckAuth(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusOK, e.ERROR_AUTH, nil)
		return
	}
	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusNonAuthoritativeInfo, e.ERROR_AUTH_TOKEN, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"token": token,
	})
}
