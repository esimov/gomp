// Package gomp implements the Porter-Duff composition operations
// used for mixing a graphic element with its backdrop.
// Porter and Duff presented in their paper 12 different composition operation, but the
// core image/draw core package implements only the source-over-destination and source.
// This package implements all of the 12 composite operation together with some blending modes.
package gomp

import (
	"sort"

	"github.com/esimov/gomp/math"
)

const (
	Normal     = "normal"
	Darken     = "darken"
	Lighten    = "lighten"
	Multiply   = "multiply"
	Screen     = "screen"
	Overlay    = "overlay"
	SoftLight  = "soft_light"
	HardLight  = "hard_light"
	ColorDodge = "color_dodge"
	ColorBurn  = "color_burn"
	Difference = "difference"
	Exclusion  = "exclusion"

	// Non-separable blend modes
	Hue        = "hue"
	Saturation = "saturation"
	ColorMode  = "color"
	Luminosity = "luminosity"
)

// Blend holds the currently active blend mode.
type Blend struct {
	Mode  string
	Modes []string
}

type Color struct {
	R, G, B float64
}

// NewBlend initializes a new Blend.
func NewBlend() *Blend {
	return &Blend{
		Modes: []string{
			Normal,
			Darken,
			Lighten,
			Multiply,
			Screen,
			Overlay,
			SoftLight,
			HardLight,
			ColorDodge,
			ColorBurn,
			Difference,
			Exclusion,
			Hue,
			Saturation,
			ColorMode,
			Luminosity,
		},
	}
}

// Set activate one of the supported blend mode.
func (bl *Blend) Set(blendType string) {
	if math.Contains(bl.Modes, blendType) {
		bl.Mode = blendType
	}
}

// Get returns the currently active blend mode.
func (bl *Blend) Get() string {
	if len(bl.Mode) > 0 {
		return bl.Mode
	}
	return ""
}

func (bl *Blend) Lum(rgb Color) float64 {
	return 0.3*rgb.R + 0.59*rgb.G + 0.11*rgb.B
}

func (bl *Blend) SetLum(rgb Color, l float64) Color {
	delta := l - bl.Lum(rgb)
	return bl.Clip(Color{
		rgb.R + delta,
		rgb.G + delta,
		rgb.B + delta,
	})
}

func (bl *Blend) Clip(rgb Color) Color {
	r, g, b := rgb.R, rgb.G, rgb.B

	l := bl.Lum(rgb)
	min := math.Min(r, g, b)
	max := math.Max(r, g, b)

	if min < 0 {
		r = l + (((r - l) * l) / (l - min))
		g = l + (((g - l) * l) / (l - min))
		b = l + (((b - l) * l) / (l - min))
	}
	if max > 1 {
		r = l + (((r - l) * (1 - l)) / (max - l))
		g = l + (((g - l) * (1 - l)) / (max - l))
		b = l + (((b - l) * (1 - l)) / (max - l))
	}

	return Color{R: r, G: g, B: b}
}

func (bl *Blend) Sat(rgb Color) float64 {
	return math.Max(rgb.R, rgb.G, rgb.B) - math.Min(rgb.R, rgb.G, rgb.B)
}

type channel struct {
	key string
	val float64
}

func (bl *Blend) SetSat(rgb Color, s float64) Color {
	color := map[string]float64{
		"R": rgb.R,
		"G": rgb.G,
		"B": rgb.B,
	}
	channels := make([]channel, 0, 3)
	for k, v := range color {
		channels = append(channels, channel{k, v})
	}
	sort.Slice(channels, func(i, j int) bool { return channels[i].val < channels[j].val })
	minChan, midChan, maxChan := channels[0].key, channels[1].key, channels[2].key

	if color[maxChan] > color[minChan] {
		color[midChan] = (((color[midChan] - color[minChan]) * s) / (color[maxChan] - color[minChan]))
		color[maxChan] = s
	} else {
		color[midChan], color[maxChan] = 0, 0
	}
	color[minChan] = 0

	return Color{
		R: color["R"],
		G: color["G"],
		B: color["B"],
	}
}
