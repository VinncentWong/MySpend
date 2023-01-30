package util

import (
	"os"

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
