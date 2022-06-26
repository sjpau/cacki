package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/theonlymoby/cacki/asset"
	"github.com/theonlymoby/cacki/component"
)

func (g *Game) InitObjects() {
	if g.player == nil {
		sw, sh := ebiten.ScreenSizeInFullscreen()
		g.player = &component.Player{
			Image: asset.KidStillImage,
			O: component.Object{
				X:      float64(sw) / component.Unit,
				Y:      float64(sh) / component.Unit,
				VX:     0,
				VY:     0,
				Width:  component.Unit * int(ebiten.DeviceScaleFactor()),
				Height: component.Unit * int(ebiten.DeviceScaleFactor()),
			},
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	asset.LoadStaticImages()
}
