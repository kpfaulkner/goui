package widgets

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/kpfaulkner/goui/pkg/common"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
)

var (
	nonPressedButtonBorder color.RGBA
	pressedButtonBorder color.RGBA
)
// TextButton is a button that just has a background colour, text and text colour.
type TextButton struct {
	BaseButton
	buttonText       string
	backgroundColour color.RGBA
	fontInfo         *common.Font
	uiFont           font.Face
	Rect             image.Rectangle
}

func init() {
	nonPressedButtonBorder = color.RGBA{0xff,0,0,0xff}
	pressedButtonBorder = color.RGBA{0,0xff,0,0xff}
}

func NewTextButton(text string, x float64, y float64, width int, height int, backgroundColour color.RGBA, fontInfo *common.Font) TextButton {
	b := TextButton{}
	b.BaseButton = NewBaseButton(x, y, width, height)
	b.backgroundColour = backgroundColour
	b.buttonText = text
	b.fontInfo = fontInfo
	b.generateButtonImage(b.backgroundColour, nonPressedButtonBorder)


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

func (b *TextButton) generateButtonImage(colour color.RGBA, border color.RGBA) error {
	log.Infof("TextButton generateButtonImage")
	emptyImage, _ := ebiten.NewImage(b.Width, b.Height, ebiten.FilterDefault)
	_ = emptyImage.Fill(colour)

	ebitenutil.DrawLine(emptyImage,0,0,0,float64(b.Width),border)
	ebitenutil.DrawLine(emptyImage,0,float64(b.Width), float64(b.Width), float64(b.Height),border)
	ebitenutil.DrawLine(emptyImage,float64(b.Width), float64(b.Height),0,float64(b.Height),border)
	ebitenutil.DrawLine(emptyImage,0,float64(b.Height),0,0,border)

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
			b.generateButtonImage(b.backgroundColour, pressedButtonBorder)
		} else {
			log.Debugf("changing to nonpressed colour")
			b.generateButtonImage(b.backgroundColour, nonPressedButtonBorder)
		}
		b.stateChangedSinceLastDraw = false
	}

	text.Draw(b.rectImage, b.buttonText, b.uiFont, 20, 50, color.Black)
	_ = screen.DrawImage(b.rectImage, op)

	return nil
}
