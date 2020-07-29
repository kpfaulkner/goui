package main

import (
	"github.com/kpfaulkner/goui/pkg"
	"github.com/kpfaulkner/goui/pkg/events"
	"github.com/kpfaulkner/goui/pkg/widgets"
	log "github.com/sirupsen/logrus"
	"image/color"
)

type MyApp struct {

}


func (m MyApp) ButtonAction1( event events.IEvent) error {
  log.Debugf("My button1 action!!!")
  return nil
}

func (m MyApp) ButtonAction2( event events.IEvent) error {
	log.Debugf("My button2 action!!!")
	return nil
}

func main() {
	log.SetLevel(log.DebugLevel)

	a := MyApp{}

	app := pkg.NewWindow(600, 600, "my title")
	//button := widgets.NewImageButton("c:/temp/test.png",0,0,100,100)

	panel := widgets.NewPanel(100, 0, 200, 200)
	button := widgets.NewTextButton("my button1", 0, 0, 100, 100, color.RGBA{0,0,0xff,0xff}, nil)
	button.RegisterEventHandler(events.EventTypeButtonDown, a.ButtonAction1)

	button2 := widgets.NewTextButton("my button2", 100, 0, 100, 100, color.RGBA{0,0xff,0,0xff}, nil)
	button2.RegisterEventHandler(events.EventTypeButtonDown, a.ButtonAction2)

	panel.AddButton(&button)
	panel.AddButton(&button2)

	app.AddPanel(panel)
	app.MainLoop()

}
