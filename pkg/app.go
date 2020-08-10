package pkg

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/kpfaulkner/goui/pkg/common"
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
  mouseX int
	mouseY int

	haveMenuBar bool
	haveToolBar bool

	// widget that has focus...  I think that will do?
	FocusedWidget *widgets.IWidget

	// app level mouse + keyboard handlers.
	// Sometimes we want the app to get the event and not just a specific widget.
	// Example, if making a calculator app and I type '1' on the keyboard, I want the
	// app to react even though the last widget I clicked on might have been a '9'.
	// I this a hack or should I be having events propagate to all parents regardless?
	keyboardHandler func(event events.KeyboardEvent) error
	mouseHandler    func(event events.MouseEvent) error


}

var defaultFontInfo common.Font

func init() {
	defaultFontInfo = common.LoadFont("", 16, color.RGBA{0xff, 0xff, 0xff, 0xff})
}

func NewWindow(width int, height int, title string, haveMenuBar bool, haveToolBar bool) Window {
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
	w.haveToolBar = haveToolBar

	return w
}

func (w *Window) AddKeyboardHandler(handler func(event events.KeyboardEvent) error) error {
	w.keyboardHandler = handler
	return nil
}

func (w *Window) AddMouseHandler(handler func(event events.MouseEvent) error) error {
	w.mouseHandler = handler
	return nil
}

func (w *Window) AddPanel(panel widgets.IPanel) error {
	panel.SetTopLevel(true)
	panel.SetSize(w.width, w.height)
	w.panels = append(w.panels, panel)

	return nil
}

func (w *Window) FindWidgetRecursive(x float64, y float64, widget widgets.IWidget) widgets.IWidget {

	if widget == nil {
		return nil
	}

	// check recursive.
	panel, ok := widget.(widgets.IPanel)
	if ok {
		for _, ww := range panel.ListWidgets() {
			res := w.FindWidgetRecursive(x, y, ww)
			if res != nil {
				return res
			}
		}
	} else {
		if widget.ContainsCoords(x, y) {
			// have match
			w.FocusedWidget = &widget
			return widget
		}
	}
	return nil
}

// FindWidgetForInput
// Need to make recursive for panels in panels etc... but just leave pretty linear for now.
func (w *Window) FindWidgetForInput(x float64, y float64) (*widgets.IWidget, error) {

	// all things are panels at this level.
	for _, panel := range w.panels {
		if panel.ContainsCoords(x, y) {

			for _, widget := range panel.ListWidgets() {
				res := w.FindWidgetRecursive(x, y, widget)
				if res != nil {
					return &res, nil
				}
			}
		}
	}
	//return nil, errors.New("Unable to panel/widget that was clicked on....  impossible!!!")
	return nil, nil
}


