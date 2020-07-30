package widgets

import (
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
	"image/color"
)

// MenuBar is just a custom panel....  (just an idea for now)
type MenuBar struct {
	Panel

}

// NewMenuBar creates the menubar panel. *has* to be located at 0,0 and full length of the window.
func NewMenuBar(ID string, x float64, y float64, width int, height int, colour *color.RGBA) MenuBar {
	mb := MenuBar{}
	mb.Panel = NewPanel(ID,x,y,width,height,colour)
	return mb
}

// AddMenuHeading adds the header for the menu eg. File, Edit, etc etc
// Tt does NOT add the contents/panel of when we click on that menu heading
func (mb *MenuBar) AddMenuHeading(menuName string) error {

	menuButton := NewTextButton(menuName,menuName,0,0,50,20,nil,nil,nil)
	err := mb.AddWidget(&menuButton)
	return err
}


func (mb *MenuBar) HandleEvent(event events.IEvent, local bool) error {

	eventType := event.EventType()
	switch eventType {
	case events.EventTypeButtonDown:
		{
			mb.HandleMouseEvent(event, local)
		}
	case events.EventTypeButtonUp:
		{
			mb.HandleMouseEvent(event, local)
		}

	case events.EventTypeKeyboard:
		{
			mb.HandleKeyboardEvent(event, local)
		}
	}

	return nil
}

func (mb *MenuBar) HandleMouseEvent(event events.IEvent, local bool) error {
	inPanel, _ := mb.BaseWidget.CheckMouseEventCoords(event, local)

	if inPanel {
		mouseEvent := event.(events.MouseEvent)
		log.Debugf("HandleMouseEvent panel %s :  %f %f", mb.ID, mouseEvent.X, mouseEvent.Y)
		localCoordMouseEvent := mb.GenerateLocalCoordMouseEvent(mouseEvent)

		for _, widget := range mb.widgets {
			widget.HandleEvent(localCoordMouseEvent)
		}
	}

	// should propagate to children nodes?
	return nil
}

