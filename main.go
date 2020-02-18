package main

import "github.com/gin-gonic/gin"

func steupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello world!")
	})
	return r
}

func main() {
	r := steupRouter()
	r.Run()
}
