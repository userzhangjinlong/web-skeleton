package route

import (
	"github.com/gin-gonic/gin"
	"linkr-frame/app/http/controller/web"
	"linkr-frame/app/http/middleware"
)

var (
	webIndex web.IndexController
)

//web 服务路由
func WebRouter(r *gin.Engine) {
	//group1
	route := r.Group("/")
	{
		route.GET("/", webIndex.Index)
	}

	//group2
	userGroup := r.Group("user").
		Use(middleware.Auth())
	userGroup.GET("/info", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "this is user api!"})
	})
}
