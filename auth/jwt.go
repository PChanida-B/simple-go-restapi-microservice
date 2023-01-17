package auth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const maxTimeSession = 10

type JWTToken struct {
	signature string
}

func NewJWTAuth(signature string) *JWTToken {
	return &JWTToken{signature}
}

func (maker *JWTToken) CreateToken() (string, error) {
	payload := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(maxTimeSession * time.Minute).Unix(),
		Audience:  "Chanida",
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.signature))
	return token, err
}

func (maker *JWTToken) VerifyToken(token string) (interface{}, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(maker.signature), nil
	}

	jwtToken, err := jwt.Parse(token, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("ErrInvalidToken")
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("ErrInvalidToken")
	}
	return claims, nil
}
