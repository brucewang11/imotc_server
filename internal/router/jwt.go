package router

import (
	log "github.com/alecthomas/log4go"
	"github.com/brucewang11/frame/internal/controller"
	"github.com/brucewang11/frame/internal/service"
	"github.com/brucewang11/frame/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

const TOKEN_EXP = 24 // token有效周期

func JWTAuth(userType int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		//非法token定义


		msgtoken := service.CodeModel{Code: 2000}
		token,_ = c.Cookie("token")
		if token == ""{
			if strings.ToUpper(c.Request.Method) == "GET" {
				token = c.Query("token")
			}
			if strings.ToUpper(c.Request.Method) == "POST" {
				token = c.PostForm("token")
			}
		}


		if token == "" {
			log.Error("JWTAuth token is null")
			c.Abort()
			controller.ResData(&msgtoken,c)
			return
		}

		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if claims == nil {
			log.Error("JWTAuth token parse error")
			c.Abort()
			controller.ResData(&msgtoken,c)
			return
		}
		if err != nil {
			c.Abort()
			log.Error("token parse error", err.Error())
			controller.ResData(&msgtoken,c)
			return
		}
		if userType != claims.UserType {
			log.Error("JWTAuth token is error")
			c.Abort()
			controller.ResData(&msgtoken,c)
			return
		}
		c.Set("user_id", claims.ID)
		c.Set("account",claims.Account)
	}

}



