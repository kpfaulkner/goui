package widgets

import (
	"image/color"
)

// Toolbar stacks internal objects horizonally... left to right.
type Toolbar struct {
	HPanel

	// X offset for next widget to be added.
	XLoc float64
}

func NewToolBar(ID string, colour *color.RGBA) *Toolbar {
	p := Toolbar{}
	p.HPanel = *NewHPanel(ID, colour)
	return &p
}

// AddToolBarItem add toolbaritem.
// Basically just an alias to HPanel.AddWidget for now... but that could change.
func (p *Toolbar) AddToolBarItem(i IWidget) error {
	p.AddWidget(i)
	return nil
}
