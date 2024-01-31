package core

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Program struct {
	id uint32
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

	return program, nil
}

// uniform interaction code

func (program *Program) SetMaterial(material *Material) {
	gl.UseProgram(program.id)

	textureLocation := getUniformLocation(program.id, MATERIAL_TEXTURE_UNIFORM)

	r, g, b, _ := material.Color.Unpack()

	gl.Uniform3f(textureLocation, r, g, b)

	//TODO also set textures
}

// set uniforms

func getUniformLocation(programHandle uint32, uniformName string) int32 {
	return gl.GetUniformLocation(programHandle, gl.Str(uniformName+"\x00"))
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
