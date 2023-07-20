package background

import "github.com/hajimehoshi/ebiten/v2"

type Background struct {
	images []*ebiten.Image
	offset float32
}

func New(images []*ebiten.Image) *Background {
	return &Background{
		images: images,
	}
}

func (b *Background) Update() {
	b.offset += 0.5
	if int(b.offset) > len(b.images)-1 {
		b.offset = 0
	}
}

func (b *Background) Draw(screen *ebiten.Image) {
	screen.DrawImage(b.images[int(b.offset)], nil)
}
