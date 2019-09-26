package imglib

import (
	"github.com/disintegration/imaging"
	"github.com/stretchr/testify/require"
	"image/color"
	"path/filepath"
	"testing"
)

const testdata = "testdata"

func TestTrimGrayscale(t *testing.T) {
	im, err := imaging.Open(filepath.Join(testdata, "trim.png"))
	require.Nil(t, err)

	imGray := TrimGrayscale(ToGrayscale(im), color.Gray{Y: 255})

	require.Equal(t, imGray.Bounds().Max.X, 137)
	require.Equal(t, imGray.Bounds().Max.Y, 43)
}

func TestTrimGrayscaleLeft(t *testing.T) {
	im, err := imaging.Open(filepath.Join(testdata, "trim.png"))
	require.Nil(t, err)

	imGray := TrimGrayscaleLeft(ToGrayscale(im), color.Gray{Y: 255})

	require.Equal(t, imGray.Bounds().Max.X, 180)
	require.Equal(t, imGray.Bounds().Max.Y, 60)
}

func TestTrimGrayscaleRight(t *testing.T) {
	im, err := imaging.Open(filepath.Join(testdata, "trim.png"))
	require.Nil(t, err)

	imGray := TrimGrayscaleRight(ToGrayscale(im), color.Gray{Y: 255})

	require.Equal(t, imGray.Bounds().Max.X, 157)
	require.Equal(t, imGray.Bounds().Max.Y, 60)
}

func TestTrimGrayscaleTop(t *testing.T) {
	im, err := imaging.Open(filepath.Join(testdata, "trim.png"))
	require.Nil(t, err)

	imGray := TrimGrayscaleTop(ToGrayscale(im), color.Gray{Y: 255})

	require.Equal(t, imGray.Bounds().Max.X, 200)
	require.Equal(t, imGray.Bounds().Max.Y, 45)
}

func TestTrimGrayscaleBottom(t *testing.T) {
	im, err := imaging.Open(filepath.Join(testdata, "trim.png"))
	require.Nil(t, err)

	imGray := TrimGrayscaleBottom(ToGrayscale(im), color.Gray{Y: 255})

	require.Equal(t, imGray.Bounds().Max.X, 200)
	require.Equal(t, imGray.Bounds().Max.Y, 58)
}
