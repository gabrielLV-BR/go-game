package materials

import (
	"gabriellv/game/core"
	"gabriellv/game/structs"
)

type PhongMaterial struct {
	Color structs.Color
	Diffuse core.Texture
}

func (material *PhongMaterial) Id() core.MaterialId {
	return "phong"
}

func (material *PhongMaterial) Uniforms() []core.UniformDescriptor {
	return []core.UniformDescriptor {
		{
			Name: "uTexture",
			Value: material.Diffuse,
		},
		{
			Name: "uColor",
			Value: material.Color,
		},
	}
}

func (material *PhongMaterial) Prepare() {}

