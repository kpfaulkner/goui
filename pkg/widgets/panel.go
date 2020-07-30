package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
	"image/color"
)

var defaultPanelColour color.RGBA

type IPanel interface {
	AddWidget(w IWidget) error
	Draw(screen *ebiten.Image) error
	HandleEvent(event events.IEvent, local bool) error
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

func NewPanel(ID string, x float64, y float64, width int, height int, colour *color.RGBA) Panel {
	p := Panel{}
	p.BaseWidget = NewBaseWidget(ID, x, y, width, height)

	if colour != nil {
		p.panelColour = *colour
	} else {
		p.panelColour = defaultPanelColour
	}

	return p
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

func (p *Panel) HandleEvent(event events.IEvent, local bool) error {

	eventType := event.EventType()
	switch eventType {
	case events.EventTypeButtonDown:
		{
			p.HandleMouseEvent(event, local)
		}
	case events.EventTypeButtonUp:
		{
			p.HandleMouseEvent(event, local)
		}

	case events.EventTypeKeyboard:
		{
			p.HandleKeyboardEvent(event, local)
		}
	}

	return nil
}

func (p *Panel) HandleMouseEvent(event events.IEvent, local bool) error {
	inPanel, _ := p.BaseWidget.CheckMouseEventCoords(event, local)

	if inPanel {
		mouseEvent := event.(events.MouseEvent)
		log.Debugf("HandleMouseEvent panel %s :  %f %f", p.ID, mouseEvent.X, mouseEvent.Y)
		localCoordMouseEvent := p.GenerateLocalCoordMouseEvent(mouseEvent)

		for _, widget := range p.widgets {
			widget.HandleEvent(localCoordMouseEvent)
		}
	}

	// should propagate to children nodes?
	return nil
}

func (p *Panel) HandleKeyboardEvent(event events.IEvent, local bool) error {

	keyboardEvent := event.(events.KeyboardEvent)
	for _, widget := range p.widgets {
		widget.HandleEvent(keyboardEvent)
	}

	// should propagate to children nodes?
	return nil
}
