package app

import (
	"fmt"
	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/XcXerxes/go-blog-server/pkg/util"
	"github.com/gin-gonic/gin"
)
// GetUserByToken 根据token 解析出 用户名
func GetUserByToken(c *gin.Context) (string, int)  {
	authorization := c.GetHeader("Authorization")
	fmt.Printf("authorization================", authorization)
	claims, err := util.ParseToken(authorization)
	if err != nil {
		return "", e.ERROR
	}
	return claims.Username, e.SUCCESS
}
