package imglib

import (
	"image"
	"image/color"
)

/*
Все что светлее граничного цвета заменяем на replaceColor.
*/
func ReplaceGrayscaleLighter(im *image.Gray, borderColor color.Gray, replaceColor color.Gray) {
	for x := 0; x < im.Bounds().Dx(); x++ {
		for y := 0; y < im.Bounds().Dy(); y++ {
			if im.GrayAt(x, y).Y > borderColor.Y {
				im.SetGray(x, y, replaceColor)
			}
		}
	}
}

/*
Все что темнее граничного цвета заменяем на replaceColor.
*/
func ReplaceGrayscaleDarker(im *image.Gray, borderColor color.Gray, replaceColor color.Gray) {
	for x := 0; x < im.Bounds().Dx(); x++ {
		for y := 0; y < im.Bounds().Dy(); y++ {
			if im.GrayAt(x, y).Y < borderColor.Y {
				im.SetGray(x, y, replaceColor)
			}
		}
	}
}
