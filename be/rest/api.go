package rest

import (
	"module/app/user/handler"

	"github.com/gin-gonic/gin"
)

func UserRouting(route *gin.Engine, handler *handler.UserHandler) {
	group := route.Group("/user")
	group.POST("/register", handler.CreateUser)
	group.POST("/login", handler.Login)
}
