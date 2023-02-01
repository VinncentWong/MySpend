package middleware

import (
	"module/util"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if !strings.HasPrefix(token, "Bearer ") {
			util.SendResponse(c, http.StatusForbidden, "unauthorized", false, nil)
			c.Abort()
			return
		} else {
			token = token[7:]
			fixToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("secret_key")), nil
			})
			if err != nil {
				util.SendResponse(c, http.StatusForbidden, err.Error(), false, nil)
				c.Abort()
				return
			}
			_, ok := fixToken.Claims.(jwt.MapClaims)
			if !ok {
				util.SendResponse(c, http.StatusForbidden, "wrong claims", false, nil)
				c.Abort()
				return
			} else {
				c.Next()
			}
		}
	}
}
