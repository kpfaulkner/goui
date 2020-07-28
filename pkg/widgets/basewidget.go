package widgets

import (
	"github.com/hajimehoshi/ebiten"
)

// BaseWidget is the common element of ALL disable items. Widgets, buttons, panels etc.
// Should only handle some basic items like location, width, height and possibly events.
type BaseWidget struct {

	// Location of top left of widget.
	// This is always relative to the parent widget.
	X float64
	Y float64

	Width  int
	Height int

	Disabled bool

	// Every widget will have its own image.
	// All child images will draw TO this image.
	// Might be really inefficient or might be ok, will experiment. TODO(kpfaulkner) confirm if perf is ok.
	rectImage *ebiten.Image
}

func NewBaseWidget(x float64, y float64, width int, height int) BaseWidget {
	bw := BaseWidget{}
	bw.X = x
	bw.Y = y
	bw.Width = width
	bw.Height = height
	bw.Disabled = false
	bw.rectImage, _ = ebiten.NewImage(width, height, ebiten.FilterDefault)
	return bw
}

func (b *BaseWidget) Draw(screen *ebiten.Image) error {
	return nil
}

// IWidget defines what actions can be performed on a widget.
// Hate using the I* notation... but will do for now.
type IWidget interface {
	Draw(screen *ebiten.Image) error
}
