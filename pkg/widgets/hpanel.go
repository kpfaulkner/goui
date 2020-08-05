package widgets

import (
	"image/color"
)

// HPanel stacks internal objects horizonally... left to right.
type HPanel struct {
	Panel

	// X offset for next widget to be added.
	XLoc float64
}

func NewHPanel(ID string, colour *color.RGBA) *HPanel {
	p := HPanel{}
	p.Panel = *NewPanel(ID, colour)
	return &p
}

// AddWidget adds a widget to the panel, but each widget is to the right of the previous one.
func (p *HPanel) AddWidget(w IWidget) error {
	w.AddParentPanel(p)

	// find X,Y for widget...
	w.SetXY(p.XLoc, p.Y)
	width, height := w.GetSize()

	// grow panel height if widget is taller.
	if height > float64(p.Height) {
		p.Height = int(height)
		p.SetSize(p.Width,p.Height)
	}

	if p.XLoc + width > float64(p.Width) {
		p.Width = int(p.XLoc + width)
		p.SetSize(p.Width,p.Height)
	}

	p.XLoc += width

	p.widgets = append(p.widgets, w)
	return nil
}
