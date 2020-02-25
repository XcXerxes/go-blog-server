/*
 * @Description: 用户
 * @Author: leo
 * @Date: 2020-02-25 13:17:18
 * @LastEditors: leo
 * @LastEditTime: 2020-02-25 13:41:20
 */
package models

import "github.com/jinzhu/gorm"

type User struct {
	Model
	ID       int    `json:"id"`       // id
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

// GetUser 获取用户信息
func GetUser(username string) (*User, error) {
	var user *User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return user, nil
}

// ExistUserByUsername 判断用户名是否存在库中
func ExistUserByUsername(username string) (bool, error) {
	var user User
	err := db.Select("id").Where("username = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}