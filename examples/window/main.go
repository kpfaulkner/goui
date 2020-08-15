package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg"
	"github.com/kpfaulkner/goui/pkg/common"
	"github.com/kpfaulkner/goui/pkg/events"
	"github.com/kpfaulkner/goui/pkg/widgets"
	log "github.com/sirupsen/logrus"
	"image/color"
	"net/http"
	_ "net/http/pprof"
	"time"
)

type MyApp struct {
	mytext string
}

func (m *MyApp) ButtonAction1(event events.IEvent) error {
	log.Debugf("My button1 action!!!")
	return nil
}

func (m *MyApp) ButtonAction2(event events.IEvent) error {
	log.Debugf("My button2 action!!!")
	return nil
}

func (m *MyApp) ButtonAction3(event events.IEvent) error {
	log.Debugf("My button3 action!!!")
	return nil
}

func (m *MyApp) ToolBarItem1(event events.IEvent) error {
	log.Debugf("toolbar!!!")
	return nil
}

func (m *MyApp) CheckboxChanged(event events.IEvent) error {
	log.Debugf("checkbox changed!!!")

	cbe := event.(events.CheckBoxEvent)
	log.Debugf("checkbox is %v", cbe.Checked)
	return nil
}

func (m *MyApp) HandleTextInput(event events.IEvent) error {
	log.Debugf("text input changed.!!!")

	kbEvent := event.(events.KeyboardEvent)
	m.mytext += string(kbEvent.Character)
	return nil
}

func addPanel(panelName string, width int, height int, win *pkg.Window, buttonAction1 func(event events.IEvent) error, buttonAction2 func(event events.IEvent) error) error {
	panel := widgets.NewPanel(panelName, nil,nil)
	button := widgets.NewTextButton("button1", "my button1", false, 100, 100, nil, nil, nil, buttonAction1)
	panel.AddWidget(button)
	button2 := widgets.NewTextButton("button2", "my button2", true, 100, 100, nil, nil, nil, buttonAction2)
	panel.AddWidget(button2)

	win.AddPanel(panel)

	return nil
}

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	log.SetLevel(log.DebugLevel)

	a := MyApp{}
	app := pkg.NewWindow(600, 600, "my title", false, true)

	tb := widgets.NewToolBar("my toolbar", &color.RGBA{0, 0, 0, 0xff})
	tbi := widgets.NewToolbarItem("tbi", a.ToolBarItem1)

	tb.AddToolBarItem(tbi)
	tb.SetSize(600, 30)
	panel := widgets.NewVPanel("panel3", &color.RGBA{0, 0, 0, 0xff})

	panel.AddWidget(tb)

	app.AddPanel(panel)

	hPanel := widgets.NewHPanel("hpanel1", &color.RGBA{0, 100, 0, 255})

	button := widgets.NewTextButton("button1", "my button1", false, 100, 100, nil, nil, nil, a.ButtonAction1)
	hPanel.AddWidget(button)
	button2 := widgets.NewTextButton("button2", "my button2", false, 100, 100, nil, nil, nil, a.ButtonAction2)
	hPanel.AddWidget(button2)

	spacer := widgets.NewEmptySpace("empty", 100, 10)
	hPanel.AddWidget(spacer)
	cb1 := widgets.NewCheckBox("my checkbox1", "cb1", "", "", a.CheckboxChanged)
	hPanel.AddWidget(cb1)

	panel.AddWidget(hPanel)

	button3 := widgets.NewTextButton("etxt button 3", "my button3", false, 100, 100, nil, nil, nil, a.ButtonAction3)
	panel.AddWidget(button3)

	cb2 := widgets.NewCheckBox("my checkbox2", "checkbox 2", "", "", a.CheckboxChanged)
	panel.AddWidget(cb2)
	cb3 := widgets.NewCheckBox("my checkbox3", "checkbox 3", "", "", a.CheckboxChanged)
	panel.AddWidget(cb3)

	f := common.LoadFont("", 16, color.RGBA{0xff, 0xff, 0xff, 0xff})
	ti := widgets.NewTextInput("testinput1", 200, 20, &color.RGBA{0x55, 0x55, 0x55, 0xff}, &f, a.HandleTextInput)
	panel.AddWidget(ti)

	l := widgets.NewLabel("label", "my text", 200, 20, &color.RGBA{0x55, 0x55, 0x55, 0xff}, &f)
	panel.AddWidget(l)

	go func() {
		fullText := "the quick brown fox jumps over the lazy dogs"
		i := 0
		for {
			<-time.After(500 * time.Millisecond)
			if i+10 >= len(fullText) {
				i = 0
			}
			subText := fullText[i : i+10]
			te := events.NewSetTextEvent(subText)
			ti.HandleEvent(te)
			i++
		}
	}()

	ebiten.SetRunnableInBackground(true)
	ebiten.SetWindowResizable(true)
	app.MainLoop()

}
