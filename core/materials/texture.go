package materials

import (
	"gabriellv/game/core"
	"gabriellv/game/structs"
)

type TextureMaterial struct {
	Color   structs.Color
	Texture core.Texture
}

func (material *TextureMaterial) Id() core.MaterialId {
	return core.MaterialId("texture")
}

func (material *TextureMaterial) Prepare(shader core.Shader) {
	shader.SetColor("uColor", material.Color)
	// shader.SetTexture("uTexture0", material.Texture)
}
