package karten

// Interface siehe Fyne-Dokumentation: developer.fyne.io/extend/custom-widget

// Quelle für die nachfolgende Implementierung der Fyne-Funktionalitäten:
// https://github.com/stuartdd2/developer.fyne.io/blob/master/extend/custom-widget.md

import (
	"fmt"
	"image/color"

	"app/musikAbspieler"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// A text widget with theamed background and foreground
type data struct {
	widget.BaseWidget        // Inherit from BaseWidget
	text              string // The text to display in the widget
	offen             bool   // true = offen; false = verdeckt
	onTapped          func()
	weg               bool                          // Karte ist verschwunden, d.h. sie wird nicht mehr angezeigt.
	musikAbspieler    musikAbspieler.MusikAbspieler //Importierte Klasse zum Abspielen der Musik
}

func NewKarte(text string, brettFunc func()) *data { // Create a Widget and Extend (initialiase) the BaseWidget
	w := &data{ // Create this widget with an initial text value
		text:           text,      // Text, z.B. auf der Rückseite der Spielkarte "memoric"
		onTapped:       brettFunc, // Aufruf der Funktion brettFunc, wenn die Karte angeklickt wird.
		musikAbspieler: musikAbspieler.New(text),
	}
	w.ExtendBaseWidget(w) // Initialiase the BaseWidget
	return w
}

func (w *data) Tapped(_ *fyne.PointEvent) { // Methode, die ausgeführt wird, wenn die Karte geklickt wird.
	w.onTapped()
}

func (w *data) IstOffen() bool { // Methode, die zurückgibt, ob die Karte geöffnet oder geschlossen ist.
	return w.offen
}

func (w *data) Oeffnen() { // Methode zum Öffnen der Karte.
	if w.weg || w.offen {
		return
	}
	w.offen = true // Karte wird geöffnet
	fmt.Println("Karten wird geoeffnet")
	w.musikAbspieler.Spielen() //Musik wird abgespielt
	fmt.Println("Musik wurde abegespielt")
	w.Refresh() // Karte wird "refreshed"
}

func (w *data) Schliessen() { // Methode zum Schliessen der Karte.
	w.offen = false
	w.musikAbspieler.Stoppen()
	w.Refresh()
}

func (w *data) MusikStoppen() {
	w.musikAbspieler.Stoppen()
}
func (w *data) Verschwinden() {
	w.weg = true
	w.offen = false
	w.Refresh()
}

func (w *data) Inhalt() string {
	return w.text
}

//Hier folgen nun Fyne spezifische Funktionalitäten (Quelle s.o. Zeile 6)

// Create the renderer. This is called by the fyne application
func (w *data) CreateRenderer() fyne.WidgetRenderer {
	// Pass this widget to the renderer so it can access the text field
	return newKarteRenderer(w)
}

// Widget Renderer code starts here
type karteRenderer struct {
	widget     *data             // Reference to the widget holding the current state
	background *canvas.Rectangle // A background rectangle
	text       *canvas.Text      // The text
}

// Create the renderer with a reference to the widget
// Note: The background and foreground colours are set from the current theme.
//
// Do not size or move canvas objects here.

// Die Karte, die beim ersten Anzeigen angezeigt wird, also die rote Rückseite der Spielkarte.
func newKarteRenderer(Karte *data) *karteRenderer {
	return &karteRenderer{
		widget:     Karte,
		background: canvas.NewRectangle(color.RGBA{250, 128, 114, 255}),
		text:       canvas.NewText("memoric", theme.ForegroundColor()),
	}
}

// The Refresh() method is called if the state of the widget changes or the
// theme is changed
//
// Note: The background and foreground colours are set from the current theme
func (r *karteRenderer) Refresh() { // Funktion, welche die Funktionalitäten der Karte aktualisieren.
	//	fmt.Println("Refresh ist aufgerufen.")
	r.text.Color = theme.ForegroundColor() // Reicht die Textfarbe in die Fyne-Logik rüber.
	if r.widget.weg {
		r.background.FillColor = color.Transparent
		r.text.Text = ""
	} else if r.widget.offen { // Wenn die Karte offen ist:
		r.background.FillColor = color.RGBA{135, 206, 235, 255} // Farbe setzen
		r.text.Text = ""                                        // Inhalt der Karte
	} else {
		r.background.FillColor = color.RGBA{250, 128, 114, 255} // Farbe der geschlossenen Karte
		r.text.Text = "memoric"                                 // Reicht den Text in die Fyne-Logik rüber.
	}
	r.background.Refresh() // Redraw the background first
	r.text.Refresh()       // Redraw the text on top

}

// Given the size required by the fyne application move and re-size the
// canvas objects.
func (r *karteRenderer) Layout(s fyne.Size) {
	// Measure the size of the text so we can calculate the center offset.
	ts := fyne.MeasureText(r.text.Text, r.text.TextSize, r.text.TextStyle)
	// Center the text
	r.text.Move(fyne.Position{X: (s.Width - ts.Width) / 2, Y: (s.Height - ts.Height) / 2})
	// Make sure the background fills the widget
	r.background.Resize(s)
}

// Create a minimum size for the widget.
// The smallest size is the size of the text with a border defined by the theme padding
func (r *karteRenderer) MinSize() fyne.Size {
	// Measure the size of the text so we can calculate a border size.
	ts := fyne.MeasureText(r.text.Text, r.text.TextSize, r.text.TextStyle)
	// Use the theme padding to set a border size
	return fyne.NewSize(ts.Width+theme.Padding()*4, ts.Height+theme.Padding()*4)
}

// Return a list of each canvas object.
func (r *karteRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.background, r.text}
}

// Cleanup if resources have been allocated
func (r *karteRenderer) Destroy() {}
