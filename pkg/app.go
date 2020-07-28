package pkg

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg/widgets"
	"image/color"
	"log"
)

// Window used to define the UI window for the application.
// Currently will just cater for single window per app. This will be
// reviewed in the future.
type Window struct {
	width  int
	height int
	title  string

	// slice of panels. Should probably do as a map???
	// Then again, slice can be used for render order?
	panels []widgets.Panel
}

func NewWindow(width int, height int, title string) Window {
	w := Window{}
	w.height = height
	w.width = width
	w.title = title
	w.panels = []widgets.Panel{}
	return w
}

func (w *Window) AddPanel(panel widgets.Panel) error {
	w.panels = append(w.panels, panel)
	return nil
}

/////////////////////// EBiten specifics below... /////////////////////////////////////////////
func (w *Window) Update(screen *ebiten.Image) error {
	return nil
}

func (w *Window) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x11, 0x11, 0x11, 0xff})

	for _, panel := range w.panels {
		panel.Draw(screen)
	}
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return w.width, w.height
}

func (w *Window) MainLoop() error {

	ebiten.SetWindowSize(w.width, w.height)
	ebiten.SetWindowTitle(w.title)
	if err := ebiten.RunGame(w); err != nil {
		log.Fatal(err)
	}

	return nil
}
