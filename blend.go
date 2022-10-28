// Package gomp implements the Porter-Duff composition operations
// used for mixing a graphic element with its backdrop.
// Porter and Duff presented in their paper 12 different composition operation, but the
// core image/draw core package implements only the source-over-destination and source.
// This package implements all of the 12 composite operation together with some blending modes.
package gomp

import "github.com/esimov/gomp/utils"

const (
	Darken     = "darken"
	Lighten    = "lighten"
	Multiply   = "multiply"
	Screen     = "screen"
	Overlay    = "overlay"
	SoftLight  = "soft_light"
	HardLight  = "hard_light"
	ColorDodge = "color_dodge"
	ColorBurn  = "color_burn"
)

// Blend holds the currently active blend mode.
type Blend struct {
	Mode  string
	Modes []string
}

// NewBlend initializes a new Blend.
func NewBlend() *Blend {
	return &Blend{
		Modes: []string{
			Darken,
			Lighten,
			Multiply,
			Screen,
			Overlay,
			SoftLight,
			HardLight,
			ColorDodge,
			ColorBurn,
		},
	}
}

// Set activate one of the supported blend mode.
func (b *Blend) Set(blendType string) {
	if utils.Contains(b.Modes, blendType) {
		b.Mode = blendType
	}
}

// Get returns the currently active blend mode.
func (b *Blend) Get() string {
	if len(b.Mode) > 0 {
		return b.Mode
	}
	return ""
}
