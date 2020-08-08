package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/kpfaulkner/goui/pkg/common"
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"image/color"
	_ "image/png"
	"math"
	"time"
)

const (
	checkedImage = "images/checkedcheckbox.png"
	emptyImage = "images/emptycheckbox.png"
)

type CheckBox struct {
	BaseWidget
	checked bool

	// image for the checkbox
	emptyImage   *ebiten.Image
	checkedImage *ebiten.Image
	fontInfo         common.Font
	text string
	lastClickedTime time.Time
}

func init() {
	defaultFontInfo = common.LoadFont("", 16, color.RGBA{0xff, 0xff, 0xff, 0xff})
}

func NewCheckBox(ID string, text string, handler func(event events.IEvent) error) *CheckBox {
	cb := CheckBox{}

	img1, err := loadImage(emptyImage)
	if err != nil {
		log.Fatalf("Unable to load image %s", emptyImage)
	}
	cb.emptyImage = img1

	img2, err := loadImage(checkedImage)
	if err != nil {
		log.Fatalf("Unable to load image %s", checkedImage)
	}
	cb.checkedImage = img2


	width, height := cb.emptyImage.Size()
	cb.BaseWidget = *NewBaseWidget(ID, width, height, handler)
	cb.checked = false
	cb.lastClickedTime = time.Now().UTC()
	cb.text = text
	cb.fontInfo = defaultFontInfo
	cb.setupCheckboxImage()
	return &cb
}

// setupCheckboxImage... setups up image + size + words
func (b *CheckBox) setupCheckboxImage() error {

	w,h := b.emptyImage.Size()
	bounds, _ := font.BoundString(b.fontInfo.UIFont, b.text)
	totalWidth := w + (bounds.Max.X - bounds.Min.X).Ceil() + 10  // 10 pixels between image and text?
	totalHeight := math.Max(float64(h), float64((bounds.Max.Y - bounds.Min.Y).Ceil()))

	b.Width = totalWidth
	b.Height = int(totalHeight)

	b.rectImage,_ = ebiten.NewImage(b.Width, b.Height, ebiten.FilterDefault)
	b.rectImage.Fill( color.Black)

	text.Draw(b.rectImage, b.text, b.fontInfo.UIFont, w + 5, 15, color.RGBA{0,0xff,0,0xff})

	return nil
}

func (b *CheckBox) HandleEvent(event events.IEvent) error {

	eventType := event.EventType()
	switch eventType {
	case events.EventTypeButtonDown:
		{
			// if recent click, then ignore it (otherwise just constant on/off)
			now := time.Now().UTC()
			if now.Sub(b.lastClickedTime) > 100*time.Millisecond {
				b.lastClickedTime = now
				mouseEvent := event.(events.MouseEvent)
				// check click is in button boundary.
				if b.ContainsCoords(mouseEvent.X, mouseEvent.Y) {
					b.hasFocus = true
					// if already pressed, then skip it.. .otherwise lots of repeats.
					b.checked = !b.checked

					cbe := events.NewCheckBoxEvent(event.Name(), event.EventType(), b.checked)
					b.eventHandler(cbe)
					b.stateChangedSinceLastDraw = true
				}
			}
		}
	}
	return nil
}

func (b *CheckBox) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(0,0)
	// if state changed since last draw, recreate colour etc.

	//if b.stateChangedSinceLastDraw {
	if true {
		if b.checked {
			_ = b.rectImage.DrawImage(b.checkedImage, op)
		} else {
			_ = b.rectImage.DrawImage(b.emptyImage, op)
		}
		b.stateChangedSinceLastDraw = false
	}

	op.GeoM.Translate(b.X, b.Y)
	_ = screen.DrawImage(b.rectImage, op)

	return nil
}

func (b *CheckBox) GetData() (interface{}, error) {
	return b.checked, nil
}
