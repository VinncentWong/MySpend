package rest

import (
	fHandler "module/app/financial_account/handler"
	pHandler "module/app/payment/handler"
	"module/app/user/handler"
	"module/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHealthRouting(route *gin.Engine) {
	route.GET("/check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
}

func UserRouting(route *gin.Engine, handler *handler.UserHandler) {
	group := route.Group("/user")
	group.POST("/register", handler.CreateUser)
	group.POST("/login", handler.Login)
	group.GET("/get/:id", middleware.JwtMiddleware(), handler.GetUserById)
}

func FinancialAccountRouting(route *gin.Engine, handler *fHandler.FinancialAccountHandler) {
	group := route.Group("/financialaccount")
	group.POST("/create/:id", middleware.JwtMiddleware(), handler.CreateFinancialAccount)
	group.GET("/get/:id", middleware.JwtMiddleware(), handler.GetFinancialAccount)
}

func PaymentHistoryRouting(route *gin.Engine, handler *pHandler.PaymentHandler) {
	group := route.Group("/payment")
	group.POST("/create/:id", middleware.JwtMiddleware(), handler.CreatePaymentHistory)
}
