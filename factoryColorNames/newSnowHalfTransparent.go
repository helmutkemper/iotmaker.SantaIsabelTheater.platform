package factoryColorNames

import "image/color"

func NewSnowHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xff, G: 0xfa, B: 0xfa, A: 0x80} // rgb(255, 250, 250)
}
