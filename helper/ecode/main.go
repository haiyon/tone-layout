package ecode

import "fmt"

// Unknown - unknown error
func Unknown(f string, a ...any) error {
	return InternalServer("Unknown", fmt.Sprintf(f, a...))
}
