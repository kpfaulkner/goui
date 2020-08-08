package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg"
	"github.com/kpfaulkner/goui/pkg/events"
	"github.com/kpfaulkner/goui/pkg/widgets"
	log "github.com/sirupsen/logrus"
	"image/color"
	_ "net/http/pprof"
)

type MyApp struct {
	mytext string
}

func (m *MyApp) ButtonAction1(event events.IEvent) error {
	log.Debugf("My button1 action 1!!!")
	return nil
}

func (m *MyApp) ButtonAction2(event events.IEvent) error {
	log.Debugf("My button1 action 2!!!")
	return nil
}


func (m *MyApp) ToolBarItem1(event events.IEvent) error {
	log.Debugf("toolbar item1")
	return nil
}

func (m *MyApp) ToolBarItem2(event events.IEvent) error {
	log.Debugf("toolbar item2")
	return nil
}

func (m *MyApp) CheckboxChanged(event events.IEvent) error {
	log.Debugf("checkbox changed!!!")
	cbe := event.(events.CheckBoxEvent)
	log.Debugf("checkbox is %v", cbe.Checked)
	return nil
}


func main() {

	log.SetLevel(log.DebugLevel)


	a := MyApp{}
	app := pkg.NewWindow(800, 600, "simple app 1", false, false)

	tb := widgets.NewToolBar("toolbar1", &color.RGBA{0,0,0,0xff})
	tbi1 := widgets.NewToolbarItem("tbi1", a.ToolBarItem1)
	tbi2 := widgets.NewToolbarItem("tbi2", a.ToolBarItem2)
	tb.AddToolBarItem(tbi1)
	tb.AddToolBarItem(tbi2)
	tb.SetSize(800,30)  // should calculate this!

	vpanel := widgets.NewVPanel("main vpanel", &color.RGBA{0,0,0,0xff})
	vpanel.AddWidget(tb)
	app.AddPanel(vpanel)

	hPanel := widgets.NewHPanel("hpanel1",&color.RGBA{0, 100, 0, 255})

	button1 := widgets.NewTextButton("text button 1", "my button1", 100, 40, nil, nil, nil, a.ButtonAction1)
	hPanel.AddWidget(button1)
	button2 := widgets.NewTextButton("text button 2", "my button2", 100, 40, nil, nil, nil, a.ButtonAction2)
	hPanel.AddWidget(button2)

	//spacer := widgets.NewEmptySpace("empty", 100,10)
	//hPanel.AddWidget(spacer)
	cb1 := widgets.NewCheckBox("my checkbox1", "my lovely checkbox", a.CheckboxChanged)
	hPanel.AddWidget(cb1)
	vpanel.AddWidget(hPanel)

	ebiten.SetRunnableInBackground(true)
	ebiten.SetWindowResizable(true)
	app.MainLoop()

}
