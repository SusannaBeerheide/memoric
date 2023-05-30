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
	top := t.anzeigetafel()
	return container.NewBorder(top, nil, nil, nil, t.spielbrett.GetBrett())

}

func (t *data) anzeigetafel() fyne.CanvasObject {
	score1 := widget.NewLabelWithData(t.scoreSpiel.Spieler1())
	score2 := widget.NewLabelWithData(t.scoreSpiel.Spieler2())
	score3 := widget.NewLabel("Testing")
	return container.New(layout.NewGridLayout(2), score1, score2, score3)
}
