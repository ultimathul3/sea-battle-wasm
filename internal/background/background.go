package background

import "github.com/hajimehoshi/ebiten/v2"

type Background struct {
	images         []*ebiten.Image
	offset         int
	clock          int
	dir            bool
	animationSpeed uint8
}

func New(images []*ebiten.Image, animationSpeed uint8) *Background {
	return &Background{
		images:         images,
		animationSpeed: animationSpeed,
	}
}

func (b *Background) Update() error {
	if b.clock > int(b.animationSpeed) {
		if b.dir {
			b.offset++
		} else {
			b.offset--
		}
		b.clock = 0
	}

	b.clock++

	if b.offset < 0 {
		b.offset = 0
		b.dir = !b.dir
	}
	if b.offset > len(b.images)-1 {
		b.offset = len(b.images) - 1
		b.dir = !b.dir
	}

	return nil
}

func (b *Background) Draw(screen *ebiten.Image) {
	screen.DrawImage(b.images[b.offset], nil)
}
