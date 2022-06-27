package main

import (
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/theonlymoby/cacki/component"
)

type Mode int

const (
	ModeGame = iota
	ModeResult
)

type Game struct {
	mode          Mode
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
	LossScreen    *ebiten.Image
	WonScreen     *ebiten.Image
}

func (g *Game) Layout(w, h int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(w) * s), int(float64(h) * s)
}

func (g *Game) Draw(screen *ebiten.Image) {
	s := ebiten.DeviceScaleFactor()
	fw, fh := ebiten.ScreenSizeInFullscreen()
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(4*s, 4*s)
	o.GeoM.Translate(float64(fw)/2, float64(fh)/2)

	switch g.mode {
	case ModeGame:
		g.player.DrawOn(screen)
		g.kid.DrawOn(screen)
		g.baskets.DrawOn(screen)
		g.stamina.DrawOn(screen)
		g.patience.DrawOn(screen)
		g.toys.DrawOn(screen)
	case ModeResult:
		if g.Win == true {
			screen.DrawImage(g.WonScreen, o)
		}
		if g.Loss == true {
			screen.DrawImage(g.LossScreen, o)
		}
	}
}

func (g *Game) Update() error {
	ebiten.SetFullscreen(true)
	switch g.mode {
	case ModeGame:
		g.InitObjects()
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
		if g.Win == true || g.Loss == true {
			g.mode = ModeResult
		}
	case ModeResult:
		//???
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			os.Exit(1)
		}
	}
	return nil
}

func main() {
	ebiten.RunGame(&Game{})
}
