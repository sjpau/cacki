package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Game struct{}

func (g *Game) Layout(w, h int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(w) * s), int(float64(h) * s)
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Update() error {
	return nil
}

func main() {
}
