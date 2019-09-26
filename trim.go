package imglib

import (
	"image"
	"image/color"
	"image/draw"
)

func TrimGrayscale(im *image.Gray, backColor color.Gray) *image.Gray {
	top := topSpaceOffset(im, backColor)
	if top < 0 {
		return nil
	}

	left := leftSpaceOffset(im, backColor)
	if left < 0 {
		return nil
	}

	right := rightSpaceOffset(im, backColor)
	bottom := bottomSpaceOffset(im, backColor)

	if left == 0 && right == im.Bounds().Max.X-1 && top == 0 && bottom == im.Bounds().Max.Y-1 {
		return im
	}

	r := image.Rect(0, 0, right-left+1, bottom-top+1)
	img := image.NewGray(r)
	draw.Draw(img, r, im, image.Point{X: left, Y: top}, draw.Src)
	return img
}

func TrimGrayscaleLeft(im *image.Gray, backColor color.Gray) *image.Gray {
	left := leftSpaceOffset(im, backColor)
	if left < 0 {
		return nil
	}

	if left == 0 {
		return im
	}

	r := image.Rect(0, 0, im.Bounds().Max.X-left, im.Bounds().Max.Y)
	img := image.NewGray(r)
	draw.Draw(img, r, im, image.Point{X: left, Y: 0}, draw.Src)
	return img
}

func TrimGrayscaleRight(im *image.Gray, backColor color.Gray) *image.Gray {
	right := rightSpaceOffset(im, backColor)
	if right < 0 {
		return nil
	}

	if right == im.Bounds().Max.X-1 {
		return im
	}

	r := image.Rect(0, 0, right+1, im.Bounds().Max.Y)
	img := image.NewGray(r)
	draw.Draw(img, r, im, image.ZP, draw.Src)
	return img
}

func TrimGrayscaleTop(im *image.Gray, backColor color.Gray) *image.Gray {
	top := topSpaceOffset(im, backColor)
	if top < 0 {
		return nil
	}

	if top == 0 {
		return im
	}

	r := image.Rect(0, 0, im.Bounds().Max.X, im.Bounds().Max.Y-top)
	img := image.NewGray(r)
	draw.Draw(img, r, im, image.Point{X: 0, Y: top}, draw.Src)
	return img
}

func TrimGrayscaleBottom(im *image.Gray, backColor color.Gray) *image.Gray {
	bottom := bottomSpaceOffset(im, backColor)
	if bottom < 0 {
		return nil
	}

	if bottom == im.Bounds().Max.Y-1 {
		return im
	}

	r := image.Rect(0, 0, im.Bounds().Max.X, bottom+1)
	img := image.NewGray(r)
	draw.Draw(img, r, im, image.ZP, draw.Src)
	return img
}

func topSpaceOffset(im *image.Gray, backColor color.Gray) int {
	for y := 0; y < im.Bounds().Max.Y; y++ {
		for x := 0; x < im.Bounds().Max.X; x++ {
			if im.GrayAt(x, y) != backColor {
				return y
			}
		}
	}

	return -1
}

func bottomSpaceOffset(im *image.Gray, backColor color.Gray) int {
	for y := im.Bounds().Max.Y - 1; y >= 0; y-- {
		for x := 0; x < im.Bounds().Max.X; x++ {
			if im.GrayAt(x, y) != backColor {
				return y
			}
		}
	}

	return -1
}

func leftSpaceOffset(im *image.Gray, backColor color.Gray) int {
	for x := 0; x < im.Bounds().Max.X; x++ {
		for y := 0; y < im.Bounds().Max.Y; y++ {
			if im.GrayAt(x, y) != backColor {
				return x
			}
		}
	}

	return -1
}

func rightSpaceOffset(im *image.Gray, backColor color.Gray) int {
	for x := im.Bounds().Max.X - 1; x >= 0; x-- {
		for y := 0; y < im.Bounds().Max.Y; y++ {
			if im.GrayAt(x, y) != backColor {
				return x
			}
		}
	}

	return -1
}
