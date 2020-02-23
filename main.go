/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-19 15:08:40
 * @LastEditors: leo
 * @LastEditTime: 2020-02-23 13:48:08
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	_ "github.com/XcXerxes/go-blog-server/docs"
	"github.com/XcXerxes/go-blog-server/models"
	"github.com/XcXerxes/go-blog-server/pkg/setting"
	"github.com/XcXerxes/go-blog-server/routers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	setting.Setup()
	models.Setup()
	// logging.Setup()
}

// @title 博客接口文档说明
// @version 1.0.0
// @description  博客系统的api接口文档
// @BasePath /api/v1
func main() {
	r := routers.InitRouter()

	// swagger
	// url := ginSwagger.URL("http://localhost:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	s := &http.Server{
		Addr:           endPoint,       // 监听的TCP地址
		Handler:        r,              // http 句柄 实质为 ServeHTTP
		ReadTimeout:    readTimeout,    //允许读取的最大时间
		WriteTimeout:   writeTimeout,   // 允许写入的最大时间
		MaxHeaderBytes: maxHeaderBytes, // 请求投的最大字节数
	}
	exec.Command(`open`, `http://localhost:8000/swagger/index.html`).Start()
	log.Fatal(s.ListenAndServe())
}
