package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
	"image/color"
	"math/rand"
)

var defaultPanelColour color.RGBA

type IPanel interface {
	AddWidget(w IWidget) error
	Draw(screen *ebiten.Image) error
	HandleEvent(event events.IEvent) error
	SetTopLevel(bool)
	SetSize(width int, height int) error
	GetSize() (float64, float64)
	ContainsCoords(x float64, y float64) bool // contains co-ords... co-ords are based on immediate parents location/size.
	ListWidgets() []IWidget
	ListPanels() []Panel
	GetCoords() (float64, float64)
	GetDeltaOffset() (bool, float64, float64)
	GlobalToLocalCoords(x float64, y float64) (float64, float64)
	AddParentPanel(parentPanel IPanel) error
}

// Panel has a position, width and height.
// Panels contain other widgets (and also other panels).
type Panel struct {
	BaseWidget

	hasBorder bool

	// panel can contain panels.
	panels []Panel

	// all widgets.... see if we can be super generic here.
	widgets []IWidget

	panelColour color.RGBA
	borderColour color.RGBA

	// DynamicSize is for when we based sizes off window/panels etc. ie everything is
	// proportional as opposed to absolute.
	DynamicSize bool
}

func init() {
	defaultPanelColour = color.RGBA{0x00, 0x00, 0x00, 0xff}
}

func NewPanel(ID string, colour *color.RGBA, borderColour *color.RGBA) *Panel {
	p := Panel{}

	p.DynamicSize = true
	p.BaseWidget = *NewBaseWidget(ID, 0, 0, p.HandleEvent)

	if colour != nil {
		p.panelColour = *colour
	} else {
		p.panelColour = defaultPanelColour
	}

	if borderColour != nil {
		p.borderColour = *borderColour
		p.hasBorder = true
	} else {
		p.borderColour = color.RGBA{0,0,0xff,0xff}
		p.hasBorder = false  // yes contradictory... set default colour but no border flag. WIP.
	}
	return &p
}

func (p *Panel) ListWidgets() []IWidget {
	return p.widgets
}

func (p *Panel) ListPanels() []Panel {
	return p.panels
}

func (p *Panel) GetCoords() (float64, float64) {
	return p.X, p.Y
}

// AddWidget adds a widget to a panel and subscribes the widget
// to a number of events (generated from the panel)
func (p *Panel) AddWidget(w IWidget) error {

	// widget.
	w.AddParentPanel(p)
	p.widgets = append(p.widgets, w)

	return nil
}

func (p *Panel) generateBorder(border color.RGBA) error {

	ebitenutil.DrawLine(p.rectImage, 0, 0, float64(p.Width), 0, border)
	ebitenutil.DrawLine(p.rectImage, float64(p.Width), 0, float64(p.Width), float64(p.Height), border)
	ebitenutil.DrawLine(p.rectImage, float64(p.Width), float64(p.Height), 0, float64(p.Height), border)
	ebitenutil.DrawLine(p.rectImage, 0, float64(p.Height), 0, 0, border)
	return nil
}


// Draw renders all the widgets inside the panel (and the panel itself.. .if there is anything to it?)
func (p *Panel) Draw(screen *ebiten.Image) error {

	// colour background of panel first, just so we can see it.
	_ = p.rectImage.Fill(p.panelColour)


	for _, w := range p.widgets {
		w.Draw(p.rectImage)
	}

	if p.hasBorder {
		p.generateBorder( p.borderColour)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	_ = screen.DrawImage(p.rectImage, op)
	return nil
}

// HandlEvent returns true if related to this panel. eg, mouse was in its borders etc.
// Keyboard... well, will accept anyway (need to figure out focusing for that).
func (p *Panel) HandleEvent(event events.IEvent) error {

	log.Debugf("PanelHandler for %s", p.ID)

	eventType := event.EventType()
	switch eventType {
	case events.EventTypeButtonDown:
		{
			_, _ = p.HandleMouseEvent(event)
		}
	case events.EventTypeButtonUp:
		{
			_, _ = p.HandleMouseEvent(event)
		}

	case events.EventTypeKeyboard:
		{
			_, _ = p.HandleKeyboardEvent(event)
		}
	}

	return nil
}

func (p *Panel) SetTopLevel(topLevel bool) {
	p.TopLevel = topLevel
}

func (p *Panel) HandleMouseEvent(event events.IEvent) (bool, error) {
	inPanel, _ := p.BaseWidget.CheckMouseEventCoords(event)

	if inPanel {
		p.hasFocus = true

		r := rand.Intn(255)
		g := rand.Intn(255)
		b := rand.Intn(255)
		p.panelColour = color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}
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

func (p *Panel) GetDeltaOffset() (bool, float64, float64) {
	return p.populatedGlobalDelta, p.globalDX, p.globalDY
}

func (p *Panel) SetSize(width int, height int) error {
	p.Width = width
	p.Height = height

	// new backing image.
	p.rectImage, _ = ebiten.NewImage(width, height, ebiten.FilterDefault)

	// ie we're a panel in a panel.
	if p.parentPanel != nil {
		pw, ph := p.parentPanel.GetSize()
		newW := int(pw)
		newH := int(ph)
		if p.Width > int(pw) {
			newW = p.Width
		}
		if p.Height > int(ph) {
			newH = int(p.Height)
		}
		p.parentPanel.SetSize(newW, newH)
	}

	return nil
}
