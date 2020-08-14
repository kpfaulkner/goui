package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg/events"
)

// EmptySpace lame way to put spaces in panels..
type EmptySpace struct {
	BaseWidget
}

func NewEmptySpace(ID string, width int, height int) *EmptySpace {
	b := EmptySpace{}
	b.BaseWidget = *NewBaseWidget(ID, width, height, nil)
	return &b
}

func (b *EmptySpace) HandleEvent(event events.IEvent) error {
	return nil
}

// do we event need this?
func (b *EmptySpace) Draw(screen *ebiten.Image) error {
	return nil
}
