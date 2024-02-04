package materials

import (
	"gabriellv/game/core"
	"gabriellv/game/structs"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type TextureMaterial struct {
	Color   structs.Color
	Texture core.Texture
}

func (material *TextureMaterial) Id() core.MaterialId {
	return core.MaterialId("texture")
}

func (material *TextureMaterial) Prepare() {
	material.Texture.Bind(gl.TEXTURE0)
}

func (material *TextureMaterial) Uniforms() []core.UniformDescriptor {
	return []core.UniformDescriptor{
		{
			Name:  core.MATERIAL_TEXTURE0_UNIFORM,
			Value: material.Texture,
		},
		{
			Name:  core.MATERIAL_COLOR_UNIFORM,
			Value: material.Color,
		},
	}
}
