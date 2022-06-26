package component

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/theonlymoby/cacki/asset"
)

type Player struct {
	O     Object
	Image *ebiten.Image
}

func (p *Player) DrawOn(screen *ebiten.Image) {
	s := ebiten.DeviceScaleFactor()
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(3*s, 3*s)
	o.GeoM.Translate(p.O.X/Unit, p.O.Y/Unit)
	screen.DrawImage(p.Image, o)
}

func (p *Player) Update() {
	fw, fh := ebiten.ScreenSizeInFullscreen()
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

	/*WTF?!?!?!?!?!*/
	if p.O.X > float64(fw*(Unit-1)) {
		p.O.VX = 0
		p.O.X -= 100
	} else if p.O.X < 0 {
		p.O.VX = 0
		p.O.X += 100
	}
	if p.O.Y > float64(fh*(Unit-1)) {
		p.O.VY = 0
		p.O.Y -= 100
	} else if p.O.Y < 0 {
		p.O.VY = 0
		p.O.Y += 100
	}

	p.O.Update()
	p.O.VX = 0
	p.O.VY = 0
}
