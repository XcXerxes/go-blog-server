/*
 * @Description: jwt middleware
 * @Author: leo
 * @Date: 2020-02-21 13:10:37
 * @LastEditors: leo
 * @LastEditTime: 2020-02-21 19:07:19
 */

package jwt

import (
	"time"

	"github.com/XcXerxes/go-blog-server/models"
	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/XcXerxes/go-blog-server/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// JWT 的中间件方法
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			code int
			data interface{}
		)
		code = e.SUCCESS
		token := com.StrTo(c.GetHeader("Authorization")).String()
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			if clamis, err := util.ParseToken(token); err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > clamis.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			models.SendResponse(c, code, nil, data)
			c.Abort()
			return
		}
		c.Next()
	}
}
