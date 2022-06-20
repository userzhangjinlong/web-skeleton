package throw

import (
	"github.com/gin-gonic/gin"
	"linkr-frame/app/code"
)

//响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//公用异常抛出
func Exception(ctx *gin.Context, ECode code.ECode) {
	response := Response{
		Code: ECode.Code,
		Msg:  ECode.Msg,
		Data: nil,
	}
	ctx.JSON(ECode.Code, response)
	return
}

//成功响应
func Success(ctx *gin.Context, ECode code.ECode, data interface{}) {
	response := Response{
		Code: ECode.Code,
		Msg:  ECode.Msg,
		Data: data,
	}

	ctx.JSON(ECode.Code, response)
	return
}