/////////////////////// EBiten specifics below... /////////////////////////////////////////////
func (w *Window) Update(screen *ebiten.Image) error {

	x, y := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {

		// only react to mouse button down if we're not already aware of it.
		if !w.leftMouseButtonPressed {
			w.leftMouseButtonPressed = true
			usedWidget, err := w.FindWidgetForInput(float64(x), float64(y))
			if err != nil {
				log.Errorf("Unable to find widget for click!! %s", err.Error())
			}

			if usedWidget != nil {
				me := events.NewMouseEvent(fmt.Sprintf("widget %s button down", (*usedWidget).GetID()), x, y, events.EventTypeButtonDown, (*usedWidget).GetID())
				(*usedWidget).HandleEvent(me)
			}

			if w.mouseHandler != nil {
				me := events.NewMouseEvent("button down", x, y, events.EventTypeButtonDown, "")
				w.mouseHandler(me)
			}
		}
	} else {
		if w.leftMouseButtonPressed {
			w.leftMouseButtonPressed = false
			usedWidget, err := w.FindWidgetForInput(float64(x), float64(y))
			if err != nil {
				log.Errorf("Unable to find widget for click!! %s", err.Error())
			}

			if usedWidget != nil {
				me := events.NewMouseEvent(fmt.Sprintf("widget %s button up", (*usedWidget).GetID()), x, y, events.EventTypeButtonUp, (*usedWidget).GetID())
				//w.EmitEvent(me)
				//(*usedWidget).HandleEvent(me)
				(*usedWidget).HandleEvent(me)
			}

			if w.mouseHandler != nil {
				me := events.NewMouseEvent("button upXX", x, y, events.EventTypeButtonUp,"")
				w.mouseHandler(me)
			}
		}
	}

	// check if mouse has moved
	if x != w.mouseX || y != w.mouseY {
		w.mouseX = x
		w.mouseY = y
		// call app level mouse event handler
		if w.mouseHandler != nil {
			me := events.NewMouseEvent("mouse movement", x, y, events.EventTypeMouseMove,"")
			w.mouseHandler(me)
		}
	}

	inp := ebiten.InputChars()
	if len(inp) > 0 {
		// create keyboard event

		//w.EmitEvent(ke)
		x, y := ebiten.CursorPosition()
		usedWidget, err := w.FindWidgetForInput(float64(x), float64(y))
		if err != nil {
			log.Errorf("Unable to find widget for click!! %s", err.Error())
		}

		if usedWidget != nil {
			//ke := events.NewMouseEvent(x,y, events.EventTypeKeyboard)
			ke := events.NewKeyboardEvent(ebiten.Key(inp[0]), (*usedWidget).GetID()) // only send first one?
			//w.EmitEvent(me)
			(*usedWidget).HandleEvent(ke)
		}

		// go to app level handler. hack?
		if w.keyboardHandler != nil {
			ke := events.NewKeyboardEvent(ebiten.Key(inp[0]), "") // only send first one?
			w.keyboardHandler(ke)
		}
	}

	// If the backspace key is pressed, remove one character.
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		x, y := ebiten.CursorPosition()
		usedWidget, err := w.FindWidgetForInput(float64(x), float64(y))
		if err != nil {
			log.Errorf("Unable to find widget for click!! %s", err.Error())
		}

		if usedWidget != nil {
			//ke := events.NewMouseEvent(x,y, events.EventTypeKeyboard)
			ke := events.NewKeyboardEvent(ebiten.KeyBackspace, (*usedWidget).GetID())
			//w.EmitEvent(me)
			//(*usedWidget).HandleEvent(ke)
			(*usedWidget).HandleEvent(ke)
		}

		// go to app level handler. hack?
		if w.keyboardHandler != nil {
			ke := events.NewKeyboardEvent(ebiten.KeyBackspace, "") // only send first one?
			w.keyboardHandler(ke)
		}
	}

	return nil
}

func (w *Window) HandleButtonUpEvent(event events.MouseEvent) error {
	log.Debugf("button up %f %f", event.X, event.Y)
	for _, panel := range w.panels {
		panel.HandleEvent(event)
	}

	return nil
}

func (w *Window) HandleButtonDownEvent(event events.MouseEvent) error {
	log.Debugf("button down %f %f", event.X, event.Y)

	// loop through panels and find a target!
	for _, panel := range w.panels {
		panel.HandleEvent(event)
	}

	return nil
}

func (w *Window) HandleKeyboardEvent(event events.KeyboardEvent) error {

	// loop through panels and find a target!
	for _, panel := range w.panels {
		panel.HandleEvent(event)
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

	x, y := ebiten.CursorPosition()
	//defaultFontInfo := common.LoadFont("", 16, color.RGBA{0xff, 0xff, 0xff, 0xff})
	text.Draw(screen, fmt.Sprintf("%d %d", x, y), defaultFontInfo.UIFont, 00, 500, color.White)

}

func (w *Window) reset(outsideWidth, outsideHeight int) error {

	return nil
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	//  log.Debugf("LAYOUT %d %d", outsideWidth, outsideHeight)

	if outsideWidth != w.width || outsideHeight != w.height {
		// trigger recalculation of any panels etc.

		w.width = outsideWidth
		w.height = outsideHeight
		w.panels[0].SetSize(outsideWidth, outsideHeight)

		// FIXME(kpfaulkner)  need to trigger resize!!!
		return outsideWidth, outsideHeight

	}

	return outsideWidth, outsideHeight
	//return w.width, w.height
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
