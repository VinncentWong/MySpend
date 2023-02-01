package util

import (
	"module/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(data domain.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        data.ID,
		"createdAt": data.CreatedAt,
		"updatedAt": data.UpdatedAt,
		"email":     data.Email,
		"name":      data.Name,
		"exp":       time.Now().Add(time.Minute * 15).Unix(),
	})
	fixedToken, err := token.SignedString([]byte(os.Getenv("secret_key")))
	if err != nil {
		return nil, err
	}
	return &fixedToken, err
}
