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

func (material *PhongMaterial) Prepare(shader core.Shader) {
}

