package imglib

import (
	"image/color"
	"math"
)

// цвет в оттенок серого
func ColorToGray(c color.Color) *color.Gray {
	r, g, b, _ := c.RGBA()
	r = r >> 8
	g = g >> 8
	b = b >> 8
	return &color.Gray{Y: uint8(math.Round(0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)))}
	// варианты:
	// return 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
	// return 0.212671*float64(r) + 0.71516*float64(g) + 0.072169*float64(b)
}
