package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/sjpau/cacki/asset"
	"github.com/sjpau/cacki/component"
)

const (
	maxToys = 256
)

type Toys struct {
	CachedToys  []*component.Toy
	FoodImages  []*ebiten.Image
	CarsImages  []*ebiten.Image
	PlushImages []*ebiten.Image
	Spawn       int
}

func (t *Toys) InitToysImages() {

	t.FoodImages = make([]*ebiten.Image, 4)
	t.CarsImages = make([]*ebiten.Image, 4)
	t.PlushImages = make([]*ebiten.Image, 4)
	t.FoodImages[0] = asset.FoodImage1
	t.FoodImages[1] = asset.FoodImage2
	t.FoodImages[2] = asset.FoodImage3
	t.FoodImages[3] = asset.FoodImage4
	t.CarsImages[0] = asset.CarImage1
	t.CarsImages[1] = asset.CarImage2
	t.CarsImages[2] = asset.CarImage3
	t.CarsImages[3] = asset.CarImage4
	t.PlushImages[0] = asset.PlushImage1
	t.PlushImages[1] = asset.PlushImage2
	t.PlushImages[2] = asset.PlushImage3
	t.PlushImages[3] = asset.PlushImage4
}

func (t *Toys) DrawOn(screen *ebiten.Image) {
	for i := 0; i < t.Spawn; i++ {
		if t.CachedToys[i].Image != nil {
			t.CachedToys[i].DrawOn(screen)
			t.CachedToys[i].O.Width, t.CachedToys[i].O.Height = t.CachedToys[i].Image.Size()
			t.CachedToys[i].O.Width *= 20 * int(ebiten.DeviceScaleFactor())
			t.CachedToys[i].O.Height *= 20 * int(ebiten.DeviceScaleFactor())
		}
	}
}

func (toys *Toys) TypeToImage(t *component.Toy) {
	randImage := rand.Intn(4)
	switch t.Type {
	case 1:
		t.Image = toys.CarsImages[randImage]
	case 2:
		t.Image = toys.FoodImages[randImage]
	case 3:
		t.Image = toys.PlushImages[randImage]
	}
}

func (t *Toys) Update(g *Game) {
	s := ebiten.DeviceScaleFactor()
	for i := 0; i < t.Spawn; i++ {
		t.CachedToys[i].Update()
		if t.CachedToys[i].O.CollideWith(&g.player.O) {
			t.CachedToys[i].Attached = true
			t.CachedToys[i].Stable = false
		}
		if t.CachedToys[i].Attached {
			t.CachedToys[i].O.X = g.player.O.X
			t.CachedToys[i].O.Y = g.player.O.Y + component.Unit*s*2
		}
	}
}
