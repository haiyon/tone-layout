package server

import (
	"context"
	"sample/helper/types"
)

// matchWhiteList - operation white list.
func matchWhiteList(_ context.Context, operation string) bool {
	whiteList := make(types.JSON)

	whiteList["/sample.v1.Post/GetPost"] = types.JSON{}
	whiteList["/sample.v1.Post/ListPosts"] = types.JSON{}

	if _, ok := whiteList[operation]; ok {
		return false
	}
	return true
}
