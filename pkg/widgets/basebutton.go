package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
)

type BaseButton struct {
	BaseWidget
	pressed bool
}

func NewBaseButton(ID string, x float64, y float64, width int, height int) BaseButton {
	bb := BaseButton{}
	bb.BaseWidget = NewBaseWidget(ID, x, y, width, height)
	bb.pressed = false
	bb.stateChangedSinceLastDraw = true
	return bb
}

func (b *BaseButton) Draw(screen *ebiten.Image) error {
	return nil
}

func (b *BaseButton) HandleEvent(event events.IEvent) error {

	eventType := event.EventType()
	switch eventType {
	case events.EventTypeButtonDown:
		{
			mouseEvent := event.(events.MouseEvent)

			// check click is in button boundary.
			if b.ContainsCoords(mouseEvent.X, mouseEvent.Y, true) {
				b.hasFocus = true
				// if already pressed, then skip it.. .otherwise lots of repeats.
				if !b.pressed {
					b.pressed = true

					// then do application specific stuff!!
					if ev, ok := b.eventRegister[event.EventType()]; ok {
						ev(event)
					}
					b.stateChangedSinceLastDraw = true
				}
			}
		}
	case events.EventTypeButtonUp:
		{
			log.Debugf("BUTTON UP!!!")
			mouseEvent := event.(events.MouseEvent)

			// check click is in button boundary.
			if b.ContainsCoords(mouseEvent.X, mouseEvent.Y, true) {
				b.hasFocus = true

				if b.pressed {
					// do generic button stuff here.
					b.pressed = false

					// then do application specific stuff!!
					if ev, ok := b.eventRegister[event.EventType()]; ok {
						ev(event)
					}
					b.stateChangedSinceLastDraw = true
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
