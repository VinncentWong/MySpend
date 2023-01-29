package util

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(data any) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": data,
	})
	fixedToken, err := token.SignedString([]byte(os.Getenv("secret_key")))
	if err != nil {
		return nil, err
	}
	return &fixedToken, err
}

func ValidateToken(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.Next()
			return
		}
		if !strings.HasPrefix(token, "Bearer ") {
			SendResponse(c, http.StatusForbidden, "unauthorized", false, nil)
			return
		} else {
			fixToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
				return os.Getenv("secret_key"), nil
			})
			if err != nil {
				SendResponse(c, http.StatusForbidden, "unauthorized", false, nil)
				c.Abort()
				return
			}
			_, ok := fixToken.Claims.(jwt.MapClaims)
			if !ok {
				SendResponse(c, http.StatusForbidden, "wrong claims", false, nil)
				c.Abort()
				return
			} else {
				c.Next()
			}
		}
	}
}
