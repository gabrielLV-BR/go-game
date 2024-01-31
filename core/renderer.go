package core

import (
	"path/filepath"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Renderer struct {
	programMap map[uint32]Program
}

func NewRenderer(window *Window) (Renderer, error) {
	renderer := Renderer{}

	if err := gl.Init(); err != nil {
		return renderer, err
	}

	gl.Viewport(0, 0, int32(window.width), int32(window.height))
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	renderer.programMap = make(map[uint32]Program)

	return renderer, nil
}

func (renderer *Renderer) ProgramMap() map[uint32]Program {
	return renderer.programMap
}

func (renderer *Renderer) GetProgramForMaterial(material *Material) Program {
	program, ok := renderer.programMap[material.Id()]

	if !ok {
		panic("Program not found for material")
	}

	return program
}

func (renderer *Renderer) LoadDefaultMaterials() error {
	defaultMaterials := []string{
		"color",
		"texture",
	}

	materialIds := []uint32{
		0, 1,
	}

	if len(defaultMaterials) != len(materialIds) {
		panic("There must be a 1:1 mapping of Shader names and Id's")
	}

	const root string = "assets/shaders/"

	for i, material := range defaultMaterials {
		vertexPath := filepath.Join(root, material+".vert.glsl")
		fragPath := filepath.Join(root, material+".frag.glsl")

		vertexShader, err := LoadShader(vertexPath, gl.VERTEX_SHADER)

		if err != nil {
			return err
		}

		fragShader, err := LoadShader(fragPath, gl.FRAGMENT_SHADER)

		if err != nil {
			return err
		}

		program, err := NewProgram(vertexShader, fragShader)

		if err != nil {
			return err
		}

		gl.DeleteShader(vertexShader.id)
		gl.DeleteShader(fragShader.id)

		renderer.programMap[materialIds[i]] = program
	}

	return nil
}

func (renderer *Renderer) Resize(width, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}

func (renderer *Renderer) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (renderer *Renderer) DrawMesh(mesh *Mesh, material *Material) {
	program := renderer.GetProgramForMaterial(material)

	program.Bind()
	mesh.Bind()

	program.SetMaterial(material)
	mesh.Draw()

	program.Unbind()
	mesh.Unbind()
}
