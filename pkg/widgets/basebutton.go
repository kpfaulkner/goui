package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg/events"
)

type BaseButton struct {
	BaseWidget
}

func NewBaseButton(x float64, y float64, width int, height int) BaseButton {
	bb := BaseButton{}
	bb.BaseWidget = NewBaseWidget(x, y, width, height)
	return bb
}

func (b *BaseButton) Draw(screen *ebiten.Image) error {
	return nil
}

func (b *BaseButton) HandleEvent(event events.IEvent) error {

	eventType := event.EventType()
	switch eventType {
	case events.EventTypeButtonDown, events.EventTypeButtonUp:
		{
			mouseEvent := event.(events.MouseEvent)

			// check click is in button boundary.
			if b.ContainsCoords(mouseEvent.X, mouseEvent.Y) {
				// do generic button stuff here.

				// then do application specific stuff!!
				if ev,ok := b.eventRegister[event.EventType()]; ok {
					ev(event)
				}
			}
		}
	}
	return nil
}

type IButton interface {
	HandleEvent(event events.IEvent) error
	Draw(screen *ebiten.Image) error
}
