package core

import (
	"errors"
	"fmt"
	"os"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Shader struct {
	id    uint32
	stage uint32
}

func NewShader(source string, stage uint32) (Shader, error) {
	shader := gl.CreateShader(stage)

	if stage != gl.VERTEX_SHADER && stage != gl.FRAGMENT_SHADER {
		return Shader{}, errors.New("Invalid shader stage")
	}

	strings, free := gl.Strs(source + "\x00")
	defer free()

	gl.ShaderSource(shader, 1, strings, nil)
	gl.CompileShader(shader)

	err := CheckShaderError(shader)

	if err != nil {
		return Shader{}, err
	}

	return Shader{id: shader, stage: stage}, nil
}

func LoadShader(path string, stage uint32) (Shader, error) {
	buffer, err := os.ReadFile(path)

	if err != nil {
		return Shader{}, err
	}

	return NewShader(string(buffer), stage)
}

func CheckShaderError(shader uint32) error {
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
