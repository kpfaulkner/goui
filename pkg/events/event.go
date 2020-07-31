package events

const (
	EventTypeButtonDown int = iota
	EventTypeButtonUp
	EventTypeKeyboard
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
