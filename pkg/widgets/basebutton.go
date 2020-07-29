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

type IButton interface {
	HandleEvent(event events.IEvent) error
	Draw(screen *ebiten.Image) error
}
