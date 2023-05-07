package spielbrett

import (
	"app/karten"

	"fyne.io/fyne/v2"
)

type spielbrett struct {
	karten []fyne.CanvasObject
}

func New() *spielbrett {
	var sb *spielbrett
	sb = new(spielbrett)
	sb.karten = make([]fyne.CanvasObject, 12)

	for i := 0; i < 12; i++ {
		k := karten.NewKarte("Hello World!")
		sb.karten[i] = k
	}

	return sb
}

func (sb *spielbrett) GetKarten() []fyne.CanvasObject {
	return sb.karten
}
