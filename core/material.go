package core

import (
	"gabriellv/game/structs"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	MATERIAL_COLOR_UNIFORM    = "uColor"
	MATERIAL_TEXTURE0_UNIFORM = "uTexture0"
)

type MaterialId string

type Material interface {
	Id() MaterialId
	Uniforms() []UniformDescriptor

	Prepare()
}

// material properties

type AmbientLight struct {
	Color structs.Color
	Intensity float32
}

type PointLight struct {
	Position mgl32.Vec3
	Color structs.Color
}
