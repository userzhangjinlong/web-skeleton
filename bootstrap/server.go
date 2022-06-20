package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"linkr-frame/app/http/middleware"
	"linkr-frame/global"
	"linkr-frame/route"
)

//初始化配置
func init() {
	InitSystemConfig()
	InitLogConfig()
	InitMysql()
	InitRedis()
}

//web服务启用
func Start() {
	engine := gin.Default()
	//定义全局中间件
	engine.
		Use(middleware.Cors()).
		Use(middleware.ThrowErr())

	route.WebRouter(engine)

	engine.Run(fmt.Sprintf(":%s", global.Config.Frame.Port))
}
