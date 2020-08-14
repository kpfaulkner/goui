package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg/common"
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
	"image/color"
	_ "image/png"
	"time"
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
	// image for the checkbox
	notSelectedImage *ebiten.Image
	selectedImage    *ebiten.Image
	fontInfo         common.Font
	lastClickedTime  time.Time
}

func init() {
	defaultFontInfo = common.LoadFont("", 16, color.RGBA{0xff, 0xff, 0xff, 0xff})
}

func NewRadioButtonGroup(ID string, vertical bool, handler func(event events.IEvent) error) *RadioButtonGroup {
	rb := RadioButtonGroup{}
	rb.vertical = vertical

	/*
		img1, err := loadImage(notSelectedImage)
		if err != nil {
			log.Fatalf("Unable to load image %s", notSelectedImage)
		}
		rb.notSelectedImage = img1

		img2, err := loadImage(selectedImage)
		if err != nil {
			log.Fatalf("Unable to load image %s", selectedImage)
		}

		rb.selectedImage = img2

	*/
	rb.Panel = *NewPanel(ID, nil)

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

	rb.lastClickedTime = time.Now().UTC()
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
