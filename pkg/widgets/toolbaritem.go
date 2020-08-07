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
	p.ImageButton = *NewImageButton(ID,"./images/ptoolbaritem1.png","./images/nptoolbaritem1.png", p.eventHandler)
	return &p
}

