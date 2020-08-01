package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/kpfaulkner/goui/pkg/common"
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"image"
	"image/color"
)

var (
	defaultNonPressedButtonColour color.RGBA
	defaultPressedButtonColour    color.RGBA
)

// TextButton is a button that just has a background colour, text and text colour.
type TextButton struct {
	BaseButton
	buttonText                 string
	pressedBackgroundColour    color.RGBA
	nonPressedBackgroundColour color.RGBA
	fontInfo                   common.Font
	uiFont                     font.Face
	Rect                       image.Rectangle
	vertPos                    int
}

func init() {
	defaultNonPressedButtonColour = color.RGBA{0x8A, 0x85, 0x81, 0xff}
	defaultPressedButtonColour = color.RGBA{0x78, 0x75, 0x72, 0xff}
}

func NewTextButton(ID string, text string, x float64, y float64, width int, height int,
	nonPressedBackgroundColour *color.RGBA,
	pressedBackgroundColour *color.RGBA, fontInfo *common.Font) *TextButton {

	b := TextButton{}
	b.BaseButton = *NewBaseButton(ID, x, y, width, height)

	if pressedBackgroundColour != nil {
		b.pressedBackgroundColour = *pressedBackgroundColour
	} else {
		b.pressedBackgroundColour = defaultPressedButtonColour
	}

	if nonPressedBackgroundColour != nil {
		b.nonPressedBackgroundColour = *nonPressedBackgroundColour
	} else {
		b.nonPressedBackgroundColour = defaultNonPressedButtonColour
	}

	b.buttonText = text

	if fontInfo != nil {
		b.fontInfo = *fontInfo
	} else {
		b.fontInfo = defaultFontInfo
	}

	b.generateButtonImage(b.pressedBackgroundColour, b.nonPressedBackgroundColour)

	// vert pos is where does text go within button. Assuming we want it centred (for now)
	// Need to just find something visually appealing.
	b.vertPos = (height - (height-int(b.fontInfo.SizeInPixels))/2) - 2

	b.eventHandler = b.HandleEvent

	// just go off and listen for all events.
	go b.ListenToIncomingEvents()
	return &b
}

func (b *TextButton) HandleEvent(event events.IEvent) (bool, error) {

	log.Debugf("TextButton::HandleEvent %s", b.ID)
	eventType := event.EventType()
	switch eventType {
	case events.EventTypeButtonDown:
		{

			mouseEvent := event.(events.MouseEvent)

			// check click is in button boundary.
			if b.ContainsCoords(mouseEvent.X, mouseEvent.Y) {
				log.Debugf("TextButton BUTTON DOWN!!!")
				b.hasFocus = true
				// if already pressed, then skip it.. .otherwise lots of repeats.
				//if !b.pressed {
				if true {
					b.pressed = true
					b.stateChangedSinceLastDraw = true
				}
			}
		}
	case events.EventTypeButtonUp:
		{

			mouseEvent := event.(events.MouseEvent)

			// check click is in button boundary.
			if b.ContainsCoords(mouseEvent.X, mouseEvent.Y) {
				log.Debugf("TextButton BUTTON UP!!!")
				b.hasFocus = true
				//if b.pressed {
				if true {
					// do generic button stuff here.
					b.pressed = false
					b.stateChangedSinceLastDraw = true
				}
			}
		}
	}
	return false, nil
}


func (b *TextButton) generateButtonImage(colour color.RGBA, border color.RGBA) error {
	//log.Infof("TextButton generateButtonImage")
	emptyImage, _ := ebiten.NewImage(b.Width, b.Height, ebiten.FilterDefault)
	_ = emptyImage.Fill(colour)

	ebitenutil.DrawLine(emptyImage, 0, 0, 0, float64(b.Width), border)
	ebitenutil.DrawLine(emptyImage, 0, float64(b.Width), float64(b.Width), float64(b.Height), border)
	ebitenutil.DrawLine(emptyImage, float64(b.Width), float64(b.Height), 0, float64(b.Height), border)
	ebitenutil.DrawLine(emptyImage, 0, float64(b.Height), 0, 0, border)

	b.rectImage = emptyImage
	return nil
}

func (b *TextButton) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.X, b.Y)

	// if state changed since last draw, recreate colour etc.
	//if b.stateChangedSinceLastDraw {
	if true {
		if b.pressed {
			log.Debugf("XXXXXXXXXX changing to pressed colour")
			b.generateButtonImage(b.pressedBackgroundColour, b.nonPressedBackgroundColour)
		} else {
			log.Debugf("YYYYYYYYY changing to nonpressed colour")
			b.generateButtonImage(b.nonPressedBackgroundColour, b.pressedBackgroundColour)
		}
		text.Draw(b.rectImage, b.buttonText, b.fontInfo.UIFont, 00, b.vertPos, color.Black)
		b.stateChangedSinceLastDraw = false
	}

	_ = screen.DrawImage(b.rectImage, op)

	return nil
}
