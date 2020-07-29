package main

import (
	"github.com/kpfaulkner/goui/pkg"
	"github.com/kpfaulkner/goui/pkg/events"
	"github.com/kpfaulkner/goui/pkg/widgets"
	log "github.com/sirupsen/logrus"
)

func myButtonAction( event events.IEvent) error {
  log.Debugf("My button action!!!")
  return nil
}

func main() {
	log.SetLevel(log.DebugLevel)

	app := pkg.NewWindow(600, 600, "my title")
	//button := widgets.NewImageButton("c:/temp/test.png",0,0,100,100)

	panel := widgets.NewPanel(100, 0, 200, 200)
	button := widgets.NewTextButton("my button", 0, 0, 100, 100, "", nil)
	button.RegisterEventHandler(events.EventTypeButtonDown, myButtonAction)

	panel.AddButton(&button)

	app.AddPanel(panel)
	app.MainLoop()

}
