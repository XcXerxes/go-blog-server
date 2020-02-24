/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-24 11:05:50
 * @LastEditors: leo
 * @LastEditTime: 2020-02-24 20:16:17
 */
package cache_service

import (
	"strconv"
	"strings"

	"github.com/XcXerxes/go-blog-server/pkg/e"
)

type Tag struct {
	ID    int
	Name  string
	State int

	PageNum  int
	PageSize int
}

func (t *Tag) GetTagsKey() string {
	keys := []string{
		e.CACHE_TAG,
		"LIST",
	}

	if t.Name != "" {
		keys = append(keys, t.Name)
	}
	if t.State >= 0 {
		keys = append(keys, strconv.Itoa(t.State))
	}
	if t.PageNum > 0 {
		keys = append(keys, strconv.Itoa(t.PageNum))
	}
	if t.PageSize > 0 {
		keys = append(keys, strconv.Itoa(t.PageSize))
	}

	return strings.Join(keys, "_")
}
