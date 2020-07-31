package pkg

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
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
	panels []widgets.IPanel

	leftMouseButtonPressed  bool
	rightMouseButtonPressed bool

	haveMenuBar bool
}

func NewWindow(width int, height int, title string, haveMenuBar bool) Window {
	w := Window{}
	w.height = height
	w.width = width
	w.title = title

	// panels are ordered. They are drawn from first to last.
	// So if we *have* to have a panel drawn last (eg, menu from a menu bar) then
	// one approach might be to create a panel (representing the menu) and it gets displayed at the end?
	// Mad idea..
	w.panels = []widgets.IPanel{}
	w.leftMouseButtonPressed = false
	w.rightMouseButtonPressed = false
	w.haveMenuBar = haveMenuBar

	if w.haveMenuBar {
		mb := widgets.NewMenuBar("menubar", 0, 0, width, 30, &color.RGBA{0x71, 0x71, 0x71, 0xff})
		mb.AddMenuHeading("test")
		w.AddPanel(&mb)
	}
	return w
}

func (w *Window) AddPanel(panel widgets.IPanel) error {
	w.panels = append(w.panels, panel)
	return nil
}

/////////////////////// EBiten specifics below... /////////////////////////////////////////////
func (w *Window) Update(screen *ebiten.Image) error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		w.leftMouseButtonPressed = true
		x, y := ebiten.CursorPosition()
		me := events.NewMouseEvent(x, y, events.EventTypeButtonDown)
		w.HandleEvent(me)
	} else {
		if w.leftMouseButtonPressed {
			w.leftMouseButtonPressed = false
			// it *WAS* pressed previous frame... but isn't now... this means released!!!
			x, y := ebiten.CursorPosition()
			me := events.NewMouseEvent(x, y, events.EventTypeButtonUp)
			w.HandleEvent(me)
		}
	}

	inp := ebiten.InputChars()
	if len(inp) > 0 {
		// create keyboard event
		ke := events.NewKeyboardEvent(ebiten.Key(inp[0])) // only send first one?
		w.HandleEvent(ke)
	}

	// If the backspace key is pressed, remove one character.
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		ke := events.NewKeyboardEvent(ebiten.KeyBackspace)
		w.HandleEvent(ke)
	}

	return nil
}

func (w *Window) HandleButtonUpEvent(event events.MouseEvent) error {
	log.Debugf("button up %f %f", event.X, event.Y)
	for _, panel := range w.panels {
		panel.HandleEvent(event, false)
	}

	return nil
}

func (w *Window) HandleButtonDownEvent(event events.MouseEvent) error {
	log.Debugf("button down %f %f", event.X, event.Y)

	// loop through panels and find a target!
	for _, panel := range w.panels {
		panel.HandleEvent(event, false)
	}

	return nil
}

func (w *Window) HandleKeyboardEvent(event events.KeyboardEvent) error {

	// loop through panels and find a target!
	for _, panel := range w.panels {
		panel.HandleEvent(event, false)
	}
	return nil
}

func (w *Window) HandleEvent(event events.IEvent) error {
	//log.Debugf("Window handled event %v", event)

	switch event.EventType() {
	case events.EventTypeButtonUp:
		{
			err := w.HandleButtonUpEvent(event.(events.MouseEvent))
			return err
		}

	case events.EventTypeButtonDown:
		{
			err := w.HandleButtonDownEvent(event.(events.MouseEvent))
			return err
		}

	case events.EventTypeKeyboard:
		{
			err := w.HandleKeyboardEvent(event.(events.KeyboardEvent))
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

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}
