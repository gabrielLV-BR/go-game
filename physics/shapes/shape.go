package shapes

import (
	"github.com/go-gl/mathgl/mgl32"
)

// The way these are structured, the shape should return its
// support point and bounding box centered at (0, 0)
// I'm currently doing that because it's more versatile.
// Maybe I'll put it to some better use in the future
type PhysicsShape interface {
	SupportPoint(dir mgl32.Vec3) mgl32.Vec3
	BoundingBox() BoundingBox
}

// Bounding Box

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

func (box BoundingBox) CenteredAt(center mgl32.Vec3) BoundingBox {
	return BoundingBox{
		Min: box.Min.Add(center),
		Max: box.Max.Add(center),
	}
}

func (box BoundingBox) Scaled(scale mgl32.Vec3) BoundingBox {
	// should disable bounds check in each array access
	// hyahyahyahya
	if len(box.Min) != len(scale) || len(box.Max) != len(scale) {
		panic("Impossible: all should be Vec3")
	}

	for i := 0; i < 3; i++ {
		box.Min[i] *= scale[i]
		box.Max[i] *= scale[i]
	}

	return box
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
