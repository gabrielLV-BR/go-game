package utils

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

func AnglesToVector(yaw, pitch float32) mgl32.Vec3 {
	yaw64 := float64(yaw)
	pitch64 := float64(pitch)

	xsLen := math.Cos(pitch64)

	return mgl32.Vec3{
		float32(math.Cos(yaw64) * xsLen),
		float32(math.Sin(pitch64)),
		float32(math.Sin(-yaw64) * xsLen),
	}
}
