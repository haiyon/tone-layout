package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sample/helper/utils"
	"sample/internal/data/ent"
	"sample/internal/data/ent/greeter"
	"sample/internal/service"
	_struct "sample/internal/structs"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	_               service.GreeterRepo = (*greeterRepo)(nil)
	greeterCacheKey                     = "greeter:%s"
)

type greeterRepo struct {
	d   *Data
	log *log.Helper
}

// NewGreeterRepo - new greeter repository.
func NewGreeterRepo(d *Data, logger log.Logger) service.GreeterRepo {
	return &greeterRepo{
		d:   d,
		log: log.NewHelper(logger),
	}
}

// CreateGreeter - create greeter.
func (r *greeterRepo) CreateGreeter(ctx context.Context, g *_struct.Greeter) (*ent.Greeter, error) {
	// create builder.
	builder := r.d.ec.Greeter.Create()

	// set related fields.
	// builder.SetUser(uid)

	// set the fields.
	builder.SetName(g.Name)

	// execute the builder.
	row, err := builder.Save(ctx)
	if utils.IsNotNil(err) {
		return nil, err
	}
	return row, nil
}

// GetGreeter - get greeter.
func (r *greeterRepo) GetGreeter(ctx context.Context, p *_struct.FindGreeter) (*ent.Greeter, error) {
	// try to fetch from cache.
	ck := r.d.getCacheKey(fmt.Sprintf(greeterCacheKey, p.Greeter))
	row, err := r.getFromCache(ctx, ck)
	if utils.IsNotNil(err) {
		// fetch from db when cache is empty.
		// use internal get method.
		row, err = r.getGreeter(ctx, p)
		if utils.IsNotNil(err) {
			return nil, err
		}
	}

	return row, err
}

// UpdateGreeter - update greeter.
func (r *greeterRepo) UpdateGreeter(ctx context.Context, body *_struct.Greeter) (*ent.Greeter, error) {
	// query the greeter.
	row, err := r.getGreeter(ctx, &_struct.FindGreeter{
		Greeter: body.ID,
	})

	// create builder.
	builder := row.Update()

	// set the fields.
	builder.SetName("Sample")

	// execute the builder.
	row, err = builder.Save(ctx)
	if utils.IsNotNil(err) {
		return nil, err
	}

	// reset the cache.
	r.resetCache(ctx, row)

	return row, err
}

// DeleteGreeter - delete greeter.
func (r *greeterRepo) DeleteGreeter(ctx context.Context, p *_struct.FindGreeter) error {
	// create builder.
	builder := r.d.ec.Greeter.Delete()

	// match user id
	builder.Where(greeter.IDEQ(p.Greeter))

	// execute the builder.
	_, err := builder.Exec(ctx)
	if utils.IsNotNil(err) {
		return err
	}

	// delete the cache.
	ck := r.d.getCacheKey(fmt.Sprintf(greeterCacheKey, p.Greeter))
	r.d.deleteCache(ctx, ck)

	return nil
}

// ListGreeters - get list of greeters.
func (r *greeterRepo) ListGreeters(ctx context.Context, p *_struct.ListGreeters) ([]*ent.Greeter, error) {
	// create list builder
	builder, err := r.listBuilder(ctx, p)
	if utils.IsNotNil(err) {
		return nil, err
	}

	// limit the result
	builder.Limit(int(p.Limit))

	// order by create time desc
	builder.Order(ent.Desc(greeter.FieldCreatedAt))

	// execute the builder.
	greeters, err := builder.All(ctx)
	if utils.IsNotNil(err) {
		return nil, err
	}

	return greeters, nil
}

// CountX - taxonnomy count.
func (r *greeterRepo) CountX(ctx context.Context, p *_struct.ListGreeters) int {
	// create list builder
	builder, err := r.listBuilder(ctx, p)
	if utils.IsNotNil(err) {
		return 0
	}
	return builder.CountX(ctx)
}

// listBuilder - create list builder.
// internal method.
func (r *greeterRepo) listBuilder(ctx context.Context, p *_struct.ListGreeters) (*ent.GreeterQuery, error) {
	// verify query params.
	var nextGreeter *ent.Greeter
	if utils.IsNotEmpty(p.Cursor) {
		// query the greeter.
		// use internal get method.
		row, err := r.getGreeter(ctx, &_struct.FindGreeter{Greeter: p.Cursor})
		if utils.IsNotNil(err) || utils.IsNil(row) {
			return nil, errors.New(service.ErrInvalidCursor)
		}
		nextGreeter = row
	}
	// create builder.
	builder := r.d.ec.Greeter.Query()

	// lt the cursor create time
	if nextGreeter != nil {
		builder.Where(greeter.CreatedAtLT(nextGreeter.CreatedAt))
	}

	// match user id
	// if utils.IsNotEmpty(p.User) {
	// 	builder.Where(greeter.UserIDEQ(p.User))
	// }

	return builder, nil
}

// getGreeter - get greeter.
// internal method.
func (r *greeterRepo) getGreeter(ctx context.Context, p *_struct.FindGreeter) (*ent.Greeter, error) {
	// create builder.
	builder := r.d.ec.Greeter.Query()

	// set the query conditions.
	if utils.IsNotEmpty(p.Greeter) {
		builder.Where(greeter.IDEQ(p.Greeter))
	}
	// match user id.
	// if utils.IsNotEmpty(p.User) {
	// 	builder.Where(greeter.UserIDEQ(p.User))
	// }

	// execute the builder.
	row, err := builder.First(ctx)
	if utils.IsNotNil(err) {
		return nil, err
	}

	return row, err
}

// getFromCache - get greeter from cache.
func (r *greeterRepo) getFromCache(ctx context.Context, key string) (*ent.Greeter, error) {
	result, err := r.d.rc.Get(ctx, key).Result()
	if utils.IsNotNil(err) {
		return nil, err
	}
	var row = &ent.Greeter{}
	err = json.Unmarshal([]byte(result), row)
	if utils.IsNotNil(err) {
		return nil, err
	}
	return row, nil
}

// setCache - set greeter to cache.
func (r *greeterRepo) setCache(ctx context.Context, row *ent.Greeter, key string) {
	marshal, err := json.Marshal(row)
	if utils.IsNotNil(err) {
		r.log.Errorf("fail to set greeter cache:json.Marshal(%v) error(%v)", row, err)
	}
	err = r.d.rc.Set(ctx, key, string(marshal), time.Minute*30).Err()
	if utils.IsNotNil(err) {
		r.log.Errorf("fail to set greeter cache:redis.Set(%v) error(%v)", row, err)
	}
}

// resetCache - reset greeter to cache.
func (r *greeterRepo) resetCache(ctx context.Context, row *ent.Greeter) {
	// get the cache key.
	ck := r.d.getCacheKey(fmt.Sprintf(greeterCacheKey, row.ID))

	// delete the cache.
	// r.d.deleteCache(ctx, ck)

	// set the cache.
	r.setCache(ctx, row, ck)
}
