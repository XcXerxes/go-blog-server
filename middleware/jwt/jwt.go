/*
 * @Description: jwt middleware
 * @Author: leo
 * @Date: 2020-02-21 13:10:37
 * @LastEditors: leo
 * @LastEditTime: 2020-02-26 13:09:43
 */

package jwt

import (
	"net/http"
	"time"

	"github.com/XcXerxes/go-blog-server/pkg/app"
	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/XcXerxes/go-blog-server/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// JWT 的中间件方法
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{c}
		// code = e.SUCCESS
		token := com.StrTo(c.GetHeader("Authorization")).String()
		if token == "" {
			appG.Response(http.StatusBadRequest, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
			return
		}
		clamis, err := util.ParseToken(token)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
			return
		}
		if time.Now().Unix() > clamis.ExpiresAt {
			appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT, nil)
			return
		}
		c.Set("username", clamis.Username)
		c.Next()
		// if token == "" {
		// 	code = e.INVALID_PARAMS
		// } else {
		// 	if clamis, err := util.ParseToken(token); err != nil {
		// 		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		// 	} else if time.Now().Unix() > clamis.ExpiresAt {
		// 		code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		// 	}
		// }
		// if code != e.SUCCESS {
		// 	models.SendResponse(c, code, nil, data)
		// 	c.Abort()
		// 	return
		// }
		// c.Next()
	}
}
