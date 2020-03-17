package controller

import (
	"github.com/brucewang11/frame/internal/service"
	"github.com/brucewang11/frame/internal/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendCode(ctx *gin.Context) {
	var vo vo.UserVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"Code": 1001, "Msg":"参数错误"})
	} else {
		res := service.SendCode(&vo)
		ResData(res,ctx)
	}

}

func Register(ctx *gin.Context) {
	var vo vo.UserVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"Code": 1001, "Msg":"参数错误"})
	} else {
		res := service.Register(&vo)
		ResData(res,ctx)
	}

}


func Load(ctx *gin.Context) {
	var vo vo.UserVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"Code": 1001, "Msg":"参数错误"})
	} else {
		res := service.Load(&vo)
		ResData(res,ctx)
	}

}

func LoadAdmin(ctx *gin.Context) {
	var vo vo.UserVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"Code": 1001, "Msg":"参数错误"})
	} else {
		res := service.LoadAdmin(&vo)
		ResData(res,ctx)
	}

}


func UserList(ctx *gin.Context) {
	var vo vo.UserVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"Code": 1001, "Msg":"参数错误"})
	} else {
		res := service.UserList(&vo)
		ResData(res,ctx)
	}

}

func EditPwd(ctx *gin.Context) {
	var vo vo.UserVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"Code": 1001, "Msg":"参数错误"})
	} else {
		res := service.EditPwd(&vo)
		ResData(res,ctx)
	}

}

func FindUserById(ctx *gin.Context) {
	var vo vo.UserVo
	if err := ctx.ShouldBind(&vo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"Code": 1001, "Msg":"参数错误"})
	} else {
		userId,_ := ctx.Get("user_id")

		vo.UserId = userId.(string)
		res := service.FindUserById(&vo)
		ResData(res,ctx)
	}

}




