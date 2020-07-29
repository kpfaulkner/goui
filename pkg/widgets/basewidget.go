package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
)

// BaseWidget is the common element of ALL disable items. Widgets, buttons, panels etc.
// Should only handle some basic items like location, width, height and possibly events.
type BaseWidget struct {

	// Location of top left of widget.
	// This is always relative to the parent widget.
	X float64
	Y float64

	Width  int
	Height int

	Disabled bool

	// Every widget will have its own image.
	// All child images will draw TO this image.
	// Might be really inefficient or might be ok, will experiment. TODO(kpfaulkner) confirm if perf is ok.
	rectImage *ebiten.Image

	// register event types with widgets?
	eventRegister map[int]func(event events.IEvent) error
}

func NewBaseWidget(x float64, y float64, width int, height int) BaseWidget {
	bw := BaseWidget{}
	bw.X = x
	bw.Y = y
	bw.Width = width
	bw.Height = height
	bw.Disabled = false
	bw.rectImage, _ = ebiten.NewImage(width, height, ebiten.FilterDefault)
	bw.eventRegister = make(map[int]func(event events.IEvent) error)
	return bw
}

func (b *BaseWidget) RegisterEventHandler(eventType int, eventHandler func(events.IEvent) error) error {
	b.eventRegister[eventType] = eventHandler
	return nil
}

func (b *BaseWidget) Draw(screen *ebiten.Image) error {
	return nil
}

func (b *BaseWidget) HandleEvent(event events.IEvent) error {

	eventType := event.EventType()
	if handler,ok := b.eventRegister[eventType]; ok {
		handler(event)
	}
	return nil
}

// ContainsCoords determines if co-ordinates (based off parent!)
func (b *BaseWidget) ContainsCoords(x float64, y float64) bool {
	localX := x - b.X
	localY := y - b.Y
	return localX >= 0 && localX <= b.X + float64(b.Width) && localY >= 0 && localY <= b.Y + float64(b.Height)
}

// GenerateLocalCoordMouseEvent takes an incoming MouseEvent and converts the X,Y co-ordinates
// to something relative to the current widget. The incoming MouseEvent co-ords are relative to
// the parent.
func (b *BaseWidget) GenerateLocalCoordMouseEvent(incomingEvent events.MouseEvent) events.MouseEvent {
	localX := incomingEvent.X - b.X
	localY := incomingEvent.Y - b.Y
	outgoingMouseEvent := events.NewMouseEvent(int(localX), int(localY))
	return outgoingMouseEvent
}

func (b *BaseWidget) HandleMouseEvent(event events.IEvent) (bool, error) {
	mouseEvent := event.(events.MouseEvent)

	if b.ContainsCoords(mouseEvent.X, mouseEvent.Y) {
		log.Debugf("basewidget handled mouse event at %f %f", mouseEvent.X, mouseEvent.Y)
		return true, nil
	}

	// should propagate to children nodes?
	return false, nil
}



// IWidget defines what actions can be performed on a widget.
// Hate using the I* notation... but will do for now.
type IWidget interface {
	Draw(screen *ebiten.Image) error

	HandleEvent(event events.IEvent) error
}
