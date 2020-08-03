package widgets

import (
	"github.com/hajimehoshi/ebiten"
	log "github.com/sirupsen/logrus"
	_ "image/png"
)

type ImageButton struct {
	BaseButton

	// image for the button.
	nonPressedImage *ebiten.Image
	pressedImage    *ebiten.Image
}

func NewImageButton(ID string, pressedImageName string, nonPressedImageName string) *ImageButton {
	b := ImageButton{}

	img1, err := loadImage(pressedImageName)
	if err != nil {
		log.Fatalf("Unable to load image %s", pressedImageName)
	}
	b.pressedImage = img1
	img2, err := loadImage(nonPressedImageName)
	if err != nil {
		log.Fatalf("Unable to load image %s", nonPressedImageName)
	}
	b.nonPressedImage = img2

	width, height := b.nonPressedImage.Size()
	b.BaseButton = *NewBaseButton(ID, width, height, b.HandleEvent)
	//b.eventHandler = b.HandleEvent
	return &b
}

func (b *ImageButton) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.X, b.Y)

	// if state changed since last draw, recreate colour etc.

	if b.pressed {
		_ = screen.DrawImage(b.pressedImage, op)
	} else {
		_ = screen.DrawImage(b.nonPressedImage, op)
	}

	return nil
}
