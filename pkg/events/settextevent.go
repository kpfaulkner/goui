package events

type SetTextEvent struct {
	Event
	Text string
}

func NewSetTextEvent(text string) SetTextEvent {
	e := SetTextEvent{}
	e.eventType = EventTypeSetText
	e.Text = text

	return e
}

func (e SetTextEvent) Name() string {
	return e.eventName
}

func (e SetTextEvent) EventType() int {
	return e.eventType
}

func (e SetTextEvent) Action() error {
	return nil
}
