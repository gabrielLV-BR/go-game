package core

import (
	"gabriellv/game/structs"
	"path/filepath"
	"reflect"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type RenderFeatures uint16

const (
	RENDER_FEATURE_INSTANCED RenderFeatures = 1 << iota
)

type Renderer struct {
	programMap map[uint32]Program
}

type RenderPass struct {
	renderer         *Renderer
	features 		 RenderFeatures
	projectionMatrix mgl32.Mat4
	viewMatrix       mgl32.Mat4
}

func NewRenderer(window Window) (Renderer, error) {
	renderer := Renderer{}

	if err := gl.Init(); err != nil {
		return renderer, err
	}

	gl.Viewport(0, 0, int32(window.width), int32(window.height))
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	gl.Enable(gl.DEPTH_TEST)

	renderer.programMap = make(map[uint32]Program)

	return renderer, nil
}

func (renderer *Renderer) GetProgramForMaterial(material Material, features RenderFeatures) Program {
	id := programId(material.Id(), features)
	program, ok := renderer.programMap[id]

	if !ok {
		panic("Program not found for material")
	}

	return program
}

func (renderer *Renderer) LoadDefaultMaterials() error {
	//TODO put this into config file
	defaultMaterials := []string{
		"color",
		"texture",
		"texture_instanced",
	}

	materialIds := []uint32{
		0, 1, programId(1, RENDER_FEATURE_INSTANCED),
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

func (renderer *Renderer) BeginDraw(camera *structs.Camera, features RenderFeatures) RenderPass {
	return RenderPass{
		renderer:         renderer,
		features:         features,
		projectionMatrix: camera.GetProjectionMatrix(),
		viewMatrix:       camera.GetViewMatrix(),
	}
}

func (pass *RenderPass) EndDraw() {
	pass.renderer = nil
}

func (renderer *Renderer) Resize(width, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}

func (renderer *Renderer) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (pass *RenderPass) DrawMesh(mesh Mesh, transform structs.Transform, material Material) {
	program := pass.renderer.GetProgramForMaterial(material, RenderFeatures(pass.features))

	program.Bind()
	mesh.Bind()

	model := transform.GetModelMatrix()

	program.SetMaterial(material)
	program.SetMVP(&model, &pass.viewMatrix, &pass.projectionMatrix)

	gl.DrawElements(
		gl.TRIANGLES,
		int32(len(mesh.indices)),
		gl.UNSIGNED_INT,
		nil,
	)

	program.Unbind()
	mesh.Unbind()
}

func (pass *RenderPass) DrawMeshInstanced(mesh Mesh, material Material, transforms []mgl32.Mat4) {

	sizeOfMat4 := reflect.TypeOf(mgl32.Mat4{}).Size()
	sizeOfVec4 := reflect.TypeOf(mgl32.Vec4{}).Size()

	// generate data for instanced drawing	
	var buffer uint32
	gl.GenBuffers(1, &buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(transforms) * int(sizeOfMat4), gl.Ptr(transforms), gl.STATIC_DRAW)

	mesh.Bind()

	// set up data for instance array
	for j := 0; j < 4; j++ {
		i := uint32(j)
		gl.EnableVertexAttribArray(2 + i)
		gl.VertexAttribPointerWithOffset(2 + i, 4, gl.FLOAT, false, int32(sizeOfMat4), uintptr(uint32(sizeOfVec4) * i))
		gl.VertexAttribDivisor(2 + i, 1)
	}

	program := pass.renderer.GetProgramForMaterial(material, RenderFeatures(pass.features))
	program.Bind()

	program.SetMaterial(material)
	program.SetVP(&pass.viewMatrix, &pass.projectionMatrix)

	gl.DrawElementsInstanced(
		gl.TRIANGLES,
		int32(len(mesh.indices)),
		gl.UNSIGNED_INT,
		nil,
		int32(len(transforms)),
	)

	program.Unbind()
	mesh.Unbind()
}

//

func programId(materialId uint32, features RenderFeatures) uint32 {
	return (uint32(features) << 16) | materialId
}
