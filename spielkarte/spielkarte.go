package spielkarte

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Karte interface {
	// istGeoeffnet() bool
	// getQuadrat() fyne.CanvasObject
}

type KarteImpl struct {
	geoeffnet bool
	quadrat   fyne.CanvasObject
}

func New() *KarteImpl {
	var karte *KarteImpl = new(KarteImpl)
	karte.quadrat = canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	return karte
}

// func (karte *KarteImpl) istGeoeffnet() bool {
// 	return karte.geoeffnet
// }

// func (karte *KarteImpl) getQuadrat() fyne.CanvasObject {
// 	return karte.quadrat
// }
