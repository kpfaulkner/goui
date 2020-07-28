package widgets

import (
	"github.com/hajimehoshi/ebiten"
	log "github.com/sirupsen/logrus"
	"image"
	_ "image/png"
	"os"
)

type ImageButton struct {
	BaseButton

	// image for the button.
	image image.Image
}

func NewImageButton(imageName string, x float64, y float64, width int, height int) ImageButton {
	b := ImageButton{}
	b.BaseButton = NewBaseButton(x, y, width, height)
	img, err := loadImage(imageName)
	if err != nil {
		log.Fatalf("Unable to load image %s", imageName)
	}
	b.image = *img
	return b
}

// loadImage, assuming its a png
func loadImage(imageName string) (*image.Image, error) {
	f, err := os.Open(imageName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return &img, nil
}

func (b *ImageButton) Draw(screen *ebiten.Image) error {
	return nil
}
