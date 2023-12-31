package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNanoid(t *testing.T) {
	t.Run("generate string", func(t *testing.T) {
		str := NanoString()
		assert.Len(t, str, 11, "should return str of default length")
	})

	t.Run("generate number", func(t *testing.T) {
		num := NanoNumber()
		assert.Len(t, num, 6, "should return num of defautl length")
	})

	t.Run("generate primary key", func(t *testing.T) {
		pk := PrimaryKey()()
		assert.Len(t, pk, 11, "should return primary key of default length")
	})
}
