/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-19 15:08:40
 * @LastEditors: leo
 * @LastEditTime: 2020-02-19 18:31:14
 */
package main

import (
	"net/http"

	"github.com/XcXerxes/go-blog-server/setting"
	"github.com/gin-gonic/gin"
)

func steupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello world!")
	})
	return r
}

func main() {
	r := steupRouter()
	s := &http.Server{
		Addr:           setting.HTTPPort,
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
