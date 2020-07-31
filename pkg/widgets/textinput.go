package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/kpfaulkner/goui/pkg/common"
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"image/color"
	_ "image/png"
)

var defaultFontInfo common.Font

type TextInput struct {
	BaseWidget

	text             string
	backgroundColour color.RGBA
	fontInfo         common.Font
	uiFont           font.Face

	// just for cursor.
	counter int
}

func init() {
	defaultFontInfo = common.LoadFont("", 16, color.RGBA{0xff, 0xff, 0xff, 0xff})
}

func NewTextInput(ID string, x float64, y float64, width int, height int, backgroundColour *color.RGBA, fontInfo *common.Font) TextInput {
	t := TextInput{}
	t.BaseWidget = NewBaseWidget(ID, x, y, width, height)
	t.text = ""
	t.stateChangedSinceLastDraw = true
	t.counter = 0

	if backgroundColour != nil {
		t.backgroundColour = *backgroundColour
	} else {
		t.backgroundColour = color.RGBA{0, 0xff, 0, 0xff}
	}

	if fontInfo != nil {
		t.fontInfo = *fontInfo
	} else {
		t.fontInfo = defaultFontInfo
	}

	return t
}

func (t *TextInput) HandleEvent(event events.IEvent) error {

	eventType := event.EventType()
	switch eventType {
	case events.EventTypeButtonDown:
		{
			mouseEvent := event.(events.MouseEvent)

			// check click is in button boundary.
			if t.ContainsCoords(mouseEvent.X, mouseEvent.Y, true) {
				t.hasFocus = true
				t.stateChangedSinceLastDraw = true
				// then do application specific stuff!!
				if ev, ok := t.eventRegister[event.EventType()]; ok {
					ev(event)
				}
			} else {
				t.hasFocus = false
			}
		}
	case events.EventTypeKeyboard:
		{
			// check if has focus....  if so, can potentially add to string?
			if t.hasFocus {
				keyboardEvent := event.(events.KeyboardEvent)

				if keyboardEvent.Character != ebiten.KeyBackspace {
					t.text = t.text + string(keyboardEvent.Character)
				} else {
					// back space one.
					l := len(t.text)
					if l > 0 {
						t.text = t.text[0 : l-1]
					}
				}
				t.stateChangedSinceLastDraw = true
			}
		}
	}
	return nil
}

func (t *TextInput) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.X, t.Y)

	if t.stateChangedSinceLastDraw {
		log.Debugf("textinput text %s", t.text)
		// how often do we update this?
		emptyImage, _ := ebiten.NewImage(t.Width, t.Height, ebiten.FilterDefault)
		_ = emptyImage.Fill(t.backgroundColour)
		t.rectImage = emptyImage

		ebitenutil.DrawLine(t.rectImage, 0, 0, float64(t.Width), 0, color.Black)
		ebitenutil.DrawLine(t.rectImage, float64(t.Width), 0, float64(t.Width), float64(t.Height), color.Black)
		ebitenutil.DrawLine(t.rectImage, float64(t.Width), float64(t.Height), 0, float64(t.Height), color.Black)
		ebitenutil.DrawLine(t.rectImage, 0, float64(t.Height), 0, 0, color.Black)

		txt := t.text
		txt += "|"

		text.Draw(t.rectImage, txt, t.fontInfo.UIFont, 0, 15, color.Black)
		t.stateChangedSinceLastDraw = false
	}

	// if state changed since last draw, recreate colour etc.
	_ = screen.DrawImage(t.rectImage, op)

	return nil
}

func (t *TextInput) GetData() (interface{}, error) {
	return t.text, nil
}

