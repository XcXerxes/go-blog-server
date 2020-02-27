/*
 * @Description: 路由
 * @Author: leo
 * @Date: 2020-02-19 19:17:03
 * @LastEditors: leo
 * @LastEditTime: 2020-02-27 11:59:31
 */

package routers

import (
	"net/http"

	"github.com/XcXerxes/go-blog-server/middleware/jwt"
	"github.com/XcXerxes/go-blog-server/pkg/setting"
	"github.com/XcXerxes/go-blog-server/pkg/upload"
	"github.com/XcXerxes/go-blog-server/routers/api"
	"github.com/XcXerxes/go-blog-server/routers/api/v1/admin"
	"github.com/XcXerxes/go-blog-server/routers/api/v1/web"
	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8989"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "content-type", "x-requested-with"},
		AllowCredentials: true,
	}))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	gin.SetMode(setting.ServerSetting.RunMode)
	r.POST("/api/v1/admin/signin", admin.PostAuth)
	r.POST("/api/v1/admin/upload", api.UploadImage)
	// 注册后台路由
	adminApiv1 := r.Group("/api/v1/admin")
	adminApiv1.Use(jwt.JWT())
	{
		// 用户路由
		adminApiv1.GET("/user", admin.GetUserInfo)
		// 标签路由
		adminApiv1.GET("/tags", admin.GetTags)
		adminApiv1.POST("/tags", admin.AddTag)
		adminApiv1.PUT("/tags/:id", admin.EditTag)
		adminApiv1.DELETE("/tags/:id", admin.DeleteTag)

		// 文章路由

		// 文章列表
		adminApiv1.GET("/articles", admin.GetArticles)
		// 指定文章
		adminApiv1.GET("/articles/:id", admin.GetArticle)
		// 新增文章
		adminApiv1.POST("/articles", admin.AddArticle)
		// 更新指定文章
		adminApiv1.PUT("/articles/:id", admin.EditArticle)
		// 删除指定文章
		adminApiv1.DELETE("/articles/:id", admin.DeleteArticle)
	}
	// 注册前台路由
	webApiv1 := r.Group("/api/v1/web")
	{
		webApiv1.GET("/tags", admin.GetTags)
		webApiv1.GET("/articles", web.GetArticles)
		webApiv1.GET("/articles/:id", admin.GetArticle)
	}
	// 创建不同的HTTP方法绑定到Handlers中，也支持POST、PUT、DELETE、PATCH、OPTIONS、HEAD 等常用的Restful方法
	// gin.Context Context是gin中的上下文，它允许我们在中间件之间传递变量、管理流、验证JSON请求、响应JSON请求等，
	// 在gin中包含大量Context的方法，例如我们常用的DefaultQuery、Query、DefaultPostForm、PostForm等等
	// r.GET("/test", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "success"})
	// })
	return r
}
