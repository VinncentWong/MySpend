package util

import (
	"context"
	"fmt"
	"log"
	"module/domain"
	"module/infrastructure"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func GenerateNewAccessToken(c *gin.Context) {
	token := c.Request.Header.Get("access-token")
	refreshToken := c.Request.Header.Get("refresh-token")
	accessToken, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("secret_key")), nil
	})
	exp := accessToken.Claims.(jwt.MapClaims)["exp"].(float64)
	log.Default().Printf("now = %v, exp = %v", time.Now().Unix(), int64(exp))
	if int64(exp) > time.Now().Unix() {
		SendResponse(c, http.StatusUnauthorized, "token still valid", false, nil)
		return
	}
	jwtToken, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("secret_key")), nil
	})
	if err != nil {
		log.Default().Printf("refresh token with error %v", err.Error())
		if err.Error() != "Token is expired" {
			SendResponse(c, http.StatusUnauthorized, err.Error(), false, nil)
			return
		}
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
		Model: gorm.Model{
			ID: uint(id),
		},
		Name:  name,
		Email: email,
	}
	log.Default().Printf("id = %v", id)
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
