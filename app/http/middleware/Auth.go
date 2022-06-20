package middleware

import (
	"github.com/gin-gonic/gin"
	"linkr-frame/app/code"
	"linkr-frame/app/throw"
)

//鉴权
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := getToken(ctx)
		if token == "" {
			//token 不存在提示授权登录
			throw.Exception(ctx, code.AuthFailed)
			ctx.Abort()
		}
		//自定义解析token

		ctx.Next()
	}
}

//获取token
func getToken(ctx *gin.Context) string {
	token := ""
	token = ctx.Request.Header.Get("token")
	if token == "" {
		//get token
		token = ctx.Query("token")
	}
	if token == "" {
		//post token
		token = ctx.Request.Form.Get("token")
	}

	return token
}
