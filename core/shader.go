package core

import (
	"errors"
	"fmt"
	"unsafe"

	"gabriellv/game/structs"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	PROGRAM_MODEL_MATRIX_UNIFORM = "uModel"
	PROGRAM_VIEW_MATRIX_UNIFORM  = "uView"
	PROGRAM_PROJ_MATRIX_UNIFORM  = "uProj"
)

type Shader struct {
	program uint32
	stages []uint32
	uniformCache map[string]int32
}

func NewShader() Shader {
	return Shader{
		program: 0,
		stages: []uint32{},
		uniformCache: make(map[string]int32),
	}
}

func (shader *Shader) LoadStageSource(source string, stage uint32) error {
	handle := gl.CreateShader(stage)

	strings, free := gl.Strs(source + "\x00")
	defer free()

	gl.ShaderSource(handle, 1, strings, nil)
	gl.CompileShader(handle)

	err := checkShaderError(handle)

	if err != nil {
		return err
	}

	shader.stages = append(shader.stages, handle)

	return nil
}

func (shader *Shader) Link() error {
	shader.program = gl.CreateProgram()

	for _, stage := range shader.stages {
		gl.AttachShader(shader.program, stage)
	}

	gl.LinkProgram(shader.program)

	return checkProgramError(shader.program)
}

func (shader *Shader) Bind() {
	gl.UseProgram(shader.program)
}

func (shader *Shader) Unbind() {
	gl.UseProgram(0)
}
// uniform handling code

func (shader *Shader) getUniformLocation(uniform string) int32 {
	location, ok := shader.uniformCache[uniform]

	if !ok {
		uniformName, free := gl.Strs(uniform + "\x00")
		defer free()

		location = gl.GetUniformLocation(shader.program, *uniformName)
		shader.uniformCache[uniform] = location
	}

	return location
}

// Uniform setters

func (shader *Shader) SetColor(uniform string, color structs.Color) {
	location := shader.getUniformLocation(uniform)

	r, g, b, _ := color.Unpack()

	gl.Uniform3f(location, r, g, b)
}

func (shader *Shader) SetMVP(model, view, projection *mgl32.Mat4) {
	shader.SetMatrix(PROGRAM_MODEL_MATRIX_UNIFORM, model)
	shader.SetMatrix(PROGRAM_VIEW_MATRIX_UNIFORM, view)
	shader.SetMatrix(PROGRAM_PROJ_MATRIX_UNIFORM, projection)
}

func (shader *Shader) SetMatrix(uniform string, matrix *mgl32.Mat4) {
	location := shader.getUniformLocation(uniform)

	gl.UniformMatrix4fv(location, 1, false, &matrix[0])
}

func (shader *Shader) SetVec3(uniform string, vec mgl32.Vec3) {
	location := shader.getUniformLocation(uniform)

	gl.Uniform3f(location, vec.X(), vec.Y(), vec.Z())
}

func (shader *Shader) SetTexture(uniform string, texture Texture) {
	location := shader.getUniformLocation(uniform)

	gl.Uniform1i(location, int32(texture.Unit()-gl.TEXTURE0))
}

func (shader *Shader) SetInt(uniform string, val int32) {
	location := shader.getUniformLocation(uniform)

	gl.Uniform1i(location, val)
}

func (shader *Shader) SetFloat(uniform string, val float32) {
	location := shader.getUniformLocation(uniform)

	gl.Uniform1f(location, val)
}

// checks

func checkProgramError(program uint32) error {
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

func checkShaderError(shader uint32) error {
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)

	if status == gl.TRUE {
		return nil
	}

	var infoLog [512]byte

	gl.GetShaderInfoLog(shader, 512, nil, (*uint8)(unsafe.Pointer(&infoLog)))

	errorMessage := fmt.Sprintf("Shader error: %s\n", string(infoLog[:512]))

	return errors.New(errorMessage)
}
