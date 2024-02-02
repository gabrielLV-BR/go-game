package core

import (
	"errors"
	"fmt"
	"gabriellv/game/structs/color"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	PROGRAM_MODEL_MATRIX_UNIFORM = "uModel"
	PROGRAM_VIEW_MATRIX_UNIFORM  = "uView"
	PROGRAM_PROJ_MATRIX_UNIFORM  = "uProj"
)

type Program struct {
	id           uint32
	uniformCache map[string]int32
}

func (program *Program) Bind() {
	gl.UseProgram(program.id)
}

func (program *Program) Unbind() {
	gl.UseProgram(0)
}

//

func NewProgram(shaders ...Shader) (Program, error) {
	program := Program{}

	id := gl.CreateProgram()

	for _, shader := range shaders {
		gl.AttachShader(id, shader.id)
	}

	gl.LinkProgram(id)

	err := CheckProgramError(id)

	if err != nil {
		return Program{}, err
	}

	program.id = id
	program.uniformCache = make(map[string]int32)

	return program, nil
}

// uniform interaction code

func (program *Program) SetMaterial(material Material) {
	program.SetColor(MATERIAL_COLOR_UNIFORM, material.Color)

	texture := material.Textures[0]

	texture.Bind(gl.TEXTURE0)
	program.SetTexture(MATERIAL_TEXTURE0_UNIFORM, texture)
}

func (program *Program) SetColor(uniform string, color color.Color) {
	location := program.getUniformLocation(uniform)

	r, g, b, _ := color.Unpack()

	gl.Uniform3f(location, r, g, b)
}

func (program *Program) SetMVP(model, view, projection *mgl32.Mat4) {
	program.SetMatrix(PROGRAM_MODEL_MATRIX_UNIFORM, model)
	program.SetMatrix(PROGRAM_VIEW_MATRIX_UNIFORM, view)
	program.SetMatrix(PROGRAM_PROJ_MATRIX_UNIFORM, projection)
}

func (program *Program) SetMatrix(uniform string, matrix *mgl32.Mat4) {
	location := program.getUniformLocation(uniform)

	gl.UniformMatrix4fv(location, 1, false, &matrix[0])
}

func (program *Program) SetTexture(uniform string, texture Texture) {
	location := program.getUniformLocation(uniform)

	gl.Uniform1i(location, int32(texture.Unit()-gl.TEXTURE0))
}

// set uniforms

func (program *Program) getUniformLocation(uniform string) int32 {

	location, ok := program.uniformCache[uniform]

	if !ok {
		uniformName, free := gl.Strs(uniform + "\x00")
		defer free()

		location = gl.GetUniformLocation(program.id, *uniformName)

		program.uniformCache[uniform] = location
	}

	return location
}

//

func CheckProgramError(program uint32) error {
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)

	if status == gl.TRUE {
		return nil
	}

	var infoLog [512]byte

	gl.GetProgramInfoLog(program, 512, nil, (*uint8)(unsafe.Pointer(&infoLog)))

	errorMessage := fmt.Sprintf("Program error: %s\n", string(infoLog[:512]))

	return errors.New(errorMessage)
}
