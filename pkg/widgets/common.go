package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	"os"
)

// loadImage, assuming its a png
func loadImage(imageName string) (*ebiten.Image, error) {
	f, err := os.Open(imageName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	ei, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	return ei, nil
}
