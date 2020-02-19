module github.com/XcXerxes/go-blog-server

go 1.13

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.52.0
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20200217220822-9197077df867 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/ini.v1 v1.52.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace (
	github.com/XcXerxes/go-blog-server/conf => /Users/zhangjie/xinbo/myself/go/go-blog-server/conf
	github.com/XcXerxes/go-blog-server/middleware => /Users/zhangjie/xinbo/myself/go/go-blog-server/middleware
	github.com/XcXerxes/go-blog-server/models => /Users/zhangjie/xinbo/myself/go/go-blog-server/models
	github.com/XcXerxes/go-blog-server/pkg/e => /Users/zhangjie/xinbo/myself/go/go-blog-server/pkg/e
	github.com/XcXerxes/go-blog-server/pkg/setting => /Users/zhangjie/xinbo/myself/go/go-blog-server/pkg/setting
	github.com/XcXerxes/go-blog-server/pkg/util => /Users/zhangjie/xinbo/myself/go/go-blog-server/pkg/util
	github.com/XcXerxes/go-blog-server/routers => /Users/zhangjie/xinbo/myself/go/go-blog-server/routers
)
