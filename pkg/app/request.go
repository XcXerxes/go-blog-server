/*
 * @Description: 定制返回方法
 * @Author: leo
 * @Date: 2020-02-24 14:13:45
 * @LastEditors: leo
 * @LastEditTime: 2020-02-24 14:16:36
 */
package app

import (
	"log"

	"github.com/astaxie/beego/validation"
)

// MarkErrors 提前返回错误
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.Printf("err.key:%s, err.message: %s", err.Key, err.Message)
	}

	return
}
