package spieltisch

import (
	"app/score"
	"app/spielbrett"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type data struct { // Der Spieltisch besteht aus
	spielbrett spielbrett.Spielbrett // dem Spielbrett
	scoreSpiel score.Score           // den Spielstand
}

func New() *data {
	var tisch *data
	tisch = new(data)

	spielstand := score.New() // Der Spielstand wird

	tisch.scoreSpiel = spielstand

	tisch.spielbrett = spielbrett.New(spielstand)

	return tisch
}

// Die Methode zeigt den Spieltisch an:
func (t *data) GetFyneTisch() fyne.CanvasObject {
	// Die Anzeige mit den Spielständen wird in top gespeichert:
	top := t.anzeigetafel()
	// Der Spieltisch wird im Format "Border" angezeigt:
	return container.NewBorder(top, nil, nil, nil, t.spielbrett.GetBrett())
}

// Die Funktion zeigt den aktuellen Spielstand auf dem Spieltisch an:
func (t *data) anzeigetafel() fyne.CanvasObject {
	// In scoreX wird der Inhalt der Anzeige gespeichert:
	score1 := widget.NewLabel("Spieler / Spielerin 1")
	//score1.TextSize = fyne.TextSize
	score2 := widget.NewLabel("Spieler / Spielerin 2")
	score2.Alignment = fyne.TextAlignTrailing
	score3 := widget.NewLabelWithData(t.scoreSpiel.Spieler1())
	score4 := widget.NewLabelWithData(t.scoreSpiel.Spieler2())
	score4.Alignment = fyne.TextAlignTrailing
	score5 := widget.NewLabel("Spieler / Spielerin X ist dran!")
	score5.Alignment = fyne.TextAlignCenter
	score5.TextStyle.Bold = true
	score6 := widget.NewLabel("")
	// Angabe, wie die Anzeige der Inhalte erfolgen soll:
	return container.New(layout.NewGridLayout(3), score1, score5, score2, score3, score6, score4)
}
