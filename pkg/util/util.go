/*
 * @Description: util
 * @Author: leo
 * @Date: 2020-02-25 13:03:21
 * @LastEditors: leo
 * @LastEditTime: 2020-02-25 13:04:14
 */
package util

import "github.com/XcXerxes/go-blog-server/pkg/setting"

// Setup 初始化
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
