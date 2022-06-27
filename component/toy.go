package component

import "github.com/hajimehoshi/ebiten"

type Toy struct {
	O        Object
	Image    *ebiten.Image
	Type     int
	Attached bool
	Spawned  bool
	Stable   bool
}

func (t *Toy) DrawOn(screen *ebiten.Image) {
	s := ebiten.DeviceScaleFactor()
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(3*s, 3*s)
	o.GeoM.Translate(t.O.X/Unit, t.O.Y/Unit)
	screen.DrawImage(t.Image, o)
	t.Spawned = true
}

func (t *Toy) Update() {
	if t.Attached == false {
		t.O.VX = 0
		t.O.VY = 0
	}
}
