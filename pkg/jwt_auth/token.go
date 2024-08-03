package jwt_auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type JwtAuth struct {
	key string
}

func NewJwtAuth(key string) *JwtAuth {
	return &JwtAuth{
		key: key,
	}
}

type TokenClaims struct {
	jwt.StandardClaims
	Login string
}

func (ja *JwtAuth) GenerateToken(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{},
		login,
	})

	return token.SignedString([]byte(ja.key))
}

func (ja *JwtAuth) ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(ja.key), nil
	})
	if err != nil {
		return "", fmt.Errorf("Parse: %w", err)
	}
	return token.Claims.(jwt.MapClaims)["Login"].(string), nil
}
