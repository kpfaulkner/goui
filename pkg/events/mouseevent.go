package events

type MouseEvent struct {
	Event
	X float64
	Y float64
}

func NewMouseEvent(name string, x int, y int, eventType int, widgetID string) MouseEvent {
	e := MouseEvent{}
	e.X = float64(x)
	e.Y = float64(y)
	e.eventName = name
	e.eventType = eventType
	e.widgetID = widgetID
	return e
}

func (m MouseEvent) Name() string {
	return m.eventName
}

func (m MouseEvent) EventType() int {
	return m.eventType
}

func (m MouseEvent) Action() error {
	return nil
}
