package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/theonlymoby/cacki/component"
)

type Game struct {
	player        *component.Player
	kid           *component.Kid
	baskets       component.Baskets
	patience      *component.Patience
	stamina       *component.Stamina
	PatienceLevel int
	StaminaLevel  int
}

func (g *Game) Layout(w, h int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(w) * s), int(float64(h) * s)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.DrawOn(screen)
	g.kid.DrawOn(screen)
	g.baskets.DrawOn(screen)
	g.stamina.DrawOn(screen)
	g.patience.DrawOn(screen)
}

func (g *Game) Update() error {
	g.InitObjects()
	ebiten.SetFullscreen(true)
	g.player.Update()
	g.kid.Update()
	g.baskets.Update()
	g.patience.Update(1)
	g.stamina.Update(1)
	return nil
}

func main() {
	ebiten.RunGame(&Game{})
}
