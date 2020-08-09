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
	a.window.AddKeyboardHandler(a.AppKeyHandler)
	a.window.AddMouseHandler(a.AppMouseHandler)
	return &a
}

func (m *MyApp) AppKeyHandler(event events.KeyboardEvent) error {
	log.Debugf("app key handler %s", string(event.Character))
	return nil
}

func (m *MyApp) AppMouseHandler(event events.MouseEvent) error {

	switch event.EventType() {
	case events.EventTypeButtonDown:
		{
			log.Debugf("app mouse handler button down")
		}

	case events.EventTypeButtonUp:
		{
			log.Debugf("app mouse handler button up")
		}

	case events.EventTypeMouseMove:
		{
			log.Debugf("app mouse handler, x: %0.f , y: %0.f", event.X, event.Y)
		}
	}

	return nil
}

func (m *MyApp) ButtonAction1(event events.IEvent) error {
	log.Debugf("My button1 action 1 : %s!!!", event.WidgetID())
	return nil
}

func (m *MyApp) ButtonAction2(event events.IEvent) error {
	log.Debugf("My button1 action 2 : %s!!!", event.WidgetID())
	return nil
}

func (m *MyApp) ButtonActionGeneric(event events.IEvent) error {
	log.Debugf("Generic button for widget : %s!!!", event.WidgetID())
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

func (m *MyApp) SetupRestOfUI(vpanel *widgets.VPanel) error {

	vpanel2 := widgets.NewVPanel("vpanel2", nil)
	vpanel.AddWidget(vpanel2)

	hPanel := widgets.NewHPanel("hpanel1", &color.RGBA{0, 100, 0, 255})

	button1 := widgets.NewTextButton("text button 1", "my button1", true, 0, 0, nil, nil, nil, m.ButtonActionGeneric)
	hPanel.AddWidget(button1)
	button3 := widgets.NewTextButton("text button 3", "my button3", true, 0, 0, nil, nil, nil, m.ButtonActionGeneric)
	hPanel.AddWidget(button3)
	button4 := widgets.NewTextButton("text button 4", "my button4", true, 0, 0, nil, nil, nil, m.ButtonActionGeneric)
	hPanel.AddWidget(button4)
	button5 := widgets.NewTextButton("text button 5", "my button5", true, 0, 0, nil, nil, nil, m.ButtonActionGeneric)
	hPanel.AddWidget(button5)

	vpanel2.AddWidget(hPanel)

	hPanel2 := widgets.NewHPanel("hpanel2", &color.RGBA{50, 50, 0, 255})

	button2 := widgets.NewTextButton("text button 2", "my button2", false, 100, 40, nil, nil, nil, m.ButtonAction2)
	hPanel2.AddWidget(button2)

	//spacer := widgets.NewEmptySpace("empty", 100,10)
	//hPanel.AddWidget(spacer)
	cb1 := widgets.NewCheckBox("my checkbox1", "my lovely checkbox", m.CheckboxChanged)
	hPanel2.AddWidget(cb1)
	vpanel2.AddWidget(hPanel2)

	return nil
}

// SetupOutPanel.
// Then returns VPanel that the rest of the interface can go into.
func (m *MyApp) SetupOuterUI() *widgets.VPanel {

	toolbar := widgets.NewToolBar("toolbar1", &color.RGBA{0, 0, 0, 0xff})
	tbi1 := widgets.NewToolbarItem("tbi1", m.ToolBarItem1)
	tbi2 := widgets.NewToolbarItem("tbi2", m.ToolBarItem2)
	tbi3 := widgets.NewToolbarItem("tbi3", m.ToolBarItem2)
	toolbar.AddToolBarItem(tbi1)
	toolbar.AddToolBarItem(tbi2)
	toolbar.AddToolBarItem(tbi3)
	toolbar.SetSize(800, 30) // should calculate this!

	vpanel := widgets.NewVPanel("main vpanel", &color.RGBA{0, 0, 0, 0xff})
	vpanel.AddWidget(toolbar)
	m.window.AddPanel(vpanel)

	return vpanel

}

func (m *MyApp) Run() error {

	vpanel := m.SetupOuterUI()
	m.SetupRestOfUI(vpanel)

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
