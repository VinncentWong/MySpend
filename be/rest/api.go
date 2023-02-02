package rest

import (
	bHandler "module/app/budget/handler"
	fHandler "module/app/financial_account/handler"
	pHandler "module/app/payment/handler"
	"module/app/user/handler"
	"module/middleware"
	"module/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHealthRouting(route *gin.Engine) {
	route.GET("/check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	route.GET("/refresh", util.GenerateNewAccessToken)
}

func UserRouting(route *gin.Engine, handler handler.IUserHandler) {
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

func BudgetRouting(route *gin.Engine, handler *bHandler.BudgetHandler) {
	group := route.Group("/budget")
	group.POST("/create", middleware.JwtMiddleware(), middleware.PremiumMiddleware(), handler.CreateBudget)
	group.GET("/get", middleware.JwtMiddleware(), middleware.PremiumMiddleware(), handler.GetBudget)
	group.PATCH("/update", middleware.JwtMiddleware(), middleware.PremiumMiddleware(), handler.UpdateBudget)
}
