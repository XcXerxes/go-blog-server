/*
 * @Description: 设置
 * @Author: leo
 * @Date: 2020-02-19 15:08:40
 * @LastEditors: leo
 * @LastEditTime: 2020-02-24 20:15:16
 */
package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	PageSize        int
	JwtSecret       string
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

func Setup() {
	// 初始化加载 conf/app.ini 文件
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadDataBase(Cfg)
	LoadServer(Cfg)
	LoadApp(Cfg)
	LoadRedis(Cfg)
}

// LoadApp 加载基础
func LoadApp(cfg *ini.File) {
	// 读取配置文件中 RUN_MODE 属性 如果没有值，默认使用 debug 默认分区可以使用空字符串表示
	// MustString 转换为 string
	if err := cfg.Section("app").MapTo(AppSetting); err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
}

// LoadServer 加载服务配置
func LoadServer(cfg *ini.File) {
	if err := cfg.Section("server").MapTo(ServerSetting); err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// LoadDataBase 加载数据库配置
func LoadDataBase(cfg *ini.File) {
	if err := cfg.Section("database").MapTo(DatabaseSetting); err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}
}

// LoadRedis 加载redis
func LoadRedis(cfg *ini.File) {
	if err := cfg.Section("redis").MapTo(RedisSetting); err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}
}
