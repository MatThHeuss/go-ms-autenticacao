package middlewares

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func JWTMiddleware(next http.Handler) http.Handler {
	logger, _ := zap.NewProduction()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			secret := []byte("my_secret_key")
			return secret, nil
		})
		if err != nil {
			logger.Error("Error:", zap.Error(err))
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			logger.Error("Error:", zap.Error(err))
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		fmt.Println(claims.GetExpirationTime())
		next.ServeHTTP(w, r)
	})

}
