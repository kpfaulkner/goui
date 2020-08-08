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

	haveMenuBar bool
	haveToolBar bool

	// These are other widgets/components that are listening to THiS widget. Ie we will broadcast to them!
	eventListeners map[int][]chan events.IEvent

	// incoming events to THIS widget (ie stuff we're listening to!)
	incomingEvents chan events.IEvent

	// widget that has focus...  I think that will do?
	FocusedWidget *widgets.IWidget
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

	w.eventListeners = make(map[int][]chan events.IEvent)
	w.incomingEvents = make(chan events.IEvent, 1000) // too much?

	/*
		if w.haveMenuBar {
			mb := *widgets.NewMenuBar("menubar",  width, 30, &color.RGBA{0x71, 0x71, 0x71, 0xff})
			mb.AddMenuHeading("test")
			w.AddPanel(&mb)
		} */

	// if have toolbar then add a vpanel in and populate toolbar at top most part of vpanel

	/*
	if w.haveToolBar {

		// force toolpanel to have some dimension?
		tb := widgets.NewToolBar("toolbar", &color.RGBA{0, 0, 0xff, 0xff})
		tb.SetSize(w.width, 40)

		vp := widgets.NewVPanel("main vpanel", nil)

		vp.AddWidget(tb)
		w.AddPanel(vp)
	}*/
	return w
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

// FindWidgetForInput
// Need to make recursive for panels in panels etc... but just leave pretty linear for now.
func (w *Window) FindWidgetForInputOLD(x float64, y float64) (*widgets.IWidget, error) {

	// all things are panels at this level.
	for _, panel := range w.panels {
		if panel.ContainsCoords(x, y) {

			for _, subPanel := range panel.ListPanels() {
				if subPanel.ContainsCoords(x, y) {
					for _, widget := range subPanel.ListWidgets() {
						if widget.ContainsCoords(x, y) {
							// have match
							w.FocusedWidget = &widget
							return &widget, nil
						}
					}
				}
			}

			// check widgets in panel.
			for _, widget := range panel.ListWidgets() {
				if widget.ContainsCoords(x, y) {
					// have match
					w.FocusedWidget = &widget
					return &widget, nil
				}
			}
		}
	}

	//return nil, errors.New("Unable to panel/widget that was clicked on....  impossible!!!")
	return nil, nil
}

/////////////////////// EBiten specifics below... /////////////////////////////////////////////
func (w *Window) Update(screen *ebiten.Image) error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		w.leftMouseButtonPressed = true
		x, y := ebiten.CursorPosition()
		usedWidget, err := w.FindWidgetForInput(float64(x), float64(y))
		if err != nil {
			log.Errorf("Unable to find widget for click!! %s", err.Error())
		}

		if usedWidget != nil {
			me := events.NewMouseEvent(fmt.Sprintf("widget %s button down", (*usedWidget).GetID()), x, y, events.EventTypeButtonDown, (*usedWidget).GetID())
			(*usedWidget).HandleEvent(me)
		}
	} else {
		if w.leftMouseButtonPressed {
			w.leftMouseButtonPressed = false

			x, y := ebiten.CursorPosition()
			usedWidget, err := w.FindWidgetForInput(float64(x), float64(y))
			if err != nil {
				log.Errorf("Unable to find widget for click!! %s", err.Error())
			}

			if usedWidget != nil {
				me := events.NewMouseEvent(fmt.Sprintf("widget %s button up", (*usedWidget).GetID()), x, y, events.EventTypeButtonUp,(*usedWidget).GetID())
				//w.EmitEvent(me)
				//(*usedWidget).HandleEvent(me)
				(*usedWidget).HandleEvent(me)
			}

			// it *WAS* pressed previous frame... but isn't now... this means released!!!
			//x, y := ebiten.CursorPosition()
			//me := events.NewMouseEvent(x, y, events.EventTypeButtonUp)
			//w.EmitEvent(me)
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
