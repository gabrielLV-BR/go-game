package scene

type LevelGenerator struct {
	
}

type Grid3D[T any] struct {
	Data []T
	Dimensions struct { X, Y, Z int }
}

func (grid *Grid3D[T]) New(x, y, z int) {
	grid.Data = make([]T, x * y * z)
	grid.Dimensions.X = x
	grid.Dimensions.Y = y
	grid.Dimensions.Z = z
}

func (grid *Grid3D[T]) At(x, y, z int) T {
}

func NewLevelGenerator() LevelGenerator {
	return LevelGenerator{}
}

func  (level *LevelGenerator) Generate() {
	// first, create 3D grid to represent world

	grid := Grid3D[uint32]{}
	grid.New(10, 10, 10)
}
