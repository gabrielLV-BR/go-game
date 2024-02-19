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

func (builder *MeshBuilder) AddTriangle(a, b, c mgl32.Vec3) {
	vertices := []float32{
		a[0], a[1], a[2],
		b[0], b[1], b[2],
		c[0], c[1], c[2],
	}

	indices := []uint32{
		0, 1, 2,
	}

	mesh := core.Mesh{
		Vertices: vertices,
		Indices:  indices,
	}

	builder.Meshes = append(builder.Meshes, mesh)
}

func (builder *MeshBuilder) Build(removeDuplicates bool) core.Mesh {
	// save some reallocs :D
	numVertices := 0
	numIndices := 0

	for _, mesh := range builder.Meshes {
		numVertices += len(mesh.Vertices)
		numIndices += len(mesh.Indices)
	}

	vertices := make([]float32, numVertices)
	indices := make([]uint32, numIndices)

	if removeDuplicates {
		//TODO same logic is in object loader, maybe don't repeat?
		// keep yourself DRY man!!
		indexMap := make(map[mgl32.Vec3]uint32)
		index := uint32(0)

		for _, mesh := range builder.Meshes {
			for i := 0; i < len(mesh.Indices); i += 3 {
				v := mgl32.Vec3{
					mesh.Vertices[i],
					mesh.Vertices[i+1],
					mesh.Vertices[i+2],
				}

				maybeIndex, ok := indexMap[v]

				if ok {
					indices = append(indices, maybeIndex)
				} else {
					indexMap[v] = index
					vertices = append(vertices, v[0], v[1], v[2])
					indices = append(indices, index)
					index += 1
				}
			}
		}
	} else {
		for _, mesh := range builder.Meshes {
			for _, v := range mesh.Vertices {
				vertices = append(vertices, v)
			}

			indexOffset := len(mesh.Vertices)
			for _, i := range mesh.Indices {
				indices = append(indices, i+uint32(indexOffset))
			}
		}
	}

	return core.Mesh{
		Vertices: vertices,
		Indices:  indices,
	}
}
