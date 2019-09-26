package imglib

import (
	"image"
	"image/color"
)

type ImageSet interface {
	image.Image
	Set(x int, y int, c color.Color)
}

type fill struct {
	image       ImageSet
	colorFill   color.Color
	colorBorder color.Color
	colorBack   color.Color
}

func (f *fill) canNotPainted(x int, y int) bool {
	cPoint := f.image.At(x, y)

	if cPoint == f.colorFill {
		return true
	}

	if f.colorBorder == nil {
		return cPoint != f.colorBack
	}

	return cPoint == f.colorBorder
}

func (f *fill) canPainted(x int, y int) bool {
	return !f.canNotPainted(x, y)
}

func (f *fill) width() int {
	return f.image.Bounds().Dx()
}

func (f *fill) height() int {
	return f.image.Bounds().Dy()
}

func (f *fill) paint(x int, y int) {
	f.image.Set(x, y, f.colorFill)
}

func (f *fill) fillCalc(x int, y int) {
	for {
		ox := x
		oy := y

		for y != 0 && f.canPainted(x, y-1) {
			y--
		}

		for x != 0 && f.canPainted(x-1, y) {
			x--
		}

		if x == ox && y == oy {
			break
		}
	}

	f.fillCore(x, y)
}

func (f *fill) fillCore(x int, y int) {
	lastRowLength := 0

	for {
		rowLength := 0
		sx := x

		if lastRowLength != 0 && f.canNotPainted(x, y) {
			for {
				lastRowLength--
				if lastRowLength == 0 {
					return
				}

				x++
				if f.canPainted(x, y) {
					break
				}
			}

			sx = x

		} else {
			for x != 0 && f.canPainted(x-1, y) {
				x--
				f.paint(x, y)

				if y != 0 && f.canPainted(x, y-1) {
					f.fillCalc(x, y-1)
				}

				rowLength++
				lastRowLength++
			}

		}

		for sx < f.width() && f.canPainted(sx, y) {
			f.paint(sx, y)
			rowLength++
			sx++
		}

		if rowLength < lastRowLength {
			for end := x + lastRowLength; sx < end; sx++ {
				if f.canNotPainted(sx, y) {
					continue
				}
				f.fillCore(sx, y)
			}

		} else if rowLength > lastRowLength && y != 0 {
			for ux := x + lastRowLength; ux < sx; ux++ {
				if f.canNotPainted(ux, y-1) {
					continue
				}
				f.fillCalc(ux, y-1)
			}

		}

		lastRowLength = rowLength
		y++
		if lastRowLength == 0 || y >= f.height() {
			break
		}
	}
}

// Если не задан цвет границы, то берем в качестве цвета фона цвет под начальной точкой
//
// many thanks to Adam Milazzo http://www.adammil.net/blog/v126_A_More_Efficient_Flood_Fill.html
func FloodFill(im image.Image, point image.Point, colorFill color.Color, colorBorder color.Color) {
	ims := im.(ImageSet)

	f := &fill{
		image:       ims,
		colorFill:   colorFill,
		colorBorder: colorBorder,
	}

	if colorBorder == nil {
		f.colorBack = ims.At(point.X, point.Y)
	}

	if f.canNotPainted(point.X, point.Y) {
		return
	}

	f.fillCalc(point.X, point.Y)
}
