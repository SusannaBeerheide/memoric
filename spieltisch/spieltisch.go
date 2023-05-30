package spieltisch

import "fyne.io/fyne/v2"

type Spieltisch interface {
	GetFyneTisch() fyne.CanvasObject
}
