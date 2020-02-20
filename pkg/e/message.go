/*
 * @Description: 错误信息
 * @Author: leo
 * @Date: 2020-02-19 15:34:09
 * @LastEditors: leo
 * @LastEditTime: 2020-02-20 18:43:01
 */

package e

var MsgFlags = map[int]string{
	SUCCESS:                        "成功",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "标签名称已存在",
	ERROR_NOT_EXIST_TAG:            "标签名称不存在",
	ERROR_NOT_EXIST_ARTICLE:        "文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token 鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token 已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

// GetMsg 根据错误码得到错误信息
func GetMsg(code int) string {
	if msg, ok := MsgFlags[code]; ok {
		return msg
	}

	return MsgFlags[ERROR]
}
