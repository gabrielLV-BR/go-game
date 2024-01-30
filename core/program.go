package core

import (
	"errors"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Program struct {
	id uint32
}

func NewProgram(vertex, fragment Shader) (Program, error) {
	program := gl.CreateProgram()
	gl.AttachShader(program, vertex.id)
	gl.AttachShader(program, fragment.id)

	gl.LinkProgram(program)

	err := CheckProgramError(program)

	if err != nil {
		return Program{}, nil
	}

	return Program {}, nil
}

// Shader
//

func CheckProgramError(program uint32) error {
	var status int32
	gl.GetShaderiv(program, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		infolog := make([]uint8, 128)
		gl.GetShaderInfoLog(program, 128, nil, (*uint8)(gl.Ptr(infolog)))

		return errors.New(string(infolog))
	}
	
	return nil
}
