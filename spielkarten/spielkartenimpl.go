package spielkarten

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type karte struct {
	geoeffnet bool
	quadrat   fyne.CanvasObject
}

func New() *karte {
	var k *karte
	k = new(karte)
	k.quadrat = canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	return k
}

func (k *karte) IstGeoeffnet() bool {
	return k.geoeffnet
}

func (k *karte) GetQuadrat() fyne.CanvasObject {
	return k.quadrat
}
