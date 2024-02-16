package algorithms

import (
	"gabriellv/game/structs"

	"github.com/go-gl/mathgl/mgl32"
)

type Tuple struct {
	A, B int
}

type Triple struct {
	A, B, C int
}

func (t Triple) Edges() [3]Tuple {
	return [3]Tuple{
		{t.A, t.B},
		{t.B, t.C},
		{t.C, t.A},
	}
}

type Triangle struct {
	A, B, C mgl32.Vec2
}

type TriMesh struct {
	Vertices  []mgl32.Vec2
	Triangles []Triple
}

func (tri *TriMesh) AddTriangle(a, b, c mgl32.Vec2) {
	baseIndex := len(tri.Vertices)

	tri.Vertices = append(tri.Vertices, a, b, c)
	tri.Triangles = append(tri.Triangles, Triple{baseIndex, baseIndex + 1, baseIndex + 2})
}

// this algorithm requires vertices, so we'll be using Vec2
func DelauneyTriangulation(graph *structs.Graph[mgl32.Vec2]) {

	triangulation := TriMesh{}

	// find minimum and maximum extends of points

	var minPoint mgl32.Vec2
	var maxPoint mgl32.Vec2

	for _, point := range graph.Nodes {
		mgl32.SetMin(&minPoint[0], &point[0])
		mgl32.SetMin(&minPoint[1], &point[1])

		mgl32.SetMax(&maxPoint[0], &point[0])
		mgl32.SetMax(&maxPoint[1], &point[1])
	}

	width := maxPoint.Sub(minPoint).X()
	height := maxPoint.Sub(minPoint).Y()

	// add super triangle to triangulation that
	// bounds every point of the graph

	// just spat this out, needs testing
	triangulation.AddTriangle(
		minPoint.Sub(mgl32.Vec2{width, 0}),
		maxPoint.Sub(minPoint.Mul(0.5)).Add(mgl32.Vec2{0, height / 2}),
		minPoint.Add(mgl32.Vec2{width * 2, 0}),
	)

	for _, point := range graph.Nodes {
		badTriangles := make(map[Triple]bool)

		for _, triangle := range triangulation.Triangles {

			tri := Triangle{
				A: triangulation.Vertices[triangle.A],
				B: triangulation.Vertices[triangle.B],
				C: triangulation.Vertices[triangle.C],
			}

			circum := TriangleCircumcirle(tri.A, tri.B, tri.C)

			if circum.Contains(point) {
				badTriangles[triangle] = true
			}
		}

		//TODO finish this mess
	}
}

// func DelauneyTetrahedrilization(graph structs.Graph[mgl32.Vec3]) {}
