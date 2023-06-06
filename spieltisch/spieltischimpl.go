package spieltisch

import (
	"app/score"
	"app/spielbrett"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type data struct {
	spielbrett spielbrett.Spielbrett
	scoreSpiel score.Score
	//	aktuellerSpieler score.Score
}

func New() *data {
	var tisch *data
	tisch = new(data)

	spielstand := score.New()

	tisch.scoreSpiel = spielstand

	tisch.spielbrett = spielbrett.New(spielstand)

	return tisch
}

func (t *data) GetFyneTisch() fyne.CanvasObject {
	// Die Anzeige mit den Spielständen wird in top gespeichert:
	top := t.anzeigetafel()
	// Der Spieltisch wird im Format "Border" angezeigt:
	return container.NewBorder(top, nil, nil, nil, t.spielbrett.GetBrett())
}

func (t *data) anzeigetafel() fyne.CanvasObject {
	// In scoreX wird der Inhalt der Anzeige gespeichert:
	score1 := widget.NewLabel("Spieler 1")
	score2 := widget.NewLabel("Spieler 2")
	score3 := widget.NewLabelWithData(t.scoreSpiel.Spieler1())
	score4 := widget.NewLabelWithData(t.scoreSpiel.Spieler2())
	//	score5 := widget.NewLabelWithData(t.)

	// Angabe, wie die Anzeige der Inhalte erfolgen soll:
	return container.New(layout.NewGridLayout(2), score1, score2, score3, score4 /*, score5*/)
}
