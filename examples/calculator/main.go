package main

import (
	"github.com/kpfaulkner/goui/pkg"
	"github.com/kpfaulkner/goui/pkg/common"
	"github.com/kpfaulkner/goui/pkg/events"
	"github.com/kpfaulkner/goui/pkg/widgets"
	log "github.com/sirupsen/logrus"
	"image/color"
	_ "net/http/pprof"
	"strconv"
)

const (

	// operations.
	Add int = iota
	Subtract
	Multiply
	Divide
	Equals
	None
)

type Calculator struct {
	currentInput string
	currentValue float64
	operation    int
	window       pkg.Window
	inputWidget  *widgets.TextInput
	fontInfo     common.Font
}

func NewCalculator() *Calculator {
	a := Calculator{}

	a.window = pkg.NewWindow(400, 400, "calculator", false, false)
	a.window.AddKeyboardHandler(a.AppKeyHandler)
	a.currentValue = 0
	a.operation = None
	a.fontInfo = common.LoadFont("", 32, color.RGBA{0xff, 0xff, 0xff, 0xff})
	return &a
}

func (m *Calculator) AppKeyHandler(event events.KeyboardEvent) error {
	log.Debugf("app key handler %s", string(event.Character))
	return nil
}

func (m *Calculator) ClearButton(event events.IEvent) error {
	log.Debugf("ClearButton : %s!!!", event.WidgetID())
	m.currentValue = 0
	m.currentInput = ""
	m.operation = None

	// update UI
	te := events.NewSetTextEvent(m.currentInput)
	m.inputWidget.HandleEvent(te)

	return nil
}

func (m *Calculator) EqualsButton(event events.IEvent) error {
	log.Debugf("EqualsButton : %s!!!", event.WidgetID())

	m.doCalculation()
	return nil
}

func (m *Calculator) applyOperation(num1 float64, num2 float64) float64 {

	switch m.operation {
	case Add:
		{
			return num1 + num2
		}
	case Subtract:
		{
			return num1 - num2
		}
	case Divide:
		{
			return num1 / num2
		}
	case Multiply:
		{
			return num1 * num2
		}
	case Equals:
		{
			return m.applyOperation(num1, num2)
		}
	}

	return 0
}

func (m *Calculator) doCalculation() {
	num, _ := strconv.ParseFloat(m.currentInput, 64)
	result := m.applyOperation(m.currentValue, num)
	m.currentValue = result
	m.currentInput = strconv.FormatFloat(m.currentValue, 'f', -1, 64)

	// update UI
	te := events.NewSetTextEvent(m.currentInput)
	m.inputWidget.HandleEvent(te)

}

func (m *Calculator) NumberButtonAction(event events.IEvent) error {
	log.Debugf("NumberButtonAction : %s!!!", event.WidgetID())

	m.currentInput += event.WidgetID()
	te := events.NewSetTextEvent(m.currentInput)
	m.inputWidget.HandleEvent(te)

	return nil
}

func (m *Calculator) AddButton(event events.IEvent) error {
	log.Debugf("AddButton")
	m.operation = Add
	m.parseNumber(m.currentInput)
	return nil
}

func (m *Calculator) parseNumber(n string) error {
	num, _ := strconv.ParseFloat(m.currentInput, 64)
	m.currentValue = num
	m.currentInput = ""
	return nil
}

func (m *Calculator) SubtractButton(event events.IEvent) error {
	log.Debugf("SubtractButton")
	m.operation = Subtract
	m.parseNumber(m.currentInput)
	return nil
}

func (m *Calculator) MultiplyButton(event events.IEvent) error {
	log.Debugf("MultiplyButton")
	m.operation = Multiply
	m.parseNumber(m.currentInput)
	return nil
}

func (m *Calculator) DivideButton(event events.IEvent) error {
	log.Debugf("DivideButton")
	m.operation = Divide
	m.parseNumber(m.currentInput)
	return nil
}

func (m *Calculator) createNumberButton(text string, handler func(event events.IEvent) error) widgets.TextButton {
	tb := widgets.NewTextButton(text, text, false, 80, 80, nil, nil, &m.fontInfo, handler)
	return *tb
}

// SetupUI.
func (m *Calculator) SetupUI() *widgets.VPanel {

	vPanel := widgets.NewVPanel("rowsofbuttons", &color.RGBA{0, 0, 0, 0xff})

	textEntry := widgets.NewTextInput("textinput", 400, 50, &color.RGBA{0, 0, 0, 0xff}, &m.fontInfo, nil)
	m.inputWidget = textEntry
	vPanel.AddWidget(textEntry)

	hPanel1 := widgets.NewHPanel("first row", nil)
	sevenButton := m.createNumberButton("7", m.NumberButtonAction)
	hPanel1.AddWidget(&sevenButton)
	eightButton := m.createNumberButton("8", m.NumberButtonAction)
	hPanel1.AddWidget(&eightButton)
	nineButton := m.createNumberButton("9", m.NumberButtonAction)
	hPanel1.AddWidget(&nineButton)
	addButton := m.createNumberButton("+", m.AddButton)
	hPanel1.AddWidget(&addButton)
	vPanel.AddWidget(hPanel1)

	hPanel2 := widgets.NewHPanel("second row", nil)
	fourButton := m.createNumberButton("4", m.NumberButtonAction)
	hPanel2.AddWidget(&fourButton)
	fiveButton := m.createNumberButton("5", m.NumberButtonAction)
	hPanel2.AddWidget(&fiveButton)
	sixButton := m.createNumberButton("6", m.NumberButtonAction)
	hPanel2.AddWidget(&sixButton)
	subButton := m.createNumberButton("-", m.SubtractButton)
	hPanel2.AddWidget(&subButton)
	vPanel.AddWidget(hPanel2)

	hPanel3 := widgets.NewHPanel("third row", nil)
	oneButton := m.createNumberButton("1", m.NumberButtonAction)
	hPanel3.AddWidget(&oneButton)
	twoButton := m.createNumberButton("2", m.NumberButtonAction)
	hPanel3.AddWidget(&twoButton)
	threeButton := m.createNumberButton("3", m.NumberButtonAction)
	hPanel3.AddWidget(&threeButton)
	multiplyButton := m.createNumberButton("*", m.MultiplyButton)
	hPanel3.AddWidget(&multiplyButton)
	vPanel.AddWidget(hPanel3)

	hPanel4 := widgets.NewHPanel("fourth row", nil)
	zeroButton := m.createNumberButton("0", m.NumberButtonAction)
	hPanel4.AddWidget(&zeroButton)
	dotButton := m.createNumberButton(".", m.NumberButtonAction)
	hPanel4.AddWidget(&dotButton)
	equalsButton := m.createNumberButton("=", m.EqualsButton)
	hPanel4.AddWidget(&equalsButton)
	divideButton := m.createNumberButton("/", m.DivideButton)
	hPanel4.AddWidget(&divideButton)
	clearButton := m.createNumberButton("C", m.ClearButton)
	hPanel4.AddWidget(&clearButton)
	vPanel.AddWidget(hPanel4)

	m.window.AddPanel(vPanel)

	return vPanel

}

func (m *Calculator) Run() error {

	_ = m.SetupUI()
	//ebiten.SetWindowResizable(true)
	m.window.MainLoop()

	return nil
}

func main() {

	log.SetLevel(log.DebugLevel)

	app := NewCalculator()
	app.Run()

}
