package spielkarte

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Karte interface {
	istGeoeffnet() bool
	getQuadrat() fyne.CanvasObject
}

type KarteImpl struct {
	geoeffnet bool
	quadrat   fyne.CanvasObject
}

func NewKarte() KarteImpl {

	return KarteImpl{
		false,
		canvas.NewRectangle(color.RGBA{255, 0, 0, 255}),
	}
}

// func (karte KarteImpl) istGeoeffnet() bool {
// 	return karte.geoeffnet
// }

func (karte KarteImpl) GetQuadrat() fyne.CanvasObject {
	return karte.quadrat
}
