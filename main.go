/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-19 15:08:40
 * @LastEditors: leo
 * @LastEditTime: 2020-02-20 12:47:28
 */
package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/XcXerxes/go-blog-server/docs"
	"github.com/XcXerxes/go-blog-server/pkg/setting"
	"github.com/XcXerxes/go-blog-server/routers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 博客接口文档说明
// @version 1.0.0
// @description  博客系统的api接口文档
// @BasePath /api/v1/
func main() {
	r := routers.InitRouter()

	// swagger
	// url := ginSwagger.URL("http://localhost:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort), // 监听的TCP地址
		Handler:        r,                                    // http 句柄 实质为 ServeHTTP
		ReadTimeout:    setting.ReadTimeout,                  //允许读取的最大时间
		WriteTimeout:   setting.WriteTimeout,                 // 允许写入的最大时间
		MaxHeaderBytes: 1 << 20,                              // 请求投的最大字节数
	}
	log.Fatal(s.ListenAndServe())
}
