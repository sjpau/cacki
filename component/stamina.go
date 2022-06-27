package component

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/theonlymoby/cacki/asset"
)

type Stamina struct {
	O     Object
	Level int
	Image *ebiten.Image
}

func (st *Stamina) DrawOn(screen *ebiten.Image) {
	s := ebiten.DeviceScaleFactor()
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(3*s, 3*s)
	o.GeoM.Translate(st.O.X/Unit, st.O.Y/Unit)
	switch st.Level {
	case 0:
		st.Image = asset.StaminaImage1
	case 5:
		st.Image = asset.StaminaImage2
	case 10:
		st.Image = asset.StaminaImage3
	case 15:
		st.Image = asset.StaminaImage4
	case 20:
		st.Image = asset.StaminaImage5
	case 25:
		st.Image = asset.StaminaImage6
	}
	screen.DrawImage(st.Image, o)
}

func (st *Stamina) Update(l int) {
	st.Level = l
}
