package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg/events"
	"image/color"
)

var defaultPanelColour color.RGBA

type IPanel interface {
	AddWidget(w IWidget) error
	Draw(screen *ebiten.Image) error
	HandleEvent(event events.IEvent) (bool, error)
	SetTopLevel(bool)
}

// Panel has a position, width and height.
// Panels contain other widgets (and also other panels).
type Panel struct {
	BaseWidget

	// panel can contain panels.
	panels []Panel

	// all widgets.... see if we can be super generic here.
	widgets []IWidget

	panelColour color.RGBA
}

func init() {
	defaultPanelColour = color.RGBA{0xff, 0x00, 0x00, 0xff}
}

func NewPanel(ID string, x float64, y float64, width int, height int, colour *color.RGBA) *Panel {
	p := Panel{}
	p.BaseWidget = *NewBaseWidget(ID, x, y, width, height)

	if colour != nil {
		p.panelColour = *colour
	} else {
		p.panelColour = defaultPanelColour
	}

	p.eventHandler = p.HandleEvent

	// just go off and listen for all events.
	go p.ListenToIncomingEvents()

	return &p
}

// AddWidget adds a already created checkbox.
func (p *Panel) AddWidget(w IWidget) error {
	p.widgets = append(p.widgets, w)
	return nil
}

// Draw renders all the widgets inside the panel (and the panel itself.. .if there is anything to it?)
func (p *Panel) Draw(screen *ebiten.Image) error {

	// colour background of panel first, just so we can see it.
	_ = p.rectImage.Fill(p.panelColour)

	for _, w := range p.widgets {
		w.Draw(p.rectImage)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	_ = screen.DrawImage(p.rectImage, op)
	return nil
}

// HandlEvent returns true if related to this panel. eg, mouse was in its borders etc.
// Keyboard... well, will accept anyway (need to figure out focusing for that).
func (p *Panel) HandleEvent(event events.IEvent) (bool, error) {

	inPanel := false

	eventType := event.EventType()
	switch eventType {
	case events.EventTypeButtonDown:
		{
			inPanel, _ = p.HandleMouseEvent(event)
		}
	case events.EventTypeButtonUp:
		{
			inPanel, _ = p.HandleMouseEvent(event)
		}

	case events.EventTypeKeyboard:
		{
			inPanel, _ = p.HandleKeyboardEvent(event)
		}
	}

	return inPanel, nil
}

func (p *Panel) SetTopLevel(topLevel bool){
  p.TopLevel = topLevel
}

func (p *Panel) HandleMouseEvent(event events.IEvent) (bool, error) {
	inPanel, _ := p.BaseWidget.CheckMouseEventCoords(event)

	if inPanel {

    p.hasFocus = true
		// in theory do any mouse related stuff specific to the panel....
		/*
		mouseEvent := event.(events.MouseEvent)
		log.Debugf("HandleMouseEvent panel %s :  %f %f", p.ID, mouseEvent.X, mouseEvent.Y)
		localCoordMouseEvent := p.GenerateLocalCoordMouseEvent(mouseEvent)

		for _, widget := range p.widgets {
			widget.HandleEvent(localCoordMouseEvent)
		} */
	} else {
		p.hasFocus = false
	}

	// should propagate to children nodes?
	return inPanel, nil
}

func (p *Panel) HandleKeyboardEvent(event events.IEvent) (bool, error) {

	/*
	keyboardEvent := event.(events.KeyboardEvent)
	for _, widget := range p.widgets {
		widget.HandleEvent(keyboardEvent)
	}
 */

	// only propagate to children if this panel had focus.
	return p.hasFocus, nil
}
