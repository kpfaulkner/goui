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

type Label struct {
	BaseWidget

	text             string
	backgroundColour color.RGBA
	fontInfo         common.Font
	uiFont           font.Face

	// vertical position for text
	vertPos int
}

func NewLabel(ID string, text string, width int, height int, backgroundColour *color.RGBA, fontInfo *common.Font) *Label {
	t := Label{}
	t.BaseWidget = *NewBaseWidget(ID, width, height, nil)
	t.text = text
	t.stateChangedSinceLastDraw = true

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

	// vert pos is where does text go within button. Assuming we want it centred (for now)
	// Need to just find something visually appealing.
	t.vertPos = (height - (height-int(t.fontInfo.SizeInPixels))/2) - 2

	return &t
}

func (t *Label) HandleEvent(event events.IEvent) error {

	return nil
}

func (t *Label) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.X, t.Y)

	if t.stateChangedSinceLastDraw {
		log.Debugf("label text %s", t.text)
		// how often do we update this?
		emptyImage, _ := ebiten.NewImage(t.Width, t.Height, ebiten.FilterDefault)
		_ = emptyImage.Fill(t.backgroundColour)
		t.rectImage = emptyImage

		ebitenutil.DrawLine(t.rectImage, 0, 0, float64(t.Width), 0, color.Black)
		ebitenutil.DrawLine(t.rectImage, float64(t.Width), 0, float64(t.Width), float64(t.Height), color.Black)
		ebitenutil.DrawLine(t.rectImage, float64(t.Width), float64(t.Height), 0, float64(t.Height), color.Black)
		ebitenutil.DrawLine(t.rectImage, 0, float64(t.Height), 0, 0, color.Black)

		text.Draw(t.rectImage, t.text, t.fontInfo.UIFont, 0, t.vertPos, t.fontInfo.Colour)
		t.stateChangedSinceLastDraw = false
	}

	// if state changed since last draw, recreate colour etc.
	_ = screen.DrawImage(t.rectImage, op)

	return nil
}

func (t *Label) GetData() (interface{}, error) {
	return t.text, nil
}
