package widgets

import (
	"github.com/hajimehoshi/ebiten"
)

type Button struct {
	BaseButton
}

func NewButton(imageName string, x float64, y float64, width int, height int) Button {
	b := Button{}
	b.BaseButton = NewBaseButton(x, y, width, height)
	return b
}

func (b *Button) Draw(screen *ebiten.Image) error {
	return nil
}
