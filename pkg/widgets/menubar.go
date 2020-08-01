package widgets

import (
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
	"image/color"
)

type MenuItem struct {

	// Text on button.
	Name string

	// what to do if clicked.
	Handler func(event events.IEvent) error
}

// MenuDescription... for simple menus.
// Simple menu meaning a header... and items underneith it.
// No submenus (for now).
// These are purely used to describe what the menu system will look like.
// These are then converted to Panels, buttons etc.
type MenuDescription struct {
	MenuHeader string

	MenuItems []MenuItem
}

// MenuBar is just a custom panel....  (just an idea for now)
type MenuBar struct {
	Panel

	// panels of the menus that appear.
	// key is menu ID
	menuPanels map[string]Panel
}

// NewMenuBar creates the menubar panel. *has* to be located at 0,0 and full length of the window.
func NewMenuBar(ID string, x float64, y float64, width int, height int, colour *color.RGBA) *MenuBar {
	mb := MenuBar{}
	mb.Panel = *NewPanel(ID, x, y, width, height, colour)
	mb.menuPanels = make(map[string]Panel)

	return &mb
}

// AddMenuHeading adds the header for the menu eg. File, Edit, etc etc
// Tt does NOT add the contents/panel of when we click on that menu heading
func (mb *MenuBar) AddMenuHeading(menuName string) error {
	menuButton := *NewTextButton(menuName, menuName, 0, 0, 50, 20, nil, nil, nil)
	err := mb.AddWidget(&menuButton)

	// add panel for menu panel
	return err
}

func (mb *MenuBar) AddMenu(menuDesc MenuDescription) error {
	menuButton := *NewTextButton(menuDesc.MenuHeader, menuDesc.MenuHeader, 0, 0, 50, 20, nil, nil, nil)
	err := mb.AddWidget(&menuButton)

	// check number of menu entries
	menuEntries := len(menuDesc.MenuItems)

	// assume 30 pixele height for each menu option.
	// will obviously need to make configurable later.
	menuPanel := *NewPanel(menuDesc.MenuHeader, 0, 0, 40, menuEntries*30, nil)
	menuPanel.Disabled = true // disabled.... dont display it!

	for i, menuItem := range menuDesc.MenuItems {
		tb := *NewTextButton(menuItem.Name, menuItem.Name, 0, float64(i*30), 50, 30, nil, nil, nil)
		menuPanel.AddWidget(&tb)
	}

	mb.menuPanels[menuDesc.MenuHeader] = menuPanel

	// add panel for menu panel
	return err
}

func (mb *MenuBar) HandleEvent(event events.IEvent) (bool, error) {

	eventType := event.EventType()
	switch eventType {
	case events.EventTypeButtonDown:
		{
			mb.HandleMouseEvent(event)
		}
	case events.EventTypeButtonUp:
		{
			//	mb.HandleMouseEvent(event, local)
		}

	case events.EventTypeKeyboard:
		{
			mb.HandleKeyboardEvent(event)
		}
	}

	return true,nil
}

func (mb *MenuBar) HandleMouseEvent(event events.IEvent) error {
	inPanel, _ := mb.BaseWidget.CheckMouseEventCoords(event)

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
