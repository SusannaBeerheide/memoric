package spielbrett

import (
	"app/karten"
	"fmt"

	"fyne.io/fyne/v2"
)

type spielbrett struct {
	karten []karten.Karte
}

func New() *spielbrett {
	var sb *spielbrett
	sb = new(spielbrett)
	sb.karten = make([]karten.Karte, 12)

	for i := 0; i < 12; i++ {
		fmt.Println("", i)
		index := i
		k := karten.NewKarte("Hello World!",
			func() {
				sb.KarteAusgewaehlt(index)
			},
		)
		sb.karten[i] = k
	}

	return sb
}

func (sb *spielbrett) GetKartenFuerFyne() []fyne.CanvasObject {

	fyneKarten := make([]fyne.CanvasObject, len(sb.karten))
	for i, v := range sb.karten {
		fyneKarten[i] = fyne.CanvasObject(v)
	}
	return fyneKarten
}

func (sb *spielbrett) KarteAusgewaehlt(kartennr int) {
	fmt.Println("Karte wurde ausgewählt")
	fmt.Println("Kartennummer ist: ", kartennr)
	ausgewaehlteKarte := sb.karten[kartennr]
	//fmt.Println("Karten ist offen: ", ausgewaehlteKarte.IstOffen())
	if ausgewaehlteKarte.IstOffen() {
		ausgewaehlteKarte.Schliessen()
	} else {
		ausgewaehlteKarte.Oeffnen()
		//fmt.Println("Karten ist offen: ", ausgewaehlteKarte.IstOffen())
	}

}
