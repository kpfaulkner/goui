package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg"
	"github.com/kpfaulkner/goui/pkg/events"
	"github.com/kpfaulkner/goui/pkg/widgets"
	log "github.com/sirupsen/logrus"
	"image/color"
)

type MyApp struct {
	window pkg.Window
}

func NewMyApp() *MyApp {
	a := MyApp{}
	a.window = pkg.NewWindow(800, 600, "test app", false, false)
	return &a
}

func (m *MyApp) ButtonAction1(event events.IEvent) error {
	log.Debugf("My button1 action 1!!!")
	return nil
}

func (m *MyApp) SetupUI() error {

	vPanel := widgets.NewVPanel("main vpanel", &color.RGBA{0, 0, 0, 0xff})
	m.window.AddPanel(vPanel)

	hPanel := widgets.NewHPanel("hpanel1", &color.RGBA{0, 100, 0, 255})

	button1 := widgets.NewTextButton("text button 1", "my button1", true, 0, 0, nil, nil, nil, m.ButtonAction1)
	hPanel.AddWidget(button1)

	cb1 := widgets.NewCheckBox("my checkbox1", "check me please", "", "", nil)
	hPanel.AddWidget(cb1)

	vPanel.AddWidget(hPanel)

	imageButton := widgets.NewImageButton("ib1", "images/pressedbutton.png", "images/nonpressedbutton.png", nil)
	vPanel.AddWidget(imageButton)

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
