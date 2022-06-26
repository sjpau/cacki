package component

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/theonlymoby/cacki/asset"
)

type Player struct {
	O     Object
	Image *ebiten.Image
}

func (p *Player) DrawOn(screen *ebiten.Image) {
	if p.Image != nil {
		s := ebiten.DeviceScaleFactor()
		o := &ebiten.DrawImageOptions{}
		o.GeoM.Scale(4*s, 4*s)
		o.Filter = ebiten.FilterLinear
		o.GeoM.Translate(p.O.X/Unit, p.O.Y/Unit)
		screen.DrawImage(p.Image, o)
	}
}

func (p *Player) Update() {
	fw, fh := ebiten.ScreenSizeInFullscreen()
	fmt.Printf("w %d, h %d", fw, fh)
	l, r, u, d := BindMovementKeys()
	if l {
		p.Image = asset.KidLeftImage
		p.O.VX = -SpeedCoef
	} else if r {
		p.Image = asset.KidRightImage
		p.O.VX = SpeedCoef
	}
	if u {
		p.Image = asset.KidUpImage
		p.O.VY = -SpeedCoef
	} else if d {
		p.Image = asset.KidStillImage
		p.O.VY = SpeedCoef
	}
	if p.O.VX != 0 && p.O.VY != 0 {
		p.O.VX = 0
		p.O.VY = 0
	}
	p.O.Update()
	p.O.VX = 0
	p.O.VY = 0
}
