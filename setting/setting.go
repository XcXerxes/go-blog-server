package main

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File
	RumMode string

	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	PageSize int
	JwtSecret string
)

func init()  {
	var err error
	if Cfg, err = ini.Load("conf/app.ini"); err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}

// 加载基础
func LoadBase () {
	RumMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// 加载服务配置
func LoadServer()  {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout =  time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second    
}

// 加载app
func LoadApp()  {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
