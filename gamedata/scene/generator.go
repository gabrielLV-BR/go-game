package scene

import (
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
	// first, create 3D grid to represent world

	roomGraph := structs.Graph[RoomNode]{}
	roomGraph.New()

	for i := 0; i < level.RoomCount; i++ {

		x := rand.Float32()
		y := rand.Float32()
		z := rand.Float32()

		width := rand.Float32()
		height := rand.Float32()

		room := RoomNode{
			Position:   mgl32.Vec3{x, y, z},
			Dimensions: mgl32.Vec3{width, 1, height},
		}

		roomGraph.AddNode(room)
	}

	// link every node
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
