package imglib

import (
	"github.com/disintegration/imaging"
	"github.com/stretchr/testify/require"
	"image/color"
	"path/filepath"
	"testing"
)

func TestReplaceGrayscaleLighter(t *testing.T) {
	im, err := imaging.Open(filepath.Join(testdata, "replace.jpg"))
	require.Nil(t, err)

	imGray := ToGrayscale(im)
	ReplaceGrayscaleLighter(imGray, color.Gray{Y: 120}, color.Gray{Y: 255})

	err = imaging.Save(imGray, filepath.Join(testdata, "replace_result.png"))
	require.Nil(t, err)
}

func TestReplaceGrayscaleDarker(t *testing.T) {
	im, err := imaging.Open(filepath.Join(testdata, "replace.jpg"))
	require.Nil(t, err)

	imGray := ToGrayscale(im)
	ReplaceGrayscaleDarker(imGray, color.Gray{Y: 80}, color.Gray{Y: 0})

	err = imaging.Save(imGray, filepath.Join(testdata, "replace_result.png"))
	require.Nil(t, err)
}
