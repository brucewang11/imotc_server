package controller

import (
	"github.com/brucewang11/frame/internal/service"
	"github.com/brucewang11/frame/internal/vo"
	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	var vo vo.OrderVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ResData(&service.CodeModel{Code:1001},ctx)
	} else {
		userId,_ := ctx.Get("user_id")
		account,_:= ctx.Get("account")

		vo.UserId = userId.(string)
		vo.UserName = account.(string)
		res := service.CreateOrder(&vo)
		ResData(res,ctx)
	}
}


func OrderList(ctx *gin.Context) {
	var vo vo.OrderVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ResData(&service.CodeModel{Code:1001},ctx)
	} else {
		userId,_ := ctx.Get("user_id")
		account,_:= ctx.Get("account")

		vo.UserId = userId.(string)
		vo.UserName = account.(string)
		res := service.OrderList(&vo)
		ResData(res,ctx)
	}
}

func EditOrder(ctx *gin.Context) {
	var vo vo.OrderVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ResData(&service.CodeModel{Code:1001},ctx)
	} else {
		res := service.EditOrder(&vo)
		ResData(res,ctx)
	}
}

func OrderStatus(ctx *gin.Context) {
	var vo vo.OrderVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ResData(&service.CodeModel{Code:1001},ctx)
	} else {
		res := service.OrderStatus(&vo)
		ResData(res,ctx)
	}
}

func OrderInfo(ctx *gin.Context) {
	var vo vo.OrderVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ResData(&service.CodeModel{Code:1001},ctx)
	} else {
		res := service.OrderInfo(&vo)
		ResData(res,ctx)
	}
}

func OrderListAdmin(ctx *gin.Context) {
	var vo vo.OrderVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ResData(&service.CodeModel{Code:1001},ctx)
	} else {
		res := service.OrderListAdmin(&vo)
		ResData(res,ctx)
	}
}


