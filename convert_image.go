package imglib

import (
	"image"
	"image/color"
	"math"
)

func ToGrayscale(im image.Image) *image.Gray {
	gray := image.NewGray(im.Bounds())

	for y := 0; y < im.Bounds().Dy(); y++ {
		for x := 0; x < im.Bounds().Dx(); x++ {
			gray.Set(x, y, ColorToGray(im.At(x, y)))
		}
	}

	return gray
}

func ToRGBA(im *image.Gray) image.Image {
	imRGBA := image.NewRGBA(im.Bounds())

	for y := 0; y < im.Bounds().Dy(); y++ {
		for x := 0; x < im.Bounds().Dx(); x++ {
			gray := im.GrayAt(x, y).Y
			imRGBA.Set(x, y, color.RGBA{R: gray, G: gray, B: gray, A: 255})
		}
	}

	return imRGBA
}

func ReduceGrayscale(im *image.Gray, shades float64) {
	d := 255 / (shades - 1)
	for x := 0; x < im.Bounds().Max.X; x++ {
		for y := 0; y < im.Bounds().Max.Y; y++ {
			c := float64(im.GrayAt(x, y).Y)
			c = math.Round(c/d) * d
			im.Set(x, y, color.Gray{Y: uint8(c)})
		}
	}
}
