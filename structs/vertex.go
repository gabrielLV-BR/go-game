package structs

import "github.com/go-gl/mathgl/mgl32"

type Vertex struct {
	Position mgl32.Vec3
	Normal   mgl32.Vec3
	UV       mgl32.Vec2
}
