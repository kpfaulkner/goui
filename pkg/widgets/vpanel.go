package widgets

import (
	"image/color"
)

// VPanel stacks internal objects vertically.... starting at top and going down... so umm reverse panel :)
type VPanel struct {
	Panel

	// Y offset for next widget to be added.
	YLoc float64
}

func init() {
	defaultPanelColour = color.RGBA{0xff, 0x00, 0x00, 0xff}
}

func NewVPanel(ID string, width int, height int, colour *color.RGBA) *VPanel {
	p := VPanel{}
	p.Panel = *NewPanel(ID, width, height, colour)
	return &p
}


// AddWidget adds a widget to the panel, but each widget is below the next one.
func (p *VPanel) AddWidget(w IWidget) error {
	w.AddParentPanel(p)

	// find X,Y for widget...
  w.SetXY(p.X,p.YLoc)
	_,height := w.GetSize()
	p.YLoc += height
	p.widgets = append(p.widgets, w)
	return nil
}

