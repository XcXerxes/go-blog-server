/*
 * @Description: 设置
 * @Author: leo
 * @Date: 2020-02-19 15:08:40
 * @LastEditors: leo
 * @LastEditTime: 2020-02-19 16:54:11
 */
package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg     *ini.File
	RumMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	// 初始化加载 conf/app.ini 文件
	if Cfg, err = ini.Load("conf/app.ini"); err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}

// 加载基础
func LoadBase() {
	// 读取配置文件中 RUN_MODE 属性 如果没有值，默认使用 debug 默认分区可以使用空字符串表示
	// MustString 转换为 string
	RumMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// 加载服务配置
func LoadServer() {
	// 读取配置文件中 server分区
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	// 通过key 找到当前分区中的值
	// MustInt 转换为 int
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

// 加载app
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
