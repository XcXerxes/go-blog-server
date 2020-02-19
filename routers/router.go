/*
 * @Description: 路由
 * @Author: leo
 * @Date: 2020-02-19 19:17:03
 * @LastEditors: leo
 * @LastEditTime: 2020-02-19 19:21:56
 */

package routers

import (
	"github.com/XcXerxes/go-blog-server/pkg/setting"
	"github.com/gin-gonic/gin"
)

// func steupRouter() *gin.Engine {
// 	// 返回 Gin 的 type Engine struct {} 里面包含RouterGroup，相当于创建一个路由Handlers，可以后期绑定各类的路由规则和函数、中间件等
// 	r := gin.Default()
// 	// 创建不同的HTTP方法绑定到Handlers中，也支持POST、PUT、DELETE、PATCH、OPTIONS、HEAD 等常用的Restful方法
// 	// gin.Context Context是gin中的上下文，它允许我们在中间件之间传递变量、管理流、验证JSON请求、响应JSON请求等，
// 	// 在gin中包含大量Context的方法，例如我们常用的DefaultQuery、Query、DefaultPostForm、PostForm等等
// 	r.GET("/test", func(c *gin.Context) {
// 		c.String(200, "hello world!")
// 	})
// 	// 返回当前实例
// 	return r
// }

// InitRouter 初始化路由管理器
func InitRouter() *gin.Engine {
	// 返回 Gin 的 type Engine struct {} 里面包含RouterGroup，相当于创建一个路由Handlers，可以后期绑定各类的路由规则和函数、中间件等
	r := gin.Default()
	// 日志的中间件
	// r.Use(gin.Logger())
	//
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	// 创建不同的HTTP方法绑定到Handlers中，也支持POST、PUT、DELETE、PATCH、OPTIONS、HEAD 等常用的Restful方法
	// gin.Context Context是gin中的上下文，它允许我们在中间件之间传递变量、管理流、验证JSON请求、响应JSON请求等，
	// 在gin中包含大量Context的方法，例如我们常用的DefaultQuery、Query、DefaultPostForm、PostForm等等
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})
	return r
}
