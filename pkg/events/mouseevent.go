package events

type MouseEvent struct {
	Event
	X int
	Y int

}

func NewMouseEvent(x int, y int) MouseEvent {
	e := MouseEvent{}
	e.X = x
	e.Y = x
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
