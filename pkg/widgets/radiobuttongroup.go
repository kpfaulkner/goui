package widgets

import (
	"github.com/kpfaulkner/goui/pkg/common"
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
	"image/color"
	_ "image/png"
)

const (
	selectedImage    = "images/radopbuttonselected.png"
	notSelectedImage = "images/radiobuttonnotselected.png"
)

type RadioButtonGroup struct {
	Panel

	// can be horizontal or vertical
	subPanel IPanel

	vertical bool
	fontInfo common.Font
}

func init() {
	defaultFontInfo = common.LoadFont("", 16, color.RGBA{0xff, 0xff, 0xff, 0xff})
}

func NewRadioButtonGroup(ID string, vertical bool, border bool, handler func(event events.IEvent) error) *RadioButtonGroup {
	rb := RadioButtonGroup{}
	rb.vertical = vertical
	rb.Panel = *NewPanel(ID, nil, &color.RGBA{0, 0, 0xff, 0xff})

	// create VPanel or HPanel
	// Add as a widget, this will auto handle rendering etc.
	// But also keep specific reference to rb.subPanel. Feel this might be required.
	if vertical {
		vp := NewVPanel(ID, nil)
		rb.AddWidget(vp)
		rb.subPanel = vp
	} else {
		hp := NewHPanel(ID, nil)
		rb.AddWidget(hp)
		rb.subPanel = hp
	}

	return &rb
}

func (rb *RadioButtonGroup) HandleEvent(event events.IEvent) error {

	log.Debugf("have event.... %v", event)
	eventType := event.EventType()
	switch eventType {
	case events.EventTypeButtonDown:
		{
			// deselect any button that is NOT part of the event.
			for _, w := range rb.subPanel.ListWidgets() {

				// if widget is nOT the one passed in with origin event.
				if event.WidgetID() != w.GetID() {
					ev := events.NewDeselectEvent(event.WidgetID())
					w.HandleEvent(ev)
				}
			}
		}
	}
	return nil
}

// AddRadioButton adds a new radio button to the radio button group.
// temporarily just trying the checkbox here.
func (rb *RadioButtonGroup) AddRadioButton(buttonText string) error {
	cb := NewCheckBox("cb-"+buttonText, buttonText, "images/radiobuttonselected.png", "images/radiobuttonnotselected.png", rb.HandleEvent)
	rb.subPanel.AddWidget(cb)
	return nil
}
