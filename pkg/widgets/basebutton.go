package widgets

import (
	"github.com/hajimehoshi/ebiten"
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
	OnPress() error
	Draw(screen *ebiten.Image) error
}
