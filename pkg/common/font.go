package common

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"image/color"
)

type Font struct {
	Size   float64
	Colour color.RGBA
	UIFont font.Face
}

func LoadFont(name string, size float64, colour color.RGBA) Font {
	f := Font{}
	f.Size = size
	f.Colour = colour

	var tt *truetype.Font
	var err error

	if name == "" {
		tt, err = truetype.Parse(fonts.MPlus1pRegular_ttf)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Cannot select own font yet. Please leave empty.")
	}

	f.UIFont = truetype.NewFace(tt, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	return f
}
