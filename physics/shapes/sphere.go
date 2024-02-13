package shapes

import "github.com/go-gl/mathgl/mgl32"

type SphereShape struct {
	Radius float32
}

func (sphere *SphereShape) SupportPoint(dir mgl32.Vec3) mgl32.Vec3 {
	return dir.Mul(sphere.Radius)
}

func (sphere *SphereShape) BoundingBox() BoundingBox {
	radiusVec := mgl32.Vec3{sphere.Radius, sphere.Radius, sphere.Radius}

	return BoundingBox{
		Min: radiusVec.Mul(-1),
		Max: radiusVec,
	}
}

func SpheresPenetrationVector(a SphereShape, posA mgl32.Vec3, b SphereShape, posB mgl32.Vec3) mgl32.Vec3 {
	dir := posA.Sub(posB)

	// if spheres are touching, stuff go haywire
	normalizedDir := dir.Normalize()

	return dir.Sub(normalizedDir.Mul(a.Radius)).Sub(normalizedDir.Mul(b.Radius))
}
