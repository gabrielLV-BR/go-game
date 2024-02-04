package materials

import (
	"gabriellv/game/core"
)

type ColorMaterial struct {
}

func (material *ColorMaterial) Id() core.MaterialId {
	return core.MaterialId("color")
}

func (material *ColorMaterial) Prepare() {
}

func (material *ColorMaterial) Uniforms() []core.UniformDescriptor {
	return []core.UniformDescriptor{}
}
