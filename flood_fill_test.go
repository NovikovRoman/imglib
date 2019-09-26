package imglib

import (
	"github.com/disintegration/imaging"
	"github.com/stretchr/testify/require"
	"image"
	"image/color"
	"path/filepath"
	"testing"
)

func TestFill(t *testing.T) {
	im, err := imaging.Open(filepath.Join(testdata, "fill.png"))
	require.Nil(t, err)

	p := image.Point{
		X: 32,
		Y: 32,
	}
	FloodFill(im, p, color.RGBA{R: 255, G: 0, B: 0, A: 128}, color.RGBA{R: 0, G: 0, B: 0, A: 255})

	err = imaging.Save(im, filepath.Join(testdata, "fill_result.png"))
	require.Nil(t, err)
}

func BenchmarkFill(b *testing.B) {
	p := image.Point{
		X: 53,
		Y: 50,
	}
	for i := 0; i < b.N; i++ {
		im, _ := imaging.Open(filepath.Join(testdata, "fill.png"))
		FloodFill(im, p, color.RGBA{R: 255, G: 0, B: 0, A: 128}, color.RGBA{R: 0, G: 0, B: 0, A: 255})
	}
}
