package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
	_ "image/png"
	"time"
)

type CheckBox struct {
	BaseWidget

	checked bool
	// image for the checkbox
	emptyImage   *ebiten.Image
	checkedImage *ebiten.Image

	lastClickedTime time.Time
}

func NewCheckBox(ID string, emptyImage string, checkedImage string, x float64, y float64) *CheckBox {
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
	cb.BaseWidget = *NewBaseWidget(ID, x, y, width, height, cb.HandleEvent)

	cb.checked = false
	cb.eventHandler = cb.HandleEvent
  cb.lastClickedTime = time.Now().UTC()

	go cb.ListenToIncomingEvents()
	return &cb
}

func (b *CheckBox) HandleEvent(event events.IEvent) (bool, error) {

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
				}
			}
		}
	}
	return false, nil
}

func (b *CheckBox) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.X, b.Y)

	// if state changed since last draw, recreate colour etc.

	if b.checked {
		_ = screen.DrawImage(b.checkedImage, op)
	} else {
		_ = screen.DrawImage(b.emptyImage, op)
	}

	return nil
}

func (b *CheckBox) GetData() (interface{}, error) {
	return b.checked, nil
}
