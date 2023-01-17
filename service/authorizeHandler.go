package service

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey    = "authorization"
	authorizationBearerPrefix = "Bearer "
	authorizationPayloadKey   = "authorization_payload"
)

type auther interface {
	CreateToken() (string, error)
	VerifyToken(token string) (interface{}, error)
}

type Auth struct {
	auth auther
}

func NewAuthHandler(auth auther) *Auth {
	return &Auth{auth: auth}
}

func (a *Auth) AccessToken(c Context) {
	fmt.Println("AccessToken Comeing")
	token, err := a.auth.CreateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}

func (a *Auth) Authorization(c Context) {
	fmt.Println("Authorization Comeing")
	auth := c.GetHeader(authorizationHeaderKey)
	tokenString := strings.TrimPrefix(auth, authorizationBearerPrefix)

	_, err := a.auth.VerifyToken(tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
