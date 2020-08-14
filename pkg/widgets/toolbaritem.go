package widgets

import (
	"github.com/kpfaulkner/goui/pkg/events"
)

// Toolbar stacks internal objects horizonally... left to right.
type ToolbarItem struct {
	ImageButton
}

func NewToolbarItem(ID string, handler func(event events.IEvent) error) *ToolbarItem {
	p := ToolbarItem{}
	p.eventHandler = handler
	//p.ImageButton = *NewImageButton(ID,"./images/square-button1-small.png","./images/square-button2-small.png", p.eventHandler)
	p.ImageButton = *NewImageButton(ID, "./images/but1.png", "./images/but2.png", p.eventHandler)
	return &p
}
