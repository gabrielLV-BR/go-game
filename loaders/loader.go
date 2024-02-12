package loaders

import "gabriellv/game/core"

type ObjectLoader interface {
	LoadFromFile(path string) (Object, error)
}

type MaterialLoader interface {
	LoadFromFile(path string) (core.Material, error)
}

type Object struct {
	Vertices []float32
	Indices  []uint32
}
