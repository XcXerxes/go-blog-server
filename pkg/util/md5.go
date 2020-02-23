/*
 * @Description: md5
 * @Author: leo
 * @Date: 2020-02-23 14:00:58
 * @LastEditors: leo
 * @LastEditTime: 2020-02-23 14:02:46
 */

package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMd5 将图片名转为 md5
func EncodeMd5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
