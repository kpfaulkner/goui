package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg"
	"github.com/kpfaulkner/goui/pkg/events"
	"github.com/kpfaulkner/goui/pkg/widgets"
	log "github.com/sirupsen/logrus"
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
	inputWidget *widgets.TextInput
}

func NewCalculator() *Calculator {
	a := Calculator{}

	a.window = pkg.NewWindow(400, 400, "calculator", false, false)
	a.window.AddKeyboardHandler(a.AppKeyHandler)
	a.currentValue = 0
	a.operation = None

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
			return num1 * num1
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


	m.operation = Add
	num, _ := strconv.ParseFloat(m.currentInput, 64)
	m.currentValue = num
	m.currentInput = ""
	log.Debugf("AddButton")

	return nil
}

func (m *Calculator) SubtractButton(event events.IEvent) error {
	m.operation = Subtract
	num, _ := strconv.ParseFloat(m.currentInput, 64)
	m.currentValue = num
	m.currentInput = ""
	log.Debugf("SubtractButton")
	return nil
}

func (m *Calculator) MultiplyButton(event events.IEvent) error {
	m.operation = Multiply
	num, _ := strconv.ParseFloat(m.currentInput, 64)
	m.currentValue = num
	m.currentInput = ""
	log.Debugf("MultiplyButton")
	return nil
}

func (m *Calculator) DivideButton(event events.IEvent) error {
	m.operation = Divide
	num, _ := strconv.ParseFloat(m.currentInput, 64)
	m.currentValue = num
	m.currentInput = ""
	log.Debugf("DivideButton")
	return nil
}

func createNumberButton(text string, handler func(event events.IEvent) error) widgets.TextButton {
	tb := widgets.NewTextButton(text, text, false, 20, 20, nil, nil, nil, handler)
	return *tb
}

// SetupOutPanel.
// Then returns VPanel that the rest of the interface can go into.
func (m *Calculator) SetupOuterUI() *widgets.VPanel {

	vPanel := widgets.NewVPanel("rowsofbuttons", nil)

	textEntry := widgets.NewTextInput("textinput", 100, 20, nil, nil, nil)
	m.inputWidget = textEntry
	vPanel.AddWidget(textEntry)

	hPanel1 := widgets.NewHPanel("first row", nil)
	sevenButton := createNumberButton("7", m.NumberButtonAction)
	hPanel1.AddWidget(&sevenButton)
	eightButton := createNumberButton("8", m.NumberButtonAction)
	hPanel1.AddWidget(&eightButton)
	nineButton := createNumberButton("9", m.NumberButtonAction)
	hPanel1.AddWidget(&nineButton)
	addButton := createNumberButton("+", m.AddButton)
	hPanel1.AddWidget(&addButton)
	vPanel.AddWidget(hPanel1)

	hPanel2 := widgets.NewHPanel("second row", nil)
	fourButton := createNumberButton("4", m.NumberButtonAction)
	hPanel2.AddWidget(&fourButton)
	fiveButton := createNumberButton("5", m.NumberButtonAction)
	hPanel2.AddWidget(&fiveButton)
	sixButton := createNumberButton("6", m.NumberButtonAction)
	hPanel2.AddWidget(&sixButton)
	subButton := createNumberButton("-", m.SubtractButton)
	hPanel2.AddWidget(&subButton)
	vPanel.AddWidget(hPanel2)

	hPanel3 := widgets.NewHPanel("third row", nil)
	oneButton := createNumberButton("1", m.NumberButtonAction)
	hPanel3.AddWidget(&oneButton)
	twoButton := createNumberButton("2", m.NumberButtonAction)
	hPanel3.AddWidget(&twoButton)
	threeButton := createNumberButton("3", m.NumberButtonAction)
	hPanel3.AddWidget(&threeButton)
	multiplyButton := createNumberButton("*", m.MultiplyButton)
	hPanel3.AddWidget(&multiplyButton)
	vPanel.AddWidget(hPanel3)

	hPanel4 := widgets.NewHPanel("fourth row", nil)
	zeroButton := createNumberButton("0", m.NumberButtonAction)
	hPanel4.AddWidget(&zeroButton)
	dotButton := createNumberButton(".", m.NumberButtonAction)
	hPanel4.AddWidget(&dotButton)
	equalsButton := createNumberButton("=", m.EqualsButton)
	hPanel4.AddWidget(&equalsButton)
	divideButton := createNumberButton("/", m.DivideButton)
	hPanel4.AddWidget(&divideButton)
	vPanel.AddWidget(hPanel4)

	clearButton := createNumberButton("C", m.ClearButton)
	hPanel4.AddWidget(&clearButton)

	m.window.AddPanel(vPanel)

	return vPanel

}

func (m *Calculator) Run() error {

	_ = m.SetupOuterUI()

	ebiten.SetRunnableInBackground(true)
	ebiten.SetWindowResizable(true)
	m.window.MainLoop()

	return nil
}

func main() {

	log.SetLevel(log.DebugLevel)

	app := NewCalculator()
	app.Run()

}
