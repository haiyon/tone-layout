package server

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

// getOperation - get operation form context.
func getOperation(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req any) (reply any, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			fmt.Println(tr.Operation())
		}
		return handler(ctx, req)
	}
}
