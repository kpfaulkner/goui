package widgets

import (
	"errors"
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

	Disabled bool

	// Every widget will have its own image.
	// All child images will draw TO this image.
	// Might be really inefficient or might be ok, will experiment. TODO(kpfaulkner) confirm if perf is ok.
	rectImage *ebiten.Image

	// has focus....  so events should go to it?
	hasFocus bool

	// has it changed?
	stateChangedSinceLastDraw bool

	// These are other widgets/components that are listening to THiS widget. Ie we will broadcast to them!
	eventListeners map[int][]chan events.IEvent

	// incoming events to THIS widget (ie stuff we're listening to!)
	incomingEvents chan events.IEvent

	// direct parent of window.... hack to sort out mouse positioning...
	TopLevel bool

	eventHandler func(event events.IEvent) (bool, error)
}

func NewBaseWidget(ID string, x float64, y float64, width int, height int) *BaseWidget {
	bw := BaseWidget{}
	bw.ID = ID
	bw.X = x
	bw.Y = y
	bw.Width = width
	bw.Height = height
	bw.Disabled = false
	bw.rectImage, _ = ebiten.NewImage(width, height, ebiten.FilterDefault)
	bw.hasFocus = false
	bw.eventListeners = make(map[int][]chan events.IEvent)
	bw.incomingEvents = make(chan events.IEvent, 1000) // too much?
	bw.TopLevel = false

	// just go off and listen for all events.
	//go bw.ListenToIncomingEvents()

	return &bw
}

func (b *BaseWidget) GetEventListenerChannel() chan events.IEvent {
	return b.incomingEvents
}

func (b *BaseWidget) AddEventListener(eventType int, ch chan events.IEvent) error {
	if _, ok := b.eventListeners[eventType]; ok {
		b.eventListeners[eventType] = append(b.eventListeners[eventType], ch)
	} else {
		b.eventListeners[eventType] = []chan events.IEvent{ch}
	}

	return nil
}

func (b *BaseWidget) RemoveEventListener(eventType int, ch chan events.IEvent) error {
	if _, ok := b.eventListeners[eventType]; ok {
		for i := range b.eventListeners[eventType] {
			if b.eventListeners[eventType][i] == ch {
				b.eventListeners[eventType] = append(b.eventListeners[eventType][:i], b.eventListeners[eventType][i+1:]...)
				break
			}
		}
	}
	return nil
}

func isMouseEvent(event events.IEvent) bool {

	if event.EventType() == events.EventTypeButtonDown ||
		event.EventType() == events.EventTypeButtonUp  {
		return true
	}

	return false
}

// Emit event for  all listeners to receive
func (b *BaseWidget) EmitEvent(event events.IEvent) error {

	eventToUse := event

	// if event is mouse related, then convert co-ords to LOCAL (ie panel) co-ords for all listeners/children.
	if isMouseEvent(event) {
		mouseEvent := event.(events.MouseEvent)
		eventToUse = b.GenerateLocalCoordMouseEvent(mouseEvent)
	}

	if _, ok := b.eventListeners[eventToUse.EventType()]; ok {
		for _, handler := range b.eventListeners[eventToUse.EventType()] {
			go func(handler chan events.IEvent) {
				handler <- eventToUse
			}(handler)
		}
	}

	return nil
}

func (b *BaseWidget) Draw(screen *ebiten.Image) error {
	return nil
}

func (b *BaseWidget) ListenToIncomingEvents() error {

	for {
		ev := <-b.incomingEvents
		// do something with event!
		log.Debugf("EVENT %v : type %d\n", ev, ev.EventType())

		// do our local event processing (HandleEvent) then pass onto other listeners (assuming order would be important here).
	//	used, err := b.HandleEvent(ev)
	  used, err := b.eventHandler(ev)
	  //used, err := eventHandler(ev)
	  //used, err := w.HandleEvent(ev)
		if err != nil {
			log.Errorf("Unable to HandleEvent from widget: %s", err.Error())
			continue
		}

		// if USED by this widget... then pass it onto the child widgets.
		// if NOT used by this widget.... its nothing to do with us... dont
		// propagate.
		if used {
			// if mouse event, convert to local co-ord system?
			err := b.EmitEvent(ev)
			if err != nil {
				log.Errorf("Unable to emit event from widget: %s", err.Error())
				// wont break out here... assuming/hoping that this is just a once off :)
			}
		}
	}
	return nil
}

func (b *BaseWidget) HandleEventXX(event events.IEvent) (bool, error) {

	// shouldn't be used.
	return false, nil
}

// BroadcastEvent signals back to main application that something has happened.
// unsure if actually needed, but see the probability of it.
func (b *BaseWidget) BroadcastEvent(event events.IEvent) error {
	return errors.New("BaseWidget shouldn't broadcast events!")
}

// ContainsCoords determines if co-ordinates (based off parent!)
func (b *BaseWidget) ContainsCoords(x float64, y float64) bool {

	localX := x
	localY := y

	// if top level, dont try and shift offset.
	if !b.TopLevel {
		localX = x - b.X
		localY = y - b.Y
	}
	return localX >= 0 && localX <= b.X+float64(b.Width) && localY >= 0 && localY <= b.Y+float64(b.Height)
}

// GenerateLocalCoordMouseEvent takes an incoming MouseEvent and converts the X,Y co-ordinates
// to something relative to the current widget. The incoming MouseEvent co-ords are relative to
// the parent.
func (b *BaseWidget) GenerateLocalCoordMouseEvent(incomingEvent events.MouseEvent) events.MouseEvent {
	localX := incomingEvent.X - b.X
	localY := incomingEvent.Y - b.Y
	outgoingMouseEvent := events.NewMouseEvent(int(localX), int(localY), incomingEvent.EventType())
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

// IWidget defines what actions can be performed on a widget.
// Hate using the I* notation... but will do for now.
type IWidget interface {
	Draw(screen *ebiten.Image) error
	HandleEvent(event events.IEvent) (bool, error)
	BroadcastEvent(event events.IEvent) error
	GetData() (interface{}, error) // absolutely HATE the empty interface, but this will need to be extremely generic I suppose?
}
