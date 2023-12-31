package utils

import gonanoid "github.com/matoous/go-nanoid/v2"

const (
	// defaultAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	defaultAlphabet = "23456789ABCDEFGHJKLMNPQRSTWXYZabcdefghijkmnopqrstuvwxyz" // remove resemble letters
	numberAlphabet  = "0123456789"
	defaultSize     = 11
)

// NanoString -  generate optional length nanoid, use const by default
func NanoString(l ...int) string {
	size := defaultSize
	if len(l) > 0 {
		size = l[0]
	}
	return gonanoid.MustGenerate(defaultAlphabet, size)
}

// NanoNumber -  generate optional length nanoid, use const by default
func NanoNumber(l ...int) string {
	size := (defaultSize + 1) / 2
	if len(l) > 0 {
		size = l[0]
	}
	return gonanoid.MustGenerate(numberAlphabet, size)
}

// PrimaryKey - generate primary key
func PrimaryKey(l ...int) func() string {
	size := defaultSize
	if len(l) > 0 {
		size = l[0]
	}
	return func() string {
		return NanoString(size)
	}
}

// IsPrimaryKey - verify is primary key
func IsPrimaryKey(id string) bool {
	if IsEmpty(id) {
		return false
	}
	size := defaultSize
	return len(id) == size
}
