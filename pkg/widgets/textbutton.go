package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/kpfaulkner/goui/pkg/common"
	"github.com/kpfaulkner/goui/pkg/events"
	"golang.org/x/image/font"
	"image"
	"image/color"
)

var (
	defaultNonPressedButtonColour color.RGBA
	defaultPressedButtonColour    color.RGBA
	defaultBorderColour    color.RGBA
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
	textIndent  int
}

func init() {
	defaultNonPressedButtonColour = color.RGBA{0x8A, 0x85, 0x81, 0xff}
	defaultPressedButtonColour = color.RGBA{0x78, 0x75, 0x72, 0xff}
	defaultBorderColour = color.RGBA{0,0,0xff, 0xff}

}

// NewTextButton. Button will have specified dimensions of width and height, useImageForSize is true.
// Then take the size from the contained text
func NewTextButton(ID string, text string, useImageforSize bool, width int, height int,
	nonPressedBackgroundColour *color.RGBA,
	pressedBackgroundColour *color.RGBA, fontInfo *common.Font, handler func(event events.IEvent) error) *TextButton {

	b := TextButton{}
	b.buttonText = text

	// font to use.
	if fontInfo != nil {
		b.fontInfo = *fontInfo
	} else {
		b.fontInfo = defaultFontInfo
	}

	widthToUse := 0
	heightToUse := 0
	if useImageforSize {
		bounds, _ := font.BoundString(b.fontInfo.UIFont, b.buttonText)
		widthToUse = (bounds.Max.X - bounds.Min.X).Ceil() + 5
		heightToUse = (bounds.Max.Y - bounds.Min.Y).Ceil() + 5
		b.textIndent = 2
	} else {
		widthToUse = width
		heightToUse = height

		bounds, _ := font.BoundString(b.fontInfo.UIFont, b.buttonText)
		textWidth := (bounds.Max.X - bounds.Min.X).Ceil()
		textHeight := (bounds.Max.Y - bounds.Min.Y).Ceil()
		b.textIndent = (widthToUse - textWidth) /2
		b.vertPos = (height - (height-int(textHeight))/2) - 2
	}

	b.BaseButton = *NewBaseButton(ID, widthToUse, heightToUse, handler)

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

	b.generateButtonImage(b.pressedBackgroundColour, defaultBorderColour)

	// vert pos is where does text go within button. Assuming we want it centred (for now)
	// Need to just find something visually appealing.
	b.vertPos = (b.Height - (b.Height-int(b.fontInfo.SizeInPixels))/2) - 2

	// just go off and listen for all events.
	//go b.ListenToIncomingEvents()
	return &b
}

func (b *TextButton) generateButtonImage(colour color.RGBA, border color.RGBA) error {
	//log.Infof("TextButton generateButtonImage")
	emptyImage, _ := ebiten.NewImage(b.Width, b.Height, ebiten.FilterDefault)
	_ = emptyImage.Fill(colour)

	ebitenutil.DrawLine(emptyImage, 0, 0, float64(b.Width),0, border)
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
			b.generateButtonImage(b.pressedBackgroundColour, defaultBorderColour)
		} else {
			b.generateButtonImage(b.nonPressedBackgroundColour, defaultBorderColour)
		}


		text.Draw(b.rectImage, b.buttonText, b.fontInfo.UIFont, b.textIndent, b.vertPos, color.Black)
		b.stateChangedSinceLastDraw = false
	}

	_ = screen.DrawImage(b.rectImage, op)
	return nil
}
