package spielkarten

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type karte struct {
	geoeffnet bool
	quadrat   fyne.CanvasObject
}

// https://github.com/stuartdd2/developer.fyne.io/blob/master/extend/custom-widget.md

type unserWidget struct {
	widget.BaseWidget
	text string
}

type myWidgetRenderer struct {
	widget     *unserWidget
	background *canvas.Rectangle
	text       *canvas.Text
}

func newWidget() *unserWidget {
	var w *unserWidget

	w.text = "TestTest"
	w.ExtendBaseWidget(w)
	return w
}

func (w *unserWidget) CreateRenderer() fyne.WidgetRenderer {
	return BaseRen	return widget.NewSimpleRenderer(canvas.NewText(w.text, color.Black))
}

func New() *karte {
	var k *karte
	k = new(karte)
	k.quadrat = newWidget()
	return k
}

func (k *karte) IstGeoeffnet() bool {
	return k.geoeffnet
}

func (k *karte) GetQuadrat() fyne.CanvasObject {
	return k.quadrat
}
