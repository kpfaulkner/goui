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

func NewBaseButton(ID string, width int, height int, handler func(event events.IEvent) error) *BaseButton {
	bb := BaseButton{}
	bb.BaseWidget = *NewBaseWidget(ID, width, height, handler)
	bb.pressed = false
	bb.stateChangedSinceLastDraw = true
	//bb.eventHandler = handler

	return &bb
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
			if b.ContainsCoords(mouseEvent.X, mouseEvent.Y) {
				log.Debugf("BaseButton::HandleEvent Down %s", b.ID)

				if b.eventHandler != nil {
					b.eventHandler(event)
				}
				b.hasFocus = true
				// if already pressed, then skip it.. .otherwise lots of repeats.
				if !b.pressed {
					b.pressed = true
					b.stateChangedSinceLastDraw = true
				}
			}
		}
	case events.EventTypeButtonUp:
		{
			mouseEvent := event.(events.MouseEvent)

			// check click is in button boundary.
			if b.ContainsCoords(mouseEvent.X, mouseEvent.Y) {
				b.hasFocus = true
				if b.pressed {
					log.Debugf("BaseButton::HandleEvent Up %s", b.ID)
					// do generic button stuff here.
					b.pressed = false
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
