package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/theonlymoby/cacki/component"
)

type Game struct {
	player *component.Player
}

func (g *Game) Layout(w, h int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(w) * s), int(float64(h) * s)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.DrawOn(screen)
}

func (g *Game) Update() error {
	g.InitObjects()
	ebiten.SetFullscreen(true)
	g.player.Update()
	return nil
}

func main() {
	ebiten.RunGame(&Game{})
}
