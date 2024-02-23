package structs

type Color struct {
	R float32
	G float32
	B float32
	A float32
}

func RGB(r, g, b float32) Color {
	return Color{
		R: r,
		G: g,
		B: b,
		A: 1.0,
	}
}

func (c Color) Unpack() (float32, float32, float32, float32) {
	return c.R, c.G, c.B, c.A
}

var Colors Color

func (c Color) White() Color {
	return RGB(1.0, 1.0, 1.0)
}

func (c Color) Black() Color {
	return RGB(0.0, 0.0, 0.0)
}
