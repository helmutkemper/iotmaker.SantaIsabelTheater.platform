package selectBox

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	mouseWebBrowser "github.com/helmutkemper/iotmaker.platform.webbrowser/mouse"
	"github.com/helmutkemper/iotmaker.platform/abstractType/basicBox"
	"github.com/helmutkemper/iotmaker.platform/abstractType/genericTypes"
)

func NewResizeBoxFromBasicBox(basicBox *basicBox.BasicBox, offsetX, offsetY, width, height, border int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *ResizeBoxes {

	dm := genericTypes.Dimensions{}
	dm = genericTypes.NewDimensions(dm, offsetX, offsetY, width, height, border, density, iDensity)

	rb := &ResizeBoxes{
		MouseCornerA:           mouseWebBrowser.KCursorNWSeResize,
		MouseCornerB:           mouseWebBrowser.KCursorNSResize,
		MouseCornerC:           mouseWebBrowser.KCursorNESwResize,
		MouseCornerD:           mouseWebBrowser.KCursorEWResize,
		MouseCornerE:           mouseWebBrowser.KCursorNWSeResize,
		MouseCornerF:           mouseWebBrowser.KCursorNSResize,
		MouseCornerG:           mouseWebBrowser.KCursorNESwResize,
		MouseCornerH:           mouseWebBrowser.KCursorEWResize,
		MouseGeneric:           mouseWebBrowser.KCursorAuto,
		Platform:               basicBox.Platform,
		ScratchPad:             basicBox.ScratchPad,
		Dimensions:             dm,
		FatherOutBoxDimensions: basicBox.OutBoxDimensions,
	}
	rb.Create()

	return rb
}
