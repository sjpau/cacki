package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/theonlymoby/cacki/asset"
	"github.com/theonlymoby/cacki/component"
)

func (g *Game) InitObjects() {
	if g.LossScreen == nil || g.WonScreen == nil {
		g.LossScreen = asset.LossImage
		g.WonScreen = asset.WonImage
	}
	s := ebiten.DeviceScaleFactor()
	fw, fh := ebiten.ScreenSizeInFullscreen()
	if g.player == nil {
		g.Loss = false
		g.Win = false
		g.PatienceLevel = 0
		g.StaminaLevel = 0
		g.player = &component.Player{
			Image: asset.KidStillImage,
			O: component.Object{
				X:  float64(fw) * (component.Unit - 1) / 2,
				Y:  float64(fh) * (component.Unit - 1) / 2,
				VX: 0,
				VY: 0,
			},
		}
		pw, ph := g.player.Image.Size()
		g.player.O.Width = pw * int(s)
		g.player.O.Height = ph * int(s)
	}

	if g.kid == nil {
		g.kid = &component.Kid{
			Image: asset.KidStillImage,
			O: component.Object{
				X:  float64(fw)*(component.Unit-1)/2 + float64(rand.Intn(100)-rand.Intn(100)*component.Unit*rand.Intn(5)),
				Y:  float64(fh)*(component.Unit-1)/2 + float64(rand.Intn(100)-rand.Intn(100)*component.Unit*rand.Intn(5)),
				VX: 0,
				VY: 0,
			},
		}
		pw, ph := g.kid.Image.Size()
		g.kid.O.Width = pw * int(s)
		g.kid.O.Height = ph * int(s)
	}

	if g.stamina == nil {
		g.stamina = &component.Stamina{
			O: component.Object{
				X:      float64(fw) / component.Unit,
				Y:      float64(fh) / component.Unit,
				Width:  component.Unit * int(s),
				Height: component.Unit * int(s),
			},
		}
	}

	if g.patience == nil {
		g.patience = &component.Patience{
			O: component.Object{
				X:      float64(fw) * (component.Unit - 1),
				Y:      float64(fh) / component.Unit,
				Width:  component.Unit * int(s),
				Height: component.Unit * int(s),
			},
		}
	}

	if g.baskets.ToyBaskets == nil {
		g.baskets.ToyBaskets = make([]*component.Basket, 3)
		for i := range g.baskets.ToyBaskets {
			g.baskets.ToyBaskets[i] = &component.Basket{
				O: component.Object{
					X:      float64(fw*(component.Unit-1)/2) + float64(6*(i+2)*component.Unit*component.Unit) - component.Unit*component.Unit*16,
					Y:      0 + 1*(component.Unit-1),
					Width:  50 * component.Unit * int(s),
					Height: 50 * component.Unit * int(s),
				},
			}
		}
	}

	if g.toys.CachedToys == nil {
		g.toys.InitToysImages()
		g.toys.CachedToys = make([]*component.Toy, maxToys)
		g.toys.Spawn = 0
		for i := range g.toys.CachedToys {
			randType := rand.Intn(4-1) + 1
			g.toys.CachedToys[i] = &component.Toy{
				Type: randType,
				O: component.Object{
					X:      g.kid.O.X,
					Y:      g.kid.O.Y,
					Width:  50 * component.Unit * int(s),
					Height: 50 * component.Unit * int(s),
				},
				Attached: false,
				Spawned:  false,
				Stable:   false,
			}
			g.toys.TypeToImage(g.toys.CachedToys[i])
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	asset.LoadStaticImages()
}
