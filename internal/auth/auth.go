package auth

import (
	"crypto/ecdsa"
	"github.com/MatThHeuss/go-user-microservice/internal/domain"
	"github.com/golang-jwt/jwt/v5"
)

var (
	key *ecdsa.PrivateKey
)

func CreateToken(user *domain.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss":   "auth-server",
		"sub":   user.ID,
		"email": user.Email,
		"role":  user.Role,
		"name":  user.Name,
	})

	tokenString, _ := token.SignedString("my_secret_Key")

	return tokenString
}
