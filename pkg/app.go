package pkg

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg/events"
	"github.com/kpfaulkner/goui/pkg/widgets"
	log "github.com/sirupsen/logrus"
	"image/color"
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
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		//log.Debugf("mouse click %d %d", x,y)
		me := events.NewMouseEvent(x,y)
		w.HandleEvent(me)
	}
	return nil
}

func (w *Window) HandleButtonUpEvent(event events.MouseEvent) error {
	//log.Debugf("button up %f %f", event.X, event.Y)
	return nil
}

func (w *Window) HandleButtonDownEvent(event events.MouseEvent) error {
	//log.Debugf("button down %f %f", event.X, event.Y)

	// loop through panels and find a target!
	for _,panel := range w.panels {
		panel.HandleEvent(event)
	}

  return nil
}


func (w *Window) HandleEvent(event events.IEvent) error {
	//log.Debugf("Window handled event %v", event)

	switch event.EventType() {
	case events.EventTypeButtonUp: {
		err := w.HandleButtonUpEvent(event.(events.MouseEvent))
		return err
	}

	case events.EventTypeButtonDown: {
		err := w.HandleButtonDownEvent(event.(events.MouseEvent))
		return err
	}
	}

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
