package events

type CheckBoxEvent struct {
	Event
	Checked bool
}

func NewCheckBoxEvent(name string, eventType int, checked bool) CheckBoxEvent {
	e := CheckBoxEvent{}
	e.eventName = name
	e.eventType = eventType
	e.Checked = checked
	return e
}

func (m CheckBoxEvent) Name() string {
	return m.eventName
}

func (m CheckBoxEvent) EventType() int {
	return m.eventType
}

func (m CheckBoxEvent) Action() error {
	return nil
}
