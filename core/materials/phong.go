package materials

import (
	"gabriellv/game/core"
	"gabriellv/game/structs"
)

type PhongMaterial struct {
	Color structs.Color
	DiffuseMaps []core.Texture
	AmbientLight core.AmbientLight
	PointLights []core.PointLight
}

func (material *PhongMaterial) Id() core.MaterialId {
	return "phong"
}

func (material *PhongMaterial) Uniforms() []core.UniformDescriptor {
	isTextured := len(material.DiffuseMaps) > 0

	return []core.UniformDescriptor {
		{
			Name: "uAmbientLight",
			Value: material.AmbientLight,
		},
		{
			Name: "uPointLight",
			Value: material.PointLights[0],
		},
		{
			Name: "uTexture",
			Value: material.DiffuseMaps,
		},
		{
			Name: "uIsTextured",
			Value: isTextured,
		},
		{
			Name: "uColor",
			Value: material.Color,
		},
	}
}

func (material *PhongMaterial) Prepare() {}

