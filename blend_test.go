package gomp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlend(t *testing.T) {
	assert := assert.New(t)

	op := NewBlend()
	assert.Empty(op.Get())
	op.Set(Darken)
	assert.Equal(Darken, op.Get())
	op.Set(Lighten)
	assert.Equal(Lighten, op.Get())

	rgb := Color{R: 0xff, G: 0xff, B: 0xff}
	lum := op.Lum(rgb)
	assert.Equal(255.0, lum)

	rgb = Color{R: 0, G: 0, B: 0}
	lum = op.Lum(rgb)
	assert.Equal(0, lum)
}
