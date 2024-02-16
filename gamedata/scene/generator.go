package scene

import (
	"gabriellv/game/algorithms"
	"gabriellv/game/structs"
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
)

type LevelGenerator struct {
	RoomCount int
}

type RoomNode struct {
	Position   mgl32.Vec3
	Dimensions mgl32.Vec3
}

func NewLevelGenerator() LevelGenerator {
	return LevelGenerator{}
}

func (level *LevelGenerator) Generate() {
	// first, create a grid to represent world

	roomGraph := structs.Graph[mgl32.Vec2]{}
	roomGraph.New()

	for i := 0; i < level.RoomCount; i++ {

		x := rand.Float32()
		y := rand.Float32()

		point := mgl32.Vec2{x, y}

		roomGraph.AddNode(point)
	}

	// link every node

	algorithms.DelauneyTriangulation(&roomGraph)

	// remove excess edges
	// build room meesh
}

type Grid3D[T any] struct {
	Data       []T
	Dimensions struct{ X, Y, Z int }
}

func (grid *Grid3D[T]) New(x, y, z int) {
	grid.Data = make([]T, x*y*z)
	grid.Dimensions.X = x
	grid.Dimensions.Y = y
	grid.Dimensions.Z = z
}

func (grid *Grid3D[T]) Index(x, y, z int) int {
	xy := grid.Dimensions.X * grid.Dimensions.Y
	index := (x % grid.Dimensions.X) + ((y * grid.Dimensions.X) % grid.Dimensions.Y) + ((z * xy) % grid.Dimensions.Z)

	return index
}

func (grid *Grid3D[T]) Place(val T, x, y, z int) {
	grid.Data[grid.Index(x, y, z)] = val
}

func (grid *Grid3D[T]) At(x, y, z int) T {
	return grid.Data[grid.Index(x, y, z)]
}
