package infra

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/srvc/fail"

	"github.com/tadoku/api/domain"
	"github.com/tadoku/api/usecases"
)

// ErrEmptyUser when context contains no user when trying to get one
var ErrEmptyUser = fail.Errorf("user is empty")

type context struct {
	echo.Context
}

func (c context) User() (*domain.User, error) {
	claims := c.Claims()
	if claims != nil && claims.User != nil {
		return claims.User, nil
	}

	return nil, ErrEmptyUser
}

func (c context) Claims() *usecases.SessionClaims {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		claims := token.Claims.(*jwtClaims)
		if claims != nil {
			return &claims.SessionClaims
		}
	}
	return nil
}
