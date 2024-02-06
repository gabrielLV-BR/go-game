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

func (material *TextureMaterial) Prepare(shader core.Shader) {
	material.Texture.Bind(gl.TEXTURE0)
}
