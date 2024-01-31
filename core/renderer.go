package core

import (
	"path/filepath"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Renderer struct {
	programMap map[uint32]Program
}

func NewRenderer(window *Window) (Renderer, error) {
	renderer := Renderer{}

	if err := gl.Init() ; err != nil {
		return renderer, err
	}

	gl.Viewport(0, 0, int32(window.width), int32(window.height))

	renderer.programMap = make(map[uint32]Program)

	return renderer, nil
}

func (renderer *Renderer) GetProgramForMaterial(material Material) *Program {
	program, ok := renderer.programMap[material.Id()]

	if !ok {
		return nil
	}

	return &program
}

func (renderer *Renderer) LoadDefaultMaterials() error {
	defaultMaterials := []string{
		"color",
	}

	materialIds := []uint32{
		0,
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

func (renderer *Renderer) DrawMesh(mesh Mesh, material Material) {
	program := renderer.GetProgramForMaterial(material)

	if program == nil {
		panic("No program for given Material")
	}

	gl.UseProgram(program.id)
	gl.BindVertexArray(mesh.vao)

	program.SetMaterial(&material)

	gl.DrawElements(
		gl.TRIANGLES,
		int32(len(mesh.indices)),
		gl.UNSIGNED_INT,
		unsafe.Pointer(nil),
	)

	gl.UseProgram(0)
	gl.BindVertexArray(0)
}
