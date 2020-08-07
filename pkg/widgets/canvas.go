package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg/events"
)

// Canvas is just an image that will be updated outside of the UI framework.
// eg client could play a video to it (frame by frame)... etc. Just to testing out/playing
type Canvas struct {
	BaseWidget
}

func NewCanvas(ID string, width int, height int) *Canvas {
	b := Canvas{}
	b.BaseWidget = *NewBaseWidget(ID, width, height, nil)
	return &b
}

func (b *Canvas) GetUnderlyingImage() *ebiten.Image {
	return b.rectImage
}

func (b *Canvas) HandleEvent(event events.IEvent) error {
	return nil
}

// do we event need this?
func (b *Canvas) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.X, b.Y)
	_ = screen.DrawImage(b.rectImage, op)
	return nil
}
