package events

type MouseEvent struct {
	Event
	X float64
	Y float64

}

func NewMouseEvent(x int, y int) MouseEvent {
	e := MouseEvent{}
	e.X = float64(x)
	e.Y = float64(y)
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
