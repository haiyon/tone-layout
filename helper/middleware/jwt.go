package middleware

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
)

// JWTServer - service server authorize middleware.
func JWTServer(key string) middleware.Middleware {
	return jwt.Server(func(token *jwtV4.Token) (any, error) {
		return []byte(key), nil
	}, jwt.WithSigningMethod(jwtV4.SigningMethodHS256))
}

// JWTClient - service client authorize middleware.
func JWTClient(key string) middleware.Middleware {
	return jwt.Client(func(token *jwtV4.Token) (any, error) {
		return []byte(key), nil
	}, jwt.WithSigningMethod(jwtV4.SigningMethodHS256))
}
