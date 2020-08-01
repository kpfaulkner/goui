package events

const (

	// raw events that will start at window and work their way down.
	// eg, mouse clicked, key pressed on keyboard etc.
	EventTypeButtonDown int = iota
	EventTypeButtonUp
	EventTypeKeyboard


	// events generated from widgets indicating a widget based event
	// eg. image button pressed.
	// Some of these might seem replicated... eg, mouse clicked (on a button) could be
	// seen as the same as "image button pressed"... but one is effectively HW signal coming in
	// where as the the widget events are out going (at least, that's the thought at the moment)
	WidgetEventTypeButtonPressed  int = iota
)

type IEventHandler interface {
	HandleEvent(event IEvent) error
}

type IEvent interface {
	Name() string
	EventType() int
	Action() error
}

type Event struct {
	eventName string
	eventType int
}

func NewEvent(eventType int) Event {
	e := Event{}
	return e
}
