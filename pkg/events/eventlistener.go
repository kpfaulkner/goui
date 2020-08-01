package events


// EventListener is generic event listener that is used in the APP itself
// and not in the UI widgets etc.
type EventListener struct {

}

func NewEventListener() *EventListener {
	e := EventListener{}

	return &e
}


