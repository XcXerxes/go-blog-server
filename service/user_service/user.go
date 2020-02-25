/*
 * @Description: 用户
 * @Author: leo
 * @Date: 2020-02-25 13:13:41
 * @LastEditors: leo
 * @LastEditTime: 2020-02-25 13:52:53
 */
package user_service

import (
	"encoding/json"
	"fmt"
	"github.com/XcXerxes/go-blog-server/models"
	"github.com/XcXerxes/go-blog-server/pkg/gredis"
	"github.com/XcXerxes/go-blog-server/service/cache_service"
)

type User struct {
	ID       int
	Username string
	Password string
}

// Get 获取单个用户
func (u *User) Get() (*models.User, error) {
	var cacheUser *models.User
	// 初始化redis结构体
	cache := cache_service.User{ID:u.ID}
	// 获取得到当前key
	key := cache.GetUserKey()
	// 如果key 存在
	if gredis.Exists(key) {
		// 通过key 获取value
		data, err := gredis.Get(key)
		if err != nil {
			fmt.Errorf("%v", err)
		} else {
			// 解析数据
			json.Unmarshal(data, &cacheUser)
			return cacheUser, nil
		}
	}
	// 如果没有缓存，直接读取数据库
	user, err := models.GetUser(u.Username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Exist
func (u *User) ExistByUsername() (bool, error) {
	return models.ExistUserByUsername(u.Username)
}