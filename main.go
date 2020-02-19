/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-19 15:08:40
 * @LastEditors: leo
 * @LastEditTime: 2020-02-19 19:20:48
 */
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/XcXerxes/go-blog-server/pkg/setting"
	"github.com/XcXerxes/go-blog-server/routers"
)

func main() {
	r := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort), // 监听的TCP地址
		Handler:        r,                                    // http 句柄 实质为 ServeHTTP
		ReadTimeout:    setting.ReadTimeout,                  //允许读取的最大时间
		WriteTimeout:   setting.WriteTimeout,                 // 允许写入的最大时间
		MaxHeaderBytes: 1 << 20,                              // 请求投的最大字节数
	}
	log.Fatal(s.ListenAndServe())
}
