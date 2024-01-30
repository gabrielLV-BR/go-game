package core

import "gabriellv/game/structs"

type Material struct {
	color structs.Color
	textures []uint32
}

// this ID is calculated based on the textures present
// and allows for binding between material and Program
// without redundant fields
func (material *Material) Id() uint32 {
	return uint32(len(material.textures))
}
