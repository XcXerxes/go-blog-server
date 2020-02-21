/*
 * @Description: 路由
 * @Author: leo
 * @Date: 2020-02-19 19:17:03
 * @LastEditors: leo
 * @LastEditTime: 2020-02-21 19:02:17
 */

package routers

import (
	"github.com/XcXerxes/go-blog-server/middleware/jwt"
	"github.com/XcXerxes/go-blog-server/pkg/setting"
	admin "github.com/XcXerxes/go-blog-server/routers/api/admin"
	v1 "github.com/XcXerxes/go-blog-server/routers/api/v1"
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

	r.POST("/api/v1/signin", admin.PostAuth)
	// 注册路由
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 标签路由
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// 文章路由

		// 文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		// 新增文章
		apiv1.POST("/articles", v1.AddArticle)
		// 更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		// 删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	// 创建不同的HTTP方法绑定到Handlers中，也支持POST、PUT、DELETE、PATCH、OPTIONS、HEAD 等常用的Restful方法
	// gin.Context Context是gin中的上下文，它允许我们在中间件之间传递变量、管理流、验证JSON请求、响应JSON请求等，
	// 在gin中包含大量Context的方法，例如我们常用的DefaultQuery、Query、DefaultPostForm、PostForm等等
	// r.GET("/test", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "success"})
	// })
	return r
}
