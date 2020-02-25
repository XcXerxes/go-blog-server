package cache_service

import (
"strconv"

"github.com/XcXerxes/go-blog-server/pkg/e"
)

type User struct {
	ID    int
	UserName  string
}

// GetArticleKey 获取key
func (u *User) GetUserKey() string {
	return e.CACHE_USER + "_" + strconv.Itoa(u.ID)
}

