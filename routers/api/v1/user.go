package v1

import (
	"github.com/XcXerxes/go-blog-server/pkg/app"
	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/XcXerxes/go-blog-server/service/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserInfo 获取用户信息
// @Summary 获取用户信息
// @Description 获取用户信息
// @Accept json
// @produce json
// @Success 200
// @Router /user [get]
func GetUserInfo(c *gin.Context)  {
	appG := app.Gin{c}
	username, errCode := app.GetUserByToken(c)
	if errCode != e.SUCCESS {
		appG.Response(http.StatusBadRequest, errCode, nil)
		return
	}
	// 获取用户的服务
	userService := user_service.User{Username:username}
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
