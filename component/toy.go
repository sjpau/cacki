package component

import "github.com/hajimehoshi/ebiten"

type Toy struct {
	o     Object
	Image *ebiten.Image
}

type FoodToys struct {
	Toys         []*Toy
	CachedImages []*ebiten.Image
}

type CarToys struct {
	Toys         []*Toy
	CachedImages []*ebiten.Image
}

type PlushToys struct {
	Toys         []*Toy
	CachedImages []*ebiten.Image
}
