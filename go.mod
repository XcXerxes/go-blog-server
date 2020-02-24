module github.com/XcXerxes/go-blog-server

go 1.13

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.52.0
	github.com/go-openapi/spec v0.19.6 // indirect
	github.com/go-openapi/swag v0.19.7 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.12
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/robfig/cron/v3 v3.0.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5
	github.com/unknwon/com v1.0.1
	golang.org/x/net v0.0.0-20200222033325-078779b8f2d8 // indirect
	golang.org/x/sys v0.0.0-20200219091948-cb0a6d8edb6c // indirect
	golang.org/x/tools v0.0.0-20200221224223-e1da425f72fd // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/ini.v1 v1.52.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace (
	github.com/XcXerxes/go-blog-server/conf => /Users/zhangjie/xinbo/myself/go/go-blog-server/conf
	github.com/XcXerxes/go-blog-server/docs => /Users/zhangjie/xinbo/myself/go/go-blog-server/docs
	github.com/XcXerxes/go-blog-server/middleware => /Users/zhangjie/xinbo/myself/go/go-blog-server/middleware
	github.com/XcXerxes/go-blog-server/models => /Users/zhangjie/xinbo/myself/go/go-blog-server/models
	github.com/XcXerxes/go-blog-server/pkg/e => /Users/zhangjie/xinbo/myself/go/go-blog-server/pkg/e
	github.com/XcXerxes/go-blog-server/pkg/setting => /Users/zhangjie/xinbo/myself/go/go-blog-server/pkg/setting
	github.com/XcXerxes/go-blog-server/pkg/util => /Users/zhangjie/xinbo/myself/go/go-blog-server/pkg/util
	github.com/XcXerxes/go-blog-server/routers => /Users/zhangjie/xinbo/myself/go/go-blog-server/routers
	github.com/XcXerxes/go-blog-server/runtime => /Users/zhangjie/xinbo/myself/go/go-blog-server/runtime
	github.com/XcXerxes/go-blog-server/service => /Users/zhangjie/xinbo/myself/go/go-blog-server/service
)
