package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"linkr-frame/app/code"
	"linkr-frame/app/throw"
	"linkr-frame/app/utils"
)

//全局panic异常捕获处理
func ThrowErr() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//系统级别异常捕获
		defer func() {
			if err := recover(); err != nil {
				//记录异常调用栈信息
				logrus.WithFields(logrus.Fields{
					"stack": utils.PanicTrace(err),
					"err":   err,
				}).Error("全局failed error")
				//系统级别panic捕获防止程序崩溃
				throw.Exception(ctx, code.SystemError)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
