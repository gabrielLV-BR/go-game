package core

const (
	MATERIAL_COLOR_UNIFORM    = "uColor"
	MATERIAL_TEXTURE0_UNIFORM = "uTexture0"
)

type MaterialId string

type Material interface {
	Id() MaterialId
	Uniforms() []UniformDescriptor

	Prepare()
}
