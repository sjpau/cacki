package component

type Object struct {
	X      float64
	Y      float64
	VX     float64
	VY     float64
	Width  int
	Height int
}

const (
	SpeedCoef = 100
)

func (o *Object) CollideWith(n *Object) bool {
	if o.X+float64(o.Width) >= n.X &&
		o.X <= n.X+float64(n.Width) &&
		o.Y+float64(o.Height) >= n.Y &&
		o.Y <= n.Y+float64(n.Height) {
		return true
	}
	return false
}

func (o *Object) Update() {
	o.X += o.VX
	o.Y += o.VY
}
