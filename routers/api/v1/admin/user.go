package admin

import (
	"net/http"

	"github.com/XcXerxes/go-blog-server/models"
	"github.com/XcXerxes/go-blog-server/pkg/app"
	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/XcXerxes/go-blog-server/pkg/util"
	"github.com/XcXerxes/go-blog-server/service/user_service"
	"github.com/gin-gonic/gin"
)

type SigninForm struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// Signin 提交登录
// @Summary 登录
// @Description 登录
// @Accept json
// @produce json
// @param username, password string body SigninForm true "登录参数"
// @Success 200
// @Failure 500
// @Router /signin [post]
func Signin(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form SigninForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	username := form.Username
	password := form.Password
	ok, err := models.CheckAuth(username, password)
	// if err != nil {
	// 	log.Fatal(err)
	// 	appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT, nil)
	// 	return
	// }
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

// GetUserInfo 获取用户信息
// @Summary 获取用户信息
// @Description 获取用户信息
// @Accept json
// @produce json
// @Success 200
// @Router /user [get]
func GetUserInfo(c *gin.Context) {
	appG := app.Gin{c}
	username, errCode := app.GetUserByToken(c)
	if errCode != e.SUCCESS {
		appG.Response(http.StatusBadRequest, errCode, nil)
		return
	}
	// 获取用户的服务
	userService := user_service.User{Username: username}
	exists, err := userService.ExistByUsername()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_USER, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_USER, nil)
		return
	}
	user, err := userService.Get()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_USER, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, user)
}
