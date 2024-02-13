package physics

import "github.com/go-gl/mathgl/mgl32"

type BoundingBox struct {
	Min mgl32.Vec3
	Max mgl32.Vec3
}

func NewBoundingBox(x, y, w, h float32) BoundingBox {
	return BoundingBox{
		Min: mgl32.Vec3{x, y},
		Max: mgl32.Vec3{x + w, y + h},
	}
}

func BoundingBoxCollide(a, b BoundingBox) bool {
	//TODO this NEEDS to be tested, I just spat this out
	//TODO check if accessing members through these methods
	// isn't introducing any unnecessary overhead
	return a.Min.X() < b.Max.X() &&
		a.Max.X() > b.Min.X() &&
		a.Min.Y() > b.Max.Y() &&
		a.Max.Y() < b.Min.Y() &&
		a.Max.Z() < b.Min.Z() &&
		a.Min.Z() > b.Max.Z()
}
