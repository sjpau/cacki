package component

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/theonlymoby/cacki/asset"
)

const (
	TypeCars = iota
	TypeFood
	TypePlush
)

type Basket struct {
	O     Object
	Image *ebiten.Image
	Type  int
}

type Baskets struct {
	ToyBaskets []*Basket
}

func (b *Basket) DrawOn(screen *ebiten.Image, t int) {
	s := ebiten.DeviceScaleFactor()
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(3*s, 3*s)
	o.GeoM.Translate(b.O.X/Unit, b.O.Y/Unit)
	switch t {
	case TypeCars:
		b.Image = asset.BasketcarsImage
	case TypeFood:
		b.Image = asset.BasketfoodImage
	case TypePlush:
		b.Image = asset.BasketplushImage
	}
	screen.DrawImage(b.Image, o)
}

func (b *Baskets) DrawOn(screen *ebiten.Image) {
	b.ToyBaskets[0].DrawOn(screen, TypeCars)
	b.ToyBaskets[1].DrawOn(screen, TypeFood)
	b.ToyBaskets[2].DrawOn(screen, TypePlush)
}

func (b *Baskets) Update(t *Toy) {
	for i := 0; i < 3; i++ {
		if b.ToyBaskets[i].O.CollideWith(&t.O) {
			if b.ToyBaskets[i].Type == t.Type {
				// plus point
			} else {
				//minus point
			}
		}
	}
}
