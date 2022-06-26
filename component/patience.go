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
	case 1:
		p.Image = asset.PatienceImage1
	case 2:
		p.Image = asset.PatienceImage2
	case 3:
		p.Image = asset.PatienceImage3
	case 4:
		p.Image = asset.PatienceImage4
	case 5:
		p.Image = asset.PatienceImage5
	case 6:
		p.Image = asset.PatienceImage6
	}
	screen.DrawImage(p.Image, o)
}

func (p *Patience) Update(l int) {
	p.Level = l
}
