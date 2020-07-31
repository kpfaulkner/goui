package widgets

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/kpfaulkner/goui/pkg/common"
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
}

func init() {
	defaultNonPressedButtonColour = color.RGBA{0x8A, 0x85, 0x81, 0xff}
	defaultPressedButtonColour = color.RGBA{0x78, 0x75, 0x72, 0xff}
}

func NewTextButton(ID string, text string, x float64, y float64, width int, height int,
	nonPressedBackgroundColour *color.RGBA,
	pressedBackgroundColour *color.RGBA, fontInfo *common.Font) TextButton {

	b := TextButton{}
	b.BaseButton = NewBaseButton(ID, x, y, width, height)

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

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
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

func (b *TextButton) generateButtonImage(colour color.RGBA, border color.RGBA) error {
	log.Infof("TextButton generateButtonImage")
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
	if b.stateChangedSinceLastDraw {
		if b.pressed {
			log.Debugf("changing to pressed colour")
			b.generateButtonImage(b.pressedBackgroundColour, b.nonPressedBackgroundColour)
		} else {
			log.Debugf("changing to nonpressed colour")
			b.generateButtonImage(b.nonPressedBackgroundColour, b.pressedBackgroundColour)
		}
		text.Draw(b.rectImage, b.buttonText, b.fontInfo.UIFont, 00, 20, color.Black)
		b.stateChangedSinceLastDraw = false
	}

	_ = screen.DrawImage(b.rectImage, op)

	return nil
}
