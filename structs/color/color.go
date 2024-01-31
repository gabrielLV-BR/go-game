package color

type Color struct {
	R float32
	G float32
	B float32
	A float32
}

func (c Color) Unpack() (float32, float32, float32, float32) {
	return c.R, c.G, c.B, c.A
}
