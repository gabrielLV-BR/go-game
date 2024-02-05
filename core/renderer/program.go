package renderer

import (
	"errors"
	"fmt"
	"gabriellv/game/core"
	"gabriellv/game/structs"
	"reflect"
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

	err := checkProgramError(id)

	if err != nil {
		return Program{}, err
	}

	program.id = id
	program.uniformCache = make(map[string]int32)

	return program, nil
}

// uniform interaction code
func (program *Program) SetMaterial(material core.Material) {
	uniforms := material.Uniforms()

	for _, uniform := range uniforms {
		name, value := uniform.Name, uniform.Value

		if err := program.trySetUniform(name, value); err != nil {
			panic(err)
		}
	}
}

func (program *Program) trySetUniform(name string, value interface{}) error {
	switch value := value.(type) {
	case bool:
		if value {
			program.SetInt(name, 1)
		} else { 
			program.SetInt(name, 0)
		}
	case int32:
		program.SetInt(name, value)
	case structs.Color:
		program.SetColor(name, value)
	case core.Texture:
		program.SetTexture(name, value)
	case core.AmbientLight:
		program.SetAmbientLight(name, value)
	case core.PointLight:
		program.SetPointLight(name, value)
	case []interface{}:
		for i, u := range value {
			program.trySetUniform(fmt.Sprintf("%s[%d]", name, i), u)
		}
	default:
		return errors.New(fmt.Sprintf("Unknown uniform %s of type %s", name, reflect.TypeOf(value).Name()))
	}

	return nil
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

// Uniform setters

func (program *Program) SetColor(uniform string, color structs.Color) {
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

func (program *Program) SetVec3(uniform string, vec mgl32.Vec3) {
	location := program.getUniformLocation(uniform)

	gl.Uniform3f(location, vec.X(), vec.Y(), vec.Z())
}

func (program *Program) SetAmbientLight(uniform string, ambientLight core.AmbientLight) {
	colorUniform := uniform + ".color"
	// intensityUniform := uniform + ".intensity"

	program.SetColor(colorUniform, ambientLight.Color)
	// program.SetFloat(intensityUniform, ambientLight.Intensity)
}

func (program *Program) SetPointLight(uniform string, pointLight core.PointLight) {
	positionUniform := uniform + ".position"
	colorUniform := uniform + ".color"

	program.SetVec3(positionUniform, pointLight.Position)
	program.SetColor(colorUniform, pointLight.Color)
}

func (program *Program) SetTexture(uniform string, texture core.Texture) {
	location := program.getUniformLocation(uniform)

	gl.Uniform1i(location, int32(texture.Unit()-gl.TEXTURE0))
}

func (program *Program) SetInt(uniform string, val int32) {
	location := program.getUniformLocation(uniform)

	gl.Uniform1i(location, val)
}

func (program *Program) SetFloat(uniform string, val float32) {
	location := program.getUniformLocation(uniform)

	gl.Uniform1f(location, val)
}
