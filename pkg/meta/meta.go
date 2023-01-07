package meta

import (
	"context"
	"errors"
	"strings"

	"sample/pkg/constant"
	"sample/pkg/utils"

	"github.com/go-kratos/kratos/v2/transport"

	"github.com/go-kratos/kratos/v2/metadata"
)

const (
	TOKEN_ERROR string = "missing token or token incorrect"
)

// GetToken - get token from context or metadata
func GetToken(ctx context.Context) (string, error) {
	var tokenString string
	// get token from metadata.
	md, _ := metadata.FromServerContext(ctx)
	tokenString = md.Get(constant.X_MD_TOKEN)
	// if token is empty continue to get token from context
	if utils.IsEmpty(tokenString) {
		tr, _ := transport.FromServerContext(ctx)
		tokenString = tr.RequestHeader().Get(constant.AUTHORIZATION)
	}
	if utils.IsEmpty(tokenString) {
		return "", errors.New(TOKEN_ERROR)
	}
	// format
	// ie Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9
	// b := "Bearer "
	if !strings.Contains(tokenString, constant.BEARER) {
		return "", errors.New(TOKEN_ERROR)
	}
	t := strings.Split(tokenString, constant.BEARER)
	if len(t) < 2 {
		return "", errors.New(TOKEN_ERROR)
	}
	token := t[1]
	return token, nil
}

// GetUser - get user from context or metadata
func GetUser(ctx context.Context) (string, error) {
	// get user from context
	user := ctx.Value(constant.X_MD_USER)
	// user := strings.ReplaceAll(fmt.Sprintf("%v", ctx.Value(constant.X_MD_USER)), "<nil>", "")
	// fmt.Printf("-------------- user is %v, type: %v\n", user, reflect.TypeOf(user))
	if utils.IsNotNil(user) && utils.IsPrimaryKey(user.(string)) {
		return user.(string), nil
	}
	// if user is empty continue to get user from metadata.
	if md, ok := metadata.FromServerContext(ctx); ok {
		if user := md.Get(constant.X_MD_USER); utils.IsNotNil(user) && len(user) > 0 && utils.IsPrimaryKey(user) {
			return user, nil
		}
	}
	return "", errors.New("failed to get user")
}

// GetUsername - get username from context or metadata
func GetUsername(ctx context.Context) (string, error) {
	// get username from context
	username := ctx.Value(constant.X_MD_USERNAME)
	// username := strings.ReplaceAll(fmt.Sprintf("%v", ctx.Value(constant.X_MD_USERNAME)), "<nil>", "")
	if utils.IsNotNil(username) {
		return username.(string), nil
	}
	// if username is empty continue to get username from metadata.
	if md, ok := metadata.FromServerContext(ctx); ok {
		if username := md.Get(constant.X_MD_USERNAME); len(username) > 0 {
			return username, nil
		}
	}
	return "", errors.New("failed to get username")
}
