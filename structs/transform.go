package structs

import "github.com/go-gl/mathgl/mgl32"

type Transform struct {
	Position mgl32.Vec3
	Scale    mgl32.Vec3
	Rotation mgl32.Quat
}

func NewTransform() Transform {
	return Transform{
		Position: mgl32.Vec3{0, 0, 0},
		Scale:    mgl32.Vec3{1.0, 1.0, 1.0},
		Rotation: mgl32.QuatIdent(),
	}
}

func (transform *Transform) GetModelMatrix() mgl32.Mat4 {
	rot := transform.Rotation.Mat4()
	trans := mgl32.Translate3D(transform.Position.X(), transform.Position.Y(), transform.Position.Z())
	scale := mgl32.Scale3D(transform.Scale.X(), transform.Scale.Y(), transform.Scale.Z())

	return rot.Mul4(trans).Mul4(scale)
}
