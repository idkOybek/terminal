// pkg/auth/jwt.go

package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/idkOybek/newNewTerminal/internal/config"
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID int) (string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", err
	}

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

func ValidateToken(tokenString string) (*Claims, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
