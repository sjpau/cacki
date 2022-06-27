package component

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/theonlymoby/cacki/asset"
)

type Patience struct {
	O     Object
	Level int
	Image *ebiten.Image
}

func (p *Patience) DrawOn(screen *ebiten.Image) {
	s := ebiten.DeviceScaleFactor()
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(3*s, 3*s)
	o.GeoM.Translate(p.O.X/Unit, p.O.Y/Unit)
	switch p.Level {
	case 0:
		p.Image = asset.PatienceImage1
	case 5:
		p.Image = asset.PatienceImage2
	case 10:
		p.Image = asset.PatienceImage3
	case 15:
		p.Image = asset.PatienceImage4
	case 20:
		p.Image = asset.PatienceImage5
	case 25:
		p.Image = asset.PatienceImage6
	}
	screen.DrawImage(p.Image, o)
}

func (p *Patience) Update(l int) {
	p.Level = l
}
