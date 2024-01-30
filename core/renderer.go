package core

import (
	"gabriellv/game/structs"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Renderer struct {
	programMap map[uint32]Program
}

func NewRenderer(window *Window) (Renderer, error) {
	renderer := Renderer{}

	err := gl.Init()

	if err != nil {
		return renderer, err
	}

	gl.Viewport(0, 0, int32(window.width), int32(window.height))

	return renderer, nil
}

func (renderer *Renderer) GetProgramForMaterial(material Material) *Program {
	program := renderer.programMap[material.Id()]

	return &program
}

func (renderer *Renderer) LoadDefaultMaterials() error {
	defaultMaterials := []string {
		"color",
	}

	materialIds := []uint32 {
		0,
	}

	const root string = "assets/shaders/"

	for i, material := range defaultMaterials {

		vertexPath := root + material + ".vert.glsl"
		fragPath := root + material + ".frag.glsl"
		
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
			panic(err)
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

func (renderer* Renderer) DrawMesh(mesh Mesh, material Material) {
	program := renderer.GetProgramForMaterial(material)

	gl.UseProgram(program.id)
	gl.BindVertexArray(mesh.vao)

	gl.DrawElements(
		gl.TRIANGLES,
		int32(len(mesh.indices)),
		gl.UNSIGNED_INT,
		gl.Ptr(nil),
	)
}

func (renderer* Renderer) DrawQuad(x, y, width, height float32, color structs.Color) {
}
