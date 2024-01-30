package core

import (
	"errors"
	"os"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Shader struct {
	id uint32
	stage uint32
}

func NewShader(source string, stage uint32) (Shader, error) {
	shader := gl.CreateShader(stage)

	strings, free := gl.Strs(source)
	defer free()

	gl.ShaderSource(shader, 1, strings, (*int32)(gl.Ptr(nil)))
	gl.CompileShader(shader)

	err := CheckShaderError(shader)

	if err != nil {
		panic(err)
	}

	return Shader { id: shader, stage: stage }, nil
}

func LoadShader(path string, stage uint32) (Shader, error) {

	buffer, err := os.ReadFile(path)

	if err != nil {
		return Shader {}, err
	}

	return NewShader(string(buffer), stage)
}


func CheckShaderError(shader uint32) error {
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		logLen := int32(0)
		gl.GetShaderInfoLog(shader, 0, &logLen, nil)
		logLen += 1

		log := gl.Str(strings.Repeat("\x00", int(logLen)))

		gl.GetShaderInfoLog(shader, logLen, nil, log)

		return errors.New(gl.GoStr(log))
	}
	
	return nil
}
