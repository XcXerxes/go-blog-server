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
	ERROR_GET_ARTICLE_FAIL: "获取文章失败",
	ERROR_COUNT_ARTICLE_FAIL: "获取条数失败",
	ERROR_GET_ARTICLES_FAIL: "获取文章列表失败",
	ERROR_ADD_ARTICLE_FAIL: "添加文章失败",
	ERROR_CHECK_EXIST_ARTICLE_FAIL: "文章id不存在",
	ERROR_EDIT_ARTICLE_FAIL: "文章编辑失败",
	ERROR_DELETE_ARTICLE_FAIL: "删除文章失败",
	ERROR_GET_TAGS_FAIL: "获取标签列表失败",
	ERROR_COUNT_TAG_FAIL: "回去标签条数失败",
	ERROR_ADD_TAG_FAIL: "添加标签失败",
	ERROR_EDIT_TAG_FAIL: "编辑标签失败",
	ERROR_DELETE_TAG_FAIL: "删除标签失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token 鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token 已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_NOT_EXIST_USER: "用户不存在",
}

// GetMsg 根据错误码得到错误信息
func GetMsg(code int) string {
	if msg, ok := MsgFlags[code]; ok {
		return msg
	}

	return MsgFlags[ERROR]
}
