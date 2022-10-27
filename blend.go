// Package gomp implements the Porter-Duff composition operations
// used for mixing a graphic element with its backdrop.
// Porter and Duff presented in their paper 12 different composition operation, but the
// core image/draw core package implements only the source-over-destination and source.
// This package implements all of the 12 composite operation together with some blending modes.
package gomp

import "github.com/esimov/gomp/utils"

const (
	Darken   = "darken"
	Lighten  = "lighten"
	Multiply = "multiply"
	Screen   = "screen"
	Overlay  = "overlay"
)

// Blend holds the currently active blend mode.
type Blend struct {
	OpType string
}

// NewBlend initializes a new Blend.
func NewBlend() *Blend {
	return &Blend{}
}

// Set activate one of the supported blend mode.
func (o *Blend) Set(opType string) {
	bModes := []string{Darken, Lighten, Multiply, Screen, Overlay}

	if utils.Contains(bModes, opType) {
		o.OpType = opType
	}
}

// Get returns the currently active blend mode.
func (o *Blend) Get() string {
	if len(o.OpType) > 0 {
		return o.OpType
	}
	return ""
}
