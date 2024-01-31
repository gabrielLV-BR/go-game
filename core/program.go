package core

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Program struct {
	id uint32
	uniformCache map[string]int32
}

func (program *Program) Bind() {
	gl.UseProgram(program.id)
}

func (program *Program) Unbind() {
	gl.UseProgram(0)
}

//

func NewProgram(vertex, fragment Shader) (Program, error) {
	program := Program {}

	id := gl.CreateProgram()
	gl.AttachShader(id, vertex.id)
	gl.AttachShader(id, fragment.id)

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

func (program *Program) SetMaterial(material *Material) {
	gl.UseProgram(program.id)

	textureLocation := program.getUniformLocation(MATERIAL_TEXTURE_UNIFORM)

	r, g, b, _ := material.Color.Unpack()

	gl.Uniform3f(textureLocation, r, g, b)

	//TODO also set textures
}

// set uniforms

func (program* Program) getUniformLocation(uniform string) int32 {

	location, ok := program.uniformCache[uniform]

	if !ok {
		location = gl.GetUniformLocation(program.id, gl.Str(uniform + "\x00"))

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
