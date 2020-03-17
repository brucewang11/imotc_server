package router

import (
	"github.com/brucewang11/frame/internal/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)



func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())
	app := router.Group("/api/v1")
	app.POST("load_admin",controller.LoadAdmin)
	adminController := app.Group("/admin",JWTAuth(1))
	adminController.GET("/ask_list",controller.AskList)
	adminController.GET("/user_list",controller.UserList)
	adminController.POST("/edit_order",controller.EditOrder)
	adminController.GET("/order_list",controller.OrderListAdmin)
	adminController.POST("/edit_ask",controller.EditAsk)
	adminController.POST("/edit_pwd",controller.EditPwd)
	adminController.GET("/coinout_list",controller.CoinOutListAdmin)
	adminController.POST("/edit_coinout",controller.EditCoinOut)

	accountController := app.Group("/account")
	accountController.POST("/send_code",controller.SendCode)
	accountController.POST("/register",controller.Register)
	accountController.POST("/load",controller.Load)

	askController := app.Group("ask")
	askController.GET("list",controller.AskList)
	app.Use(JWTAuth(2))

	userController := app.Group("/user")
	userController.GET("/info",controller.FindUserById)
	orderController := app.Group("/order")
	orderController.POST("/create",controller.CreateOrder)
	orderController.GET("/list",controller.OrderList)
	orderController.POST("/status",controller.OrderStatus)
	orderController.GET("/info",controller.OrderInfo)

	coinOutController := app.Group("/coinout")
	coinOutController.POST("/create",controller.CreateCoinOut)
	coinOutController.GET("/list",controller.CoinOutList)







	return router
}



