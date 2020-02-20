/*
 * @Description: 标签页的模型
 * @Author: leo
 * @Date: 2020-02-20 10:57:29
 * @LastEditors: leo
 * @LastEditTime: 2020-02-20 13:55:33
 */
package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// BeforeCreate  gorm callback
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

// BeforeUpdate gorm callback
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

// GetTags 获取所有标签列表
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	// 因为这里直接定义了返回参数， 操作直接返回
	return
}

// GetTagTotal 获取标签的条数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// ExistTagByName 是否存在标签名
func ExistTagByName(name string) bool {
	var tag Tag
	// 指定从tag数据库表中 检索 存在 ${name} 的所有 id 字段
	// SELECT id FROM tag WHERE name = `name`limit 1;
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// ExistTagById 是否存在id
func ExistTagById(id int) bool {
	var tag Tag
	// 指定从tag数据库表中 检索 存在 ${name} 的所有 id 字段
	// SELECT id FROM tag WHERE name = `name`limit 1;
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// AddTag 新增标签
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

// DeleteTag 删除标签
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

// EditTag 编辑标签
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}
