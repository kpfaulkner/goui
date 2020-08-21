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

	window pkg.Window
}

func NewMyApp() *MyApp {
	a := MyApp{}
	a.window = pkg.NewWindow(800, 600, "simple app 1", false, false)
	return &a
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

func (m *MyApp) SetupUI() error {

	// modified so its exactly fitting a 800x600 window :)
	vpanel1 := widgets.NewVPanelWithSize("fixedvpanel1", 800,600, &color.RGBA{0, 100, 0, 0xff})

	tb := widgets.NewToolBar("toolbar1", &color.RGBA{0, 0, 0, 0xff})
	tbi1 := widgets.NewToolbarItem("tbi1", m.ToolBarItem1)
	tbi2 := widgets.NewToolbarItem("tbi2", m.ToolBarItem2)
	tb.AddToolBarItem(tbi1)
	tb.AddToolBarItem(tbi2)
	tb.SetSize(800, 30) // should calculate this!
	vpanel1.AddWidget(tb)

	vpanel2 := widgets.NewVPanelWithSize("fixedvpanel2", 800,540, &color.RGBA{0, 0, 100, 0xff})
	vpanel1.AddWidget(vpanel2)

	ti := widgets.NewTextInput("textinput",800,30,nil,nil,nil)
	vpanel1.AddWidget(ti)

	button1 := widgets.NewTextButton("tb1","my button", false, 100,30,nil,nil,nil,nil )
  vpanel2.AddWidget(button1)
	m.window.AddPanel(vpanel1)

	return nil
}

func (m *MyApp) Run() error {

	m.SetupUI()

	ebiten.SetRunnableInBackground(true)
	ebiten.SetWindowResizable(true)
	m.window.MainLoop()

	return nil
}

func main() {

	log.SetLevel(log.DebugLevel)

	app := NewMyApp()
	app.Run()

}
