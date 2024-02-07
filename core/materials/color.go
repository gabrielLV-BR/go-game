package materials

import (
	"gabriellv/game/core"
	"gabriellv/game/structs"
)

type ColorMaterial struct {
	Color structs.Color
}

func (material *ColorMaterial) Id() core.MaterialId {
	return core.MaterialId("color")
}

func (material *ColorMaterial) Prepare(shader core.Shader) {
	shader.SetColor("uColor", material.Color)
}
