package spielbrett

import (
	"app/spielkarten"

	"fyne.io/fyne/v2"
)

type spielbrett struct {
	karten []spielkarten.Karte
}

func New() *spielbrett {
	var sb *spielbrett
	sb = new(spielbrett)
	sb.karten = make([]spielkarten.Karte, 12)

	for i := 0; i < 12; i++ {
		k := spielkarten.New()
		sb.karten[i] = k
	}

	return sb
}

func (sb *spielbrett) GetKarten() []spielkarten.Karte {
	return sb.karten
}

func (sb *spielbrett) GetFyneKarten() []fyne.CanvasObject {
	var grafikKarten []fyne.CanvasObject
	grafikKarten = make([]fyne.CanvasObject, 12)
	for i := 0; i < len(sb.karten); i++ {
		grafikKarten[i] = sb.karten[i].GetQuadrat()
	}
	return grafikKarten
}
