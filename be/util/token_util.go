package util

import (
	"context"
	"fmt"
	"module/domain"
	"module/infrastructure"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateNewAccessToken(c *gin.Context) {
	token := c.Request.Header.Get("access-token")
	refreshToken := c.Request.Header.Get("refresh-token")

	if strings.HasPrefix(token, "Bearer ") {
		SendResponse(c, http.StatusBadRequest, "invalid token", false, nil)
		return
	}
	jwtToken, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("secret_key")), nil
	})
	if err != nil {
		SendResponse(c, http.StatusBadRequest, err.Error(), false, nil)
		return
	}
	if jwtToken.Valid {
		SendResponse(c, http.StatusBadRequest, "token still valid", false, nil)
		return
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		SendResponse(c, http.StatusInternalServerError, "wrong claims", false, nil)
		return
	}
	id := claims["id"].(float64)
	name := claims["name"].(string)
	email := claims["email"].(string)
	user := domain.User{
		Name:  name,
		Email: email,
	}
	newToken, err := CreateToken(user)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	redisClient := infrastructure.GetRedisClient()
	err = redisClient.Set(context.Background(), fmt.Sprintf("jwt_token/%v", id), *newToken, 15*time.Minute).Err()
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	SendResponse(c, http.StatusOK, "success create new token", true, newToken)
}
