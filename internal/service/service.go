package service

import (
	iV1 "sample/api/interface/v1"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrInvalidID - invalid id
	ErrInvalidID = "the id is invalid"
	// ErrInvalidCursor - invalid cursor.
	ErrInvalidCursor = "cursor is invalid"
	// ErrInvalidLimit - invalid limit.
	ErrInvalidLimit = "limit is invalid"
)

// Service - sample handles.
type Service struct {
	iV1.UnimplementedGreeterServer
	iV1.UnimplementedPostServer

	greeterRepo GreeterRepo

	log *log.Helper
}

// NewService - new sample handles.
func NewService(greeterRepo GreeterRepo, logger log.Logger) *Service {
	logHelper := log.NewHelper(log.With(logger, "module", "sample/handle"))
	return &Service{
		greeterRepo: greeterRepo,
		log:         logHelper,
	}
}
