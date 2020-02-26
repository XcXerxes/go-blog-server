/*
 * @Description: 模型初始化
 * @Author: leo
 * @Date: 2020-02-19 17:04:46
 * @LastEditors: leo
 * @LastEditTime: 2020-02-26 13:06:20
 */

package models

import (
	"fmt"
	"log"
	"time"

	"github.com/XcXerxes/go-blog-server/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Model 结构体
type Model struct {
	ID         int `gorm:"primary_key" json:"id"` // id标识
	CreatedOn  int `json:"created_on"`            // 创建时间
	ModifiedOn int `json:"modified_on"`           // 修改时间
}

func Setup() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	// sec, err := setting.Cfg.GetSection("database")
	// if err != nil {
	// 	log.Fatal(2, "Fail to get section 'database': %v", err)
	// }
	dbType = setting.DatabaseSetting.Type
	dbName = setting.DatabaseSetting.Name
	user = setting.DatabaseSetting.User
	password = setting.DatabaseSetting.Password
	host = setting.DatabaseSetting.Host
	tablePrefix = setting.DatabaseSetting.TablePrefix

	if db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName)); err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
}

// CloseDB 关闭db
func CloseDB() {
	defer db.Close()
}

// updateTimeStampForCreateCallback 更新创建时间的 回调
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	// 检测db 是否有错误
	if !scope.HasError() {
		// 获取当前的时间戳
		nowTime := time.Now().Unix()
		// 获取得到创建时间字段
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
		// 获取所有的字段
		// fields := scope.Fields()
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}
