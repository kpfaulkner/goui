package widgets

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/kpfaulkner/goui/pkg/common"
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
)

// TextButton is a button that just has a background colour, text and text colour.
type TextButton struct {
	BaseButton
	buttonText       string
	backgroundColour string
	fontInfo         *common.Font
	uiFont           font.Face
	Rect             image.Rectangle
}

func NewTextButton(text string, x float64, y float64, width int, height int, backgroundColour string, fontInfo *common.Font) TextButton {
	b := TextButton{}
	b.BaseButton = NewBaseButton(x, y, width, height)
	b.backgroundColour = backgroundColour
	b.buttonText = text
	b.fontInfo = fontInfo
	b.generateButtonImage()
	tt, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}
	b.uiFont = truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	return b
}

func (b *TextButton) generateButtonImage() error {
	log.Infof("TextButton generateButtonImage")
	emptyImage, _ := ebiten.NewImage(b.Width, b.Height, ebiten.FilterDefault)
	_ = emptyImage.Fill(color.White)
	b.rectImage = emptyImage
	return nil
}

func (b *TextButton) HandleEventXX(event events.IEvent) error {
	eventType := event.EventType()
	switch eventType {
	case events.EventTypeButtonDown:
		{
			b.HandleMouseEvent(event)

		}
	case events.EventTypeButtonUp:
		{
			b.HandleMouseEvent(event)
		}
	}

	return nil
}

func (b *TextButton) HandleMouseEvent(event events.IEvent) error {
	inButton, _ := b.BaseWidget.CheckMouseEventCoords(event)

	if inButton {
		mouseEvent := event.(events.MouseEvent)
		log.Debugf("textbutton handled mouse event at %f %f", mouseEvent.X, mouseEvent.Y)

		// registered event
		if ev,ok := b.eventRegister[event.EventType()]; ok {
			ev(event)
		}
	}
	// should propagate to children nodes?
	return nil
}

func (b *TextButton) Draw(screen *ebiten.Image) error {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.X, b.Y)

	text.Draw(b.rectImage, b.buttonText, b.uiFont, 20, 50, color.Black)
	_ = screen.DrawImage(b.rectImage, op)

	return nil
}
