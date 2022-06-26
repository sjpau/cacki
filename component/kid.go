package component

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/theonlymoby/cacki/asset"
)

type Kid struct {
	O     Object
	Image *ebiten.Image
}

func (k *Kid) MoveDelay() {
	timer := time.NewTimer(time.Second * 0)
	d := rand.Intn(5)
	go func() {
		<-timer.C
		k.MoveRandomly(d)
	}()
	stopped := timer.Stop()
	if stopped {
	}
	time.Sleep(2 * time.Second)
}

func (k *Kid) DrawOn(screen *ebiten.Image) {
	s := ebiten.DeviceScaleFactor()
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(2*s, 2*s)
	o.GeoM.Translate(k.O.X/Unit, k.O.Y/Unit)
	screen.DrawImage(k.Image, o)
}

func (k *Kid) MoveRandomly(d int) {
	switch d {
	case 1:
		k.O.VX = -SpeedCoef
	case 2:
		k.O.VX = SpeedCoef
	case 3:
		k.O.VY = SpeedCoef
	case 4:
		k.O.VY = -SpeedCoef
	}
}

func (k *Kid) Update() {
	fw, fh := ebiten.ScreenSizeInFullscreen()
	if k.O.VX > 0 {
		k.Image = asset.KidRightImage
	} else if k.O.VX < 0 {
		k.Image = asset.KidLeftImage
	}
	if k.O.VY > 0 {
		k.Image = asset.KidUpImage
	} else if k.O.VY < 0 {
		k.Image = asset.KidStillImage
	}
	if k.O.VY == 0 && k.O.VX == 0 {
		k.Image = asset.KidStillImage
	}
	if k.O.X > float64(fw*(Unit-1)) {
		k.O.VX = 0
		k.O.X -= 100
	} else if k.O.X < 0 {
		k.O.VX = 0
		k.O.X += 100
	}
	if k.O.Y > float64(fh*(Unit-1)) {
		k.O.VY = 0
		k.O.Y -= 100
	} else if k.O.Y < 0 {
		k.O.VY = 0
		k.O.Y += 100
	}

	go k.MoveDelay()
	k.O.Update()
}
