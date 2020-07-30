package events

type KeyboardEvent struct {
	Event
	Character string // yeah yeah, shouldn't really be a string.
}

func NewKeyboardEvent(eventType int) KeyboardEvent {
	e := KeyboardEvent{}
	e.eventType = eventType
	return e
}

func (e KeyboardEvent) Name() string {
	return e.eventName
}

func (e KeyboardEvent) EventType() int {
	return e.eventType
}

func (e KeyboardEvent) Action() error {
	return nil
}
