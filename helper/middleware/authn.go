package middleware

import (
	"context"
	"net/http"
	"time"

	"sample/helper/constant"
	"sample/helper/ecode"
	"sample/helper/meta"
	"sample/helper/types"
	"sample/helper/utils"

	"github.com/go-kratos/kratos/v2/middleware"
)

var (
	err      error
	token    string
	user     string
	username string
)

// refresh - Refresh token
// func refresh(refreshToken string) (string, error) {
// 	return "", nil
// }

// Authn - verify user and append to context
func Authn(secret ...string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			if len(secret) > 0 {
				// get token form metadata or context
				token, err = meta.GetToken(ctx)
				if utils.IsNotNil(err) {
					return nil, ecode.Unauthorized(http.StatusText(http.StatusUnauthorized), err.Error())
				}

				tokenData, err := utils.DecodeToken(secret[0], token)
				if utils.IsNotNil(err) {
					return nil, ecode.Unauthorized(http.StatusText(http.StatusUnauthorized), err.Error())
				}

				// check token expired
				tokenExpire := int64(tokenData["expire"].(float64))
				now := time.Now().Unix()
				diff := tokenExpire - now
				// token is expired
				if diff < 0 {
					return nil, ecode.Unauthorized(http.StatusText(http.StatusUnauthorized), "token is expired")
				}
				payload := tokenData["payload"].(types.JSON)
				user = payload["user_id"].(string)
				username = payload["username"].(string)
			} else {
				// get user id from context
				user, err = meta.GetUser(ctx)
				if utils.IsNotNil(err) {
					return nil, ecode.Forbidden(http.StatusText(http.StatusForbidden), err.Error())
				}
				// get username from context
				username, err = meta.GetUsername(ctx)
				if utils.IsNotNil(err) {
					return nil, ecode.Forbidden(http.StatusText(http.StatusForbidden), err.Error())
				}
			}

			// set user id to context
			ctx = context.WithValue(ctx, constant.X_MD_USER, user)
			// set username to context
			ctx = context.WithValue(ctx, constant.X_MD_USERNAME, username)

			// fmt.Printf("------ user_id: %s, username: %s\n", user, username)

			return handler(ctx, req)
		}
	}
}

// // setContext - set metadata to context
// func setContext(ctx context.Context) context.Context {}
//
// // setMetadata - broadcast metadata
// func setMetadata(ctx context.Context) context.Context {}
