package factorySimpleBox

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.interfaces/iStage"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/simple"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/dimensions"
)

// en: This function create a path of a box with rounded corners into the platform using density
//
// pt_br: Esta função cria o caminho de um quadrado com cantos arredondados na plataforma usando densidade
func NewBoxWithRoundedCorners(id string, stage iStage.IStage, platform, scratchPad iotmaker_platform_IDraw.IDraw, x, y, width, height, border int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *simple.BoxWithRoundedCorners {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.SetInt(x)
	x = densityCalc.Int()

	densityCalc.SetInt(y)
	y = densityCalc.Int()

	densityCalc.SetInt(width)
	width = densityCalc.Int()

	densityCalc.SetInt(height)
	height = densityCalc.Int()

	densityCalc.SetInt(border)
	border = densityCalc.Int()

	rect := &simple.BoxWithRoundedCorners{
		Sprite: basic.Sprite{
			Id:         id,
			Stage:      stage,
			Platform:   platform,
			ScratchPad: scratchPad,
			Dimensions: dimensions.Dimensions{
				X:      float64(x),
				Y:      float64(y),
				Width:  float64(width),
				Height: float64(height),
				Border: float64(border),
			},
			OutBoxDimensions: dimensions.Dimensions{
				X:      float64(x),
				Y:      float64(y),
				Width:  float64(width),
				Height: float64(height),
			},
		},
	}
	rect.Crete()

	return rect
}