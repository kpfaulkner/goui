package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
)

// BaseWidget is the common element of ALL disable items. Widgets, buttons, panels etc.
// Should only handle some basic items like location, width, height and possibly events.
type BaseWidget struct {
	ID string

	// Location of top left of widget.
	// This is always relative to the parent widget.
	X float64
	Y float64
	Width  int
	Height int

	// widget disabled (and shouldn't be rendered.... OR.... greyed out?
	Disabled bool

	// Every widget will have its own image.
	// All child images will draw TO this image.
	// Might be really inefficient or might be ok, will experiment. TODO(kpfaulkner) confirm if perf is ok.
	rectImage *ebiten.Image

	// has focus....  so events should go to it?
	hasFocus bool

	// has it changed?
	stateChangedSinceLastDraw bool

	// direct parent of window.... hack to sort out mouse positioning...
	TopLevel bool

	eventHandler func(event events.IEvent) error

	// global to local offset.
	globalDX             float64
	globalDY             float64
	populatedGlobalDelta bool

	// parent... to get relative positioning.
	parentPanel IPanel
}

func NewBaseWidget(ID string, width int, height int, handler func(event events.IEvent) error) *BaseWidget {
	bw := BaseWidget{}
	bw.ID = ID
	bw.X = 0
	bw.Y = 0
	bw.Width = width
	bw.Height = height
	bw.Disabled = false
	bw.rectImage, _ = ebiten.NewImage(width, height, ebiten.FilterDefault)
	bw.hasFocus = false
	bw.TopLevel = false
	bw.eventHandler = handler
	bw.populatedGlobalDelta = false // haven't asked parent for offer.
	return &bw
}

func isMouseEvent(event events.IEvent) bool {

	if event.EventType() == events.EventTypeButtonDown ||
		event.EventType() == events.EventTypeButtonUp {
		return true
	}

	return false
}

func (b *BaseWidget) Draw(screen *ebiten.Image) error {
	return nil
}

// ContainsCoords determines if co-ordinates... co-ords passed are GLOBAL
// and need to be converted.
func (b *BaseWidget) ContainsCoords(x float64, y float64) bool {
	localX, localY := b.GlobalToLocalCoords(x, y)
	//return localX >= 0 && localX <= b.X+float64(b.Width) && localY >= 0 && localY <= b.Y+float64(b.Height)
	return localX >= 0 && localX <= float64(b.Width) && localY >= 0 && localY <= float64(b.Height)
}

// GenerateLocalCoordMouseEvent takes an incoming MouseEvent and converts the X,Y co-ordinates
// to something relative to the current widget. The incoming MouseEvent co-ords are relative to
// the parent.
func (b *BaseWidget) GenerateLocalCoordMouseEvent(incomingEvent events.MouseEvent) events.MouseEvent {
	localX := incomingEvent.X - b.X
	localY := incomingEvent.Y - b.Y
	outgoingMouseEvent := events.NewMouseEvent(incomingEvent.Name(), int(localX), int(localY), incomingEvent.EventType(), b.ID)
	return outgoingMouseEvent
}

func (b *BaseWidget) CheckMouseEventCoords(event events.IEvent) (bool, error) {
	mouseEvent := event.(events.MouseEvent)

	if b.ContainsCoords(mouseEvent.X, mouseEvent.Y) {
		log.Debugf("CheckMouseEventCoords %f %f", mouseEvent.X, mouseEvent.Y)
		return true, nil
	}

	// should propagate to children nodes?
	return false, nil
}

func (b *BaseWidget) GetData() (interface{}, error) {
	return nil, nil
}

// GlobalToLocalCoords takes global co-ords and modifies it to local co-ords.
// It figures this out by keeping a global offset to base off. It gets the global offset
// by asking its parents for offset. The parent asks its parent and so on.
// THINK this should work.
// Remember LOCAL co-ords are really based off the parents co-ords...
// as in a widgets x,y is based off the parents 0,0....   man I need to explain that more clearly.
func (b *BaseWidget) GlobalToLocalCoords(x float64, y float64) (float64, float64) {

	if !b.populatedGlobalDelta {
		if b.parentPanel != nil {
			//dx,dy:= (*b.parentWidget).GetCoords()
			//_,_ = (*b.parentPanel).GlobalToLocalCoords(x,y)

			populated, dx, dy := (b.parentPanel).GetDeltaOffset()
			if !populated {
				b.parentPanel.GlobalToLocalCoords(x,y)
			}
			populated, dx, dy = (b.parentPanel).GetDeltaOffset()
			b.globalDX = dx + b.X
			b.globalDY = dy + b.Y
			b.populatedGlobalDelta = true
		} else {
			// parent is nil, just return regular x,y. This *should* be the window.
			b.globalDX = b.X
			b.globalDY = b.Y
			b.populatedGlobalDelta = true
		}
	}

	return x - b.globalDX, y - b.globalDY
}

func (b *BaseWidget) AddParentPanel(parentPanel IPanel) error {
	b.parentPanel = parentPanel
	return nil
}

func (b *BaseWidget) GetID() string {
	return b.ID
}

func (b *BaseWidget) SetXY(x float64, y float64) error {
	b.X = x
	b.Y = y
	return nil
}

func (b *BaseWidget) GetXY() (float64, float64) {
	return b.X, b.Y
}

func (b *BaseWidget) GetSize() (float64, float64) {
	return float64(b.Width), float64(b.Height)
}

// IWidget defines what actions can be performed on a widget.
// Hate using the I* notation... but will do for now.
type IWidget interface {
	Draw(screen *ebiten.Image) error
	HandleEvent(event events.IEvent) error
	GetData() (interface{}, error) // absolutely HATE the empty interface, but this will need to be extremely generic I suppose?
	ContainsCoords(x float64, y float64) bool // contains co-ords... co-ords are based on immediate parents location/size.
	GlobalToLocalCoords(x float64, y float64) (float64, float64)
	AddParentPanel(parentPanel IPanel) error
	SetXY(x float64, y float64) error
	GetXY() (float64, float64)
	GetSize() (float64, float64)
	GetID() string
}
