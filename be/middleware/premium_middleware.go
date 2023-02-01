package middleware

import (
	"module/domain"
	"module/infrastructure"
	"module/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PremiumMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("userId")
		db := infrastructure.GetDb()
		var container domain.User
		result := db.Where("id = ?", id).Take(&container)
		if result.Error != nil {
			c.Abort()
			util.SendResponse(c, http.StatusBadRequest, "user doesn't exist!", false, nil)
			return
		}
		if !container.IsPremium && container.PremiumExpire.After(time.Now()) {
			container.IsPremium = false
			c.Abort()
			util.SendResponse(c, http.StatusForbidden, "user premium duration is expire or user is not on premium state", false, nil)
			return
		}
		c.Next()
	}
}
