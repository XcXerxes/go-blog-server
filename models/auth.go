/*
 * @Description: admin
 * @Author: leo
 * @Date: 2020-02-21 13:38:02
 * @LastEditors: leo
 * @LastEditTime: 2020-02-21 17:47:32
 */

package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"` // id
	Username string `json:"username"`              // 用户名
	Password string `json:"password"`              // 密码
}

// BeforeCreate  gorm callback
// func (auth *Auth) BeforeCreate(scope *gorm.Scope) error {
// 	scope.SetColumn("CreatedOn", time.Now().Unix())
// 	scope.SetColumn("ModifiedOn", time.Now().Unix())
// 	return nil
// }

// BeforeUpdate gorm callback
// func (auth *Auth) BeforeUpdate(scope *gorm.Scope) error {
// 	scope.SetColumn("ModifiedOn", time.Now().Unix())
// 	return nil
// }

// CheckAuth 验证登录信息
func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}
