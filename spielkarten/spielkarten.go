package spielkarten

import "fyne.io/fyne/v2"

type Karte interface {
	IstGeoeffnet() bool
	GetQuadrat() fyne.CanvasObject
}
