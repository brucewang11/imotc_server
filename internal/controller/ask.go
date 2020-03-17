package controller

import (
	"github.com/brucewang11/frame/internal/service"
	"github.com/brucewang11/frame/internal/vo"
	"github.com/gin-gonic/gin"
)

func AskList(ctx *gin.Context) {

	res := service.AskList()
	ResData(res,ctx)
}

func EditAsk(ctx *gin.Context) {
	var vo vo.AskVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ResData(&service.CodeModel{Code:1001},ctx)
	} else {
		res := service.EditAsk(&vo)
		ResData(res,ctx)
	}
}
