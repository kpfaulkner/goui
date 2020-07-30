package events

import "github.com/hajimehoshi/ebiten"

type KeyboardEvent struct {
	Event
	Character ebiten.Key // yeah yeah, shouldn't really be a string.
}

func NewKeyboardEvent(key ebiten.Key) KeyboardEvent {
	e := KeyboardEvent{}
	e.eventType = EventTypeKeyboard
	e.Character = key

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
