package service

import (
	"context"
	mV1 "sample/api/schema/v1"
	sV1 "sample/api/shared/v1"
	"sample/internal/data/ent"
	structs "sample/internal/structs"
	"sample/pkg/types"
	"sample/pkg/utils"
)

var (
	// ErrGreeterAlreadyExist - greeter already exist
	ErrGreeterAlreadyExist = "greeter already exist"
	// ErrGreeterNotExist - greeter is not exist
	ErrGreeterNotExist = "greeter is not exist"
)

// GreeterRepo - greeter repository interface.
type GreeterRepo interface {
	// CreateGreeter - create greeter.
	CreateGreeter(ctx context.Context, body *structs.Greeter) (*ent.Greeter, error)
	// GetGreeter - get greeter.
	GetGreeter(ctx context.Context, p *structs.FindGreeter) (*ent.Greeter, error)
	// UpdateGreeter - update greeter.
	UpdateGreeter(ctx context.Context, body *structs.Greeter) (*ent.Greeter, error)
	// DeleteGreeter - delete greeter.
	DeleteGreeter(ctx context.Context, p *structs.FindGreeter) error
	// ListGreeters - get list of greeter.
	ListGreeters(ctx context.Context, p *structs.ListGreeters) ([]*ent.Greeter, error)
}

// CreateGreeter - create greeter
func (s *Service) CreateGreeter(ctx context.Context, req *mV1.GreeterRequest) (*mV1.GreeterReply, error) {
	greeter, err := s.greeterRepo.CreateGreeter(ctx, s.serializeGreeterStruct(req))
	if ent.IsConstraintError(err) {
		return nil, sV1.ErrorConflict(ErrGreeterAlreadyExist)
	} else if utils.IsNotNil(err) {
		return nil, sV1.ErrorInternalServer(err.Error())
	}
	return s.serializeGreeterReply(greeter), nil
}

// GetGreeter - get greeter
func (s *Service) GetGreeter(ctx context.Context, req *mV1.GetGreeterRequest) (*mV1.GreeterReply, error) {
	greeter, err := s.greeterRepo.GetGreeter(ctx, s.serializeGetGreeterStruct(req))
	if ent.IsNotFound(err) {
		return nil, sV1.ErrorNotFound(ErrGreeterNotExist)
	} else if utils.IsNotNil(err) {
		return nil, sV1.ErrorInternalServer(err.Error())
	}

	return s.serializeGreeterReply(greeter), nil
}

// UpdateGreeter - update greeter.
func (s *Service) UpdateGreeter(ctx context.Context, req *mV1.GreeterRequest) (*mV1.GreeterReply, error) {
	greeter, err := s.greeterRepo.UpdateGreeter(ctx, s.serializeGreeterStruct(req))
	if ent.IsNotFound(err) {
		return nil, sV1.ErrorNotFound(ErrGreeterNotExist)
	} else if ent.IsConstraintError(err) {
		return nil, sV1.ErrorConflict(ErrGreeterAlreadyExist)
	} else if utils.IsNotNil(err) {
		return nil, sV1.ErrorInternalServer(err.Error())
	}

	return s.serializeGreeterReply(greeter), nil
}

// DeleteGreeter - delete user greeter.
func (s *Service) DeleteGreeter(ctx context.Context, req *mV1.GetGreeterRequest) (*sV1.Response, error) {
	err := s.greeterRepo.DeleteGreeter(ctx, s.serializeGetGreeterStruct(req))
	if ent.IsNotFound(err) {
		return nil, sV1.ErrorNotFound(ErrGreeterNotExist)
	} else if utils.IsNotNil(err) {
		return nil, sV1.ErrorInternalServer(err.Error())
	}

	return &sV1.Response{
		Success: utils.BoolPointer(utils.IsNil(err)),
	}, nil
}

// ListGreeters - get list of greeter.
func (s *Service) ListGreeters(ctx context.Context, req *mV1.ListGreetersRequest) (*mV1.ListGreetersReply, error) {
	// limit default value
	if utils.IsEmpty(req.Limit) {
		req.Limit = 20
	}
	// limit must less than 100
	if req.Limit > 100 {
		return nil, sV1.ErrorBadRequest(ErrInvalidLimit)
	}

	// execute the repository method.
	greeters, err := s.greeterRepo.ListGreeters(ctx, &structs.ListGreeters{
		Cursor: req.Cursor,
		Limit:  req.Limit,
	})
	if ent.IsNotFound(err) {
		return nil, sV1.ErrorNotFound(ErrInvalidCursor)
	} else if utils.IsNotNil(err) {
		return nil, sV1.ErrorInternalServer(err.Error())
	}

	return s.serializeListGreetersReply(greeters), nil
}

// serializeGreeterStruct - convert card create or update protobuf struct to dto struct.
// internal method.
func (s *Service) serializeGreeterStruct(req *mV1.GreeterRequest) *structs.Greeter {
	return &structs.Greeter{}
}

// serializeGetGreeterStruct - convert card query protobuf struct to dto struct.
// internal method.
func (s *Service) serializeGetGreeterStruct(req *mV1.GetGreeterRequest) *structs.FindGreeter {
	return &structs.FindGreeter{
		Greeter: req.Greeter,
		User:    req.User,
	}
}

// serializeListGreeters - serialize greeter list.
func (s *Service) serializeListGreetersReply(greeters []*ent.Greeter) *mV1.ListGreetersReply {
	rs := make([]*mV1.GreeterReply, 0, len(greeters))
	for _, greeter := range greeters {
		rs = append(rs, s.serializeGreeterReply(greeter))
	}
	return &mV1.ListGreetersReply{
		Content: rs,
	}
}

// serializeGreeter - serialize greeter.
func (s *Service) serializeGreeterReply(greeter *ent.Greeter) *mV1.GreeterReply {
	rv := &mV1.GreeterReply{
		Id:        greeter.ID,
		Name:      greeter.Name,
		CreatedAt: types.ToPBTimestamp(greeter.CreatedAt),
		UpdatedAt: types.ToPBTimestamp(greeter.UpdatedAt),
	}
	return rv
}
