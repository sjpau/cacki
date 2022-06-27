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
	toys          Toys
	PatienceLevel int
	StaminaLevel  int
	count         uint64
	Loss          bool
	Win           bool
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
	g.toys.DrawOn(screen)
}

func (g *Game) Update() error {
	g.InitObjects()
	ebiten.SetFullscreen(true)
	g.player.Update()
	for i := 0; i < g.toys.Spawn; i++ {
		g.baskets.Update(g.toys.CachedToys[i])
		if g.toys.CachedToys[i].Spawned && g.toys.CachedToys[i].Attached == false && g.toys.CachedToys[i].Stable == false {
			g.toys.CachedToys[i].O.X = g.kid.O.X
			g.toys.CachedToys[i].O.Y = g.kid.O.Y
			g.toys.CachedToys[i].Stable = true
		}
	}
	g.kid.Update()
	g.toys.Update(g)
	g.patience.Update(g.PatienceLevel)
	g.stamina.Update(g.StaminaLevel)
	g.EventDropToy()
	g.PatienceLevel = g.baskets.MinusPoint
	g.StaminaLevel = g.baskets.PlusPoint
	if g.StaminaLevel > 25 {
		g.Win = true
	}
	if g.PatienceLevel > 25 {
		g.Loss = true
	}
	g.count++
	return nil
}

func main() {
	ebiten.RunGame(&Game{})
}
