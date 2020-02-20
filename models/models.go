/*
 * @Description: 模型初始化
 * @Author: leo
 * @Date: 2020-02-19 17:04:46
 * @LastEditors: leo
 * @LastEditTime: 2020-02-20 14:40:51
 */

package models

import (
	"fmt"
	"log"

	"github.com/XcXerxes/go-blog-server/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Model 结构体
type Model struct {
	ID         int `gorm:"primary_key" json:"id"` // id标识
	CreatedOn  int `json:created_on`              // 创建时间
	ModifiedOn int `json:modified_on`             // 修改时间
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

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
}

// CloseDB 关闭db
func CloseDB() {
	defer db.Close()
}
