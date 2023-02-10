package gomp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMath(t *testing.T) {
	assert := assert.New(t)
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(0, Min(input...))
	assert.Equal(10, Max(input...))
	assert.Equal(2, Abs(-2))

	assert.Empty(Contains(input, -1))
	assert.Equal(true, Contains(input, 0))
	assert.Equal(true, Contains(input, 5))
	assert.Equal(true, Contains(input, 10))
	assert.NotEqual(true, Contains(input, 100))
}
