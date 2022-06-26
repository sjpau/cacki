package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/theonlymoby/cacki/asset"
	"github.com/theonlymoby/cacki/component"
)

func (g *Game) InitObjects() {
	fw, fh := ebiten.ScreenSizeInFullscreen()
	if g.player == nil {
		g.player = &component.Player{
			Image: asset.KidStillImage,
			O: component.Object{
				X:      float64(fw) / component.Unit,
				Y:      float64(fh) / component.Unit,
				VX:     0,
				VY:     0,
				Width:  component.Unit * int(ebiten.DeviceScaleFactor()),
				Height: component.Unit * int(ebiten.DeviceScaleFactor()),
			},
		}
	}

	if g.stamina == nil {
		g.stamina = &component.Stamina{
			O: component.Object{
				X:      float64(fw) / component.Unit,
				Y:      float64(fh) / component.Unit,
				Width:  component.Unit * int(ebiten.DeviceScaleFactor()),
				Height: component.Unit * int(ebiten.DeviceScaleFactor()),
			},
		}
	}

	if g.patience == nil {
		g.patience = &component.Patience{
			O: component.Object{
				X:      float64(fw) * (component.Unit - 1),
				Y:      float64(fh) / component.Unit,
				Width:  component.Unit * int(ebiten.DeviceScaleFactor()),
				Height: component.Unit * int(ebiten.DeviceScaleFactor()),
			},
		}
	}

	if g.baskets.ToyBaskets == nil {
		g.baskets.ToyBaskets = make([]*component.Basket, 3)
		for i := range g.baskets.ToyBaskets {
			g.baskets.ToyBaskets[i] = &component.Basket{
				O: component.Object{
					X:      float64(fw*(component.Unit-1)/2) + float64(6*i*component.Unit*component.Unit),
					Y:      0 + 1*(component.Unit-1),
					Width:  component.Unit * int(ebiten.DeviceScaleFactor()),
					Height: component.Unit * int(ebiten.DeviceScaleFactor()),
				},
			}
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	asset.LoadStaticImages()
}
