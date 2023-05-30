package score

import (
	"strconv"

	"fyne.io/fyne/v2/data/binding"
)

type data struct {
	aktuellerSpieler bool // false = Spieler 1 ist dran, true = Spieler 2 ist dran
	scoreSpieler1    int  // mit 0 initialisiert
	scoreSpieler2    int  // mit 0 initialisiert
	spieler1String   binding.String
	spieler2String   binding.String
}

func New() *data {
	var sc *data
	sc = new(data)
	sc.spieler1String = binding.NewString()
	sc.spieler2String = binding.NewString()
	sc.setScoreStrings()
	return sc
}

func (sc *data) GetAktuellerSpieler() bool {
	return sc.aktuellerSpieler
}

func (sc *data) setScoreStrings() {
	sc.spieler1String.Set("Spieler 1 Score: " + strconv.Itoa(sc.scoreSpieler1))
	sc.spieler2String.Set("Spieler 2 Score: " + strconv.Itoa(sc.scoreSpieler2))

}

func (sc *data) Spieler1() binding.String {
	return sc.spieler1String
}

func (sc *data) Spieler2() binding.String {
	return sc.spieler2String
}

func (sc *data) PaarGefunden() {
	if sc.aktuellerSpieler {
		sc.scoreSpieler2++
	} else {
		sc.scoreSpieler1++
	}
	sc.setScoreStrings()
}

func (sc *data) SpielerWechseln() {
	sc.aktuellerSpieler = !sc.aktuellerSpieler // Der jeweils andere Spieler ist mit Spielen dran.
}
