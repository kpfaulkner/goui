package events

type DeselectEvent struct {
	Event
}

func NewDeselectEvent(widgetID string) DeselectEvent {
	e := DeselectEvent{}
	e.eventType = EventTypeDeselect
	e.widgetID = widgetID

	return e
}

func (e DeselectEvent) Name() string {
	return e.eventName
}

func (e DeselectEvent) EventType() int {
	return e.eventType
}

func (e DeselectEvent) Action() error {
	return nil
}
