/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-20 14:01:55
 * @LastEditors: leo
 * @LastEditTime: 2020-02-20 14:09:02
 */
package models

import (
	"fmt"
	"testing"

	"github.com/XcXerxes/go-blog-server/models"
)

func TestGetTags(t *testing.T) {
	tagParams := make(map[string]interface{})
	tagParams["name"] = "123"
	data := models.GetTags(1, 10, tagParams)
	fmt.Println(data)
}
