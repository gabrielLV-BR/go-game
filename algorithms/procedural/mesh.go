package procedural

import (
	"gabriellv/game/core"

	"github.com/go-gl/mathgl/mgl32"
)

type MeshBuilder struct {
	Meshes []core.Mesh
}

func (builder *MeshBuilder) New() {
	builder.Meshes = []core.Mesh{}
}

func (builder *MeshBuilder) AddBox(position, size mgl32.Vec3) {
	vertices := []float32{
		-0.5, -0.5, 0.5,
		-0.5, 0.5, 0.5,
		0.5, 0.5, 0.5,
		0.5, -0.5, 0.5,
		-0.5, -0.5, -0.5,
		-0.5, 0.5, -0.5,
		0.5, 0.5, -0.5,
		0.5, -0.5, -0.5,
	}

	indices := []uint32{
		// front face
		0, 1, 2,
		2, 3, 0,
		// right face
		3, 2, 6,
		6, 7, 3,
		// back face
		7, 6, 5,
		5, 4, 7,
		// left face
		4, 5, 1,
		1, 0, 4,
		// bottom face
		4, 0, 3,
		3, 7, 4,
		// top face
		1, 5, 6,
		6, 2, 1,
	}

	mesh := core.Mesh{
		Vertices: vertices,
		Indices:  indices,
	}

	builder.Meshes = append(builder.Meshes, mesh)
}

func (builder *MeshBuilder) Build() core.Mesh {
	return core.Mesh{}
}
