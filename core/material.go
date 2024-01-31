package core

import "gabriellv/game/structs/color"

const (
	MATERIAL_COLOR_UNIFORM    = "uColor"
	MATERIAL_TEXTURE0_UNIFORM = "uTexture0"
)

type Material struct {
	Color    color.Color
	Textures []Texture
}

// this ID is calculated based on the textures present
// and allows for binding between material and Program
// without redundant fields
func (material *Material) Id() uint32 {
	return uint32(len(material.Textures))
}
