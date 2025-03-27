package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	AccountID uint64 `json:"account_id"`
	jwt.StandardClaims
}

func GenerateToken(accountID uint64, secretKey string, expirationTime time.Duration) (string, error) {
	claims := &Claims{
		AccountID: accountID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ValidateToken(tokenString string, secretKey string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
