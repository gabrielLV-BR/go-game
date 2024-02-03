package utils

import (
	math "github.com/chewxy/math32"

	"github.com/go-gl/mathgl/mgl32"
)

func AnglesToVector(yaw, pitch float32) mgl32.Vec3 {
	xsLen := math.Cos(pitch)

	return mgl32.Vec3{
		math.Cos(yaw) * xsLen,
		math.Sin(pitch),
		math.Sin(-yaw) * xsLen,
	}
}