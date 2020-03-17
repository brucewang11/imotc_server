package controller

import (
	"github.com/brucewang11/frame/internal/service"
	"github.com/brucewang11/frame/internal/vo"
	"github.com/gin-gonic/gin"
)

func CreateCoinOut(ctx *gin.Context) {
	var vo vo.CoinOutVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ResData(&service.CodeModel{Code:1001},ctx)
	} else {
		userId,_ := ctx.Get("user_id")
		account,_:= ctx.Get("account")

		vo.UserId = userId.(string)
		vo.Account = account.(string)
		res := service.CreateCoinOut(&vo)
		ResData(res,ctx)
	}
}

func CoinOutList(ctx *gin.Context) {
	var vo vo.CoinOutVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ResData(&service.CodeModel{Code:1001},ctx)
	} else {
		userId,_ := ctx.Get("user_id")
		account,_:= ctx.Get("account")

		vo.UserId = userId.(string)
		vo.Account = account.(string)
		res := service.CoinOutList(&vo)
		ResData(res,ctx)
	}
}

func CoinOutListAdmin(ctx *gin.Context) {
	var vo vo.CoinOutVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ResData(&service.CodeModel{Code:1001},ctx)
	} else {
		res := service.CoinOutList(&vo)
		ResData(res,ctx)
	}
}


func EditCoinOut(ctx *gin.Context) {
	var vo vo.CoinOutVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ResData(&service.CodeModel{Code:1001},ctx)
	} else {
		res := service.EditCoinOut(&vo)
		ResData(res,ctx)
	}
}