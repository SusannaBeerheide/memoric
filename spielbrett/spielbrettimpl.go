package spielbrett

import (
	"app/karten"
	"fmt"

	"fyne.io/fyne/v2"
)

type spielbrett struct {
	fynekarten []fyne.CanvasObject
	karten     []*karten.Karte
}

func New() *spielbrett {
	var sb *spielbrett
	sb = new(spielbrett)
	sb.fynekarten = make([]fyne.CanvasObject, 12)
	sb.karten = make([]*karten.Karte, 12)

	for i := 0; i < 12; i++ {
		fmt.Println("", i)
		index := i
		k := karten.NewKarte("Hello World!",
			func() {
				sb.KarteAusgewaehlt(index)
			},
		)
		sb.karten[i] = k
		sb.fynekarten[i] = fyne.CanvasObject(k)
	}

	return sb
}

func (sb *spielbrett) GetKarten() []fyne.CanvasObject {

	// karten := make([]karten.Karte, len(sb.karten))
	// for i, v := range karten {
	// 	karten[i] = fyne.CanvasObject(v)
	// }
	return sb.fynekarten
}

func (sb *spielbrett) KarteAusgewaehlt(kartennr int) {
	fmt.Println("Karte wurde ausgewählt")
	fmt.Println("Kartennummer ist: ", kartennr)
	ausgewaehlteKarte := sb.karten[kartennr]
	//fmt.Println("Karten ist offen: ", ausgewaehlteKarte.IstOffen())
	if ausgewaehlteKarte.Offen {
		ausgewaehlteKarte.Schliessen()
	} else {
		ausgewaehlteKarte.Oeffnen()
		//fmt.Println("Karten ist offen: ", ausgewaehlteKarte.IstOffen())
	}

}
