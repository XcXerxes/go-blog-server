/*
 * @Description: jwt
 * @Author: leo
 * @Date: 2020-02-21 12:38:02
 * @LastEditors: leo
 * @LastEditTime: 2020-02-21 14:01:37
 */

package util

import (
	"time"

	"github.com/XcXerxes/go-blog-server/pkg/setting"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	jwt.StandardClaims
}

// GenerateToken 创建生成token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "blog",
		},
	}
	// 包含SigningMethodHS256、SigningMethodHS384、SigningMethodHS512三种crypto.Hash方案
	tokenChaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenChaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 解析 token
func ParseToken(token string) (*Claims, error) {
	// 用于解析鉴权的声明
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		//  验证基于时间的声明exp, iat, nbf
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
