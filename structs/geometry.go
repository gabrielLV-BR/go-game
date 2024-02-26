package structs

import "github.com/go-gl/mathgl/mgl32"

type Circumference struct {
	Center mgl32.Vec2
	Radius float32
}

func (c Circumference) Contains(point mgl32.Vec2) bool {
	return (point.Sub(c.Center).Len()) < c.Radius
}

type Triangle struct {
	A, B, C mgl32.Vec3
}
