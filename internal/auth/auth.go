package auth

import (
	"fmt"
	"github.com/MatThHeuss/go-user-microservice/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"time"
)

func CreateToken(user *domain.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":   "auth-server",
		"sub":   user.ID,
		"email": user.Email,
		"role":  user.Role,
		"name":  user.Name,
		"exp":   time.Second * 1000,
	})

	tokenString, err := token.SignedString([]byte("my_secret_key"))
	fmt.Println(tokenString)
	if err != nil {
		zap.Error(err)
		return err.Error()
	}
	return tokenString
}
