package structs

import "github.com/go-gl/mathgl/mgl32"

type Transform struct {
	Position mgl32.Vec3
	Scale    mgl32.Vec3
	Rotation mgl32.Quat
}

func NewTransform(position, scale mgl32.Vec3, rotation mgl32.Quat) Transform {
	return Transform{
		Position: position,
		Scale:    scale,
		Rotation: rotation,
	}
}
