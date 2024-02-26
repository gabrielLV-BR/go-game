package structs

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
	index := x + y*grid.Dimensions.Y + z*grid.Dimensions.Y*grid.Dimensions.Z
	return index
}

func (grid *Grid3D[T]) Place(val T, x, y, z int) {
	grid.Data[grid.Index(x, y, z)] = val
}

func (grid *Grid3D[T]) At(x, y, z int) T {
	return grid.Data[grid.Index(x, y, z)]
}
