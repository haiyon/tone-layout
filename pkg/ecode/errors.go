package ecode

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

// Newf - returns an error with the format f and args.
func Newf(code int, reason, f string, a ...any) *errors.Error {
	return errors.New(code, reason, fmt.Sprintf(f, a...))
}
