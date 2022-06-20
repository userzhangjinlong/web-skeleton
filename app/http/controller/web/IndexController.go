package web

import (
	"github.com/gin-gonic/gin"
	"linkr-frame/app/code"
	"linkr-frame/app/http/controller"
	"linkr-frame/app/logic/web"
	"linkr-frame/app/throw"
)

type IndexController struct {
	Base       *controller.BaseController
	IndexLogic *web.IndexLogic
}

//控制器层
func (i *IndexController) Index(ctx *gin.Context) {
	throw.Success(ctx, code.SuccessCode, i.IndexLogic.Index())
}
