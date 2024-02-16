package algorithms

import (
	"gabriellv/game/structs"

	"github.com/go-gl/mathgl/mgl32"
)

func TriangleCircumcirle(a, b, c mgl32.Vec2) structs.Circumference {
	var center mgl32.Vec2
	var radius float32

	center = center.Add(a).Add(b).Add(c).Mul(1 / 3)
	radius = center.Sub(a).Len()

	return structs.Circumference{
		Center: center,
		Radius: radius,
	}
}
