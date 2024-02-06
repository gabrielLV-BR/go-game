package renderer

import (
	"gabriellv/game/core"
	"gabriellv/game/structs"
	"os"
	"path/filepath"
	"reflect"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Renderer struct {
	shaderMap map[core.MaterialId]core.Shader
}

type RenderPass struct {
	renderer         *Renderer
	projectionMatrix mgl32.Mat4
	viewMatrix       mgl32.Mat4
}

func NewRenderer(window core.Window) (Renderer, error) {
	renderer := Renderer{}

	if err := gl.Init(); err != nil {
		return renderer, err
	}

	width, height := window.Size()

	gl.Viewport(0, 0, int32(width), int32(height))
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Enable(gl.DEPTH_TEST)

	renderer.shaderMap = make(map[core.MaterialId]core.Shader)

	return renderer, nil
}

//TODO put this into config file
func (renderer *Renderer) LoadDefaultMaterials() error {
	defaultMaterials := []string{
		"color",
		"texture",
		"phong",
	}

	const root string = "assets/shaders/"

	for _, material := range defaultMaterials {
		newShader := core.NewShader()

		vertexPath := filepath.Join(root, material+".vert.glsl")
		fragPath := filepath.Join(root, material+".frag.glsl")

		vertexSource, err := os.ReadFile(vertexPath)
		if err != nil { return err }

		fragmentSource, err := os.ReadFile(fragPath)
		if err != nil { return err }

		if err := newShader.LoadStageSource(string(vertexSource), gl.VERTEX_SHADER) ; err != nil {
			return err
		}

		if err := newShader.LoadStageSource(string(fragmentSource), gl.FRAGMENT_SHADER) ; err != nil {
			return err
		}

		if err := newShader.Link(); err != nil {
			return err
		}

		renderer.shaderMap[core.MaterialId(material)] = newShader
	}

	return nil
}

func (renderer *Renderer) GetShaderFor(material core.Material) core.Shader {
	program, ok := renderer.shaderMap[material.Id()]

	if !ok {
		panic("No shader for material")
	}

	return program
}

func (renderer *Renderer) BeginDraw(camera *structs.Camera) RenderPass {
	return RenderPass{
		renderer:         renderer,
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

func (pass *RenderPass) DrawMesh(mesh core.Mesh, transform structs.Transform, material core.Material) {
	// get program from material
	shader := pass.renderer.GetShaderFor(material)
	shader.Bind()

	material.Prepare(shader)

	model := transform.ModelMatrix()

	shader.SetMVP(&model, &pass.viewMatrix, &pass.projectionMatrix)

	mesh.Bind()

	gl.DrawElements(
		gl.TRIANGLES,
		int32(mesh.IndexCount()),
		gl.UNSIGNED_INT,
		nil,
	)

	shader.Unbind()
	mesh.Unbind()
}

func (pass *RenderPass) DrawMeshInstanced(mesh core.Mesh, material core.Material, transforms []mgl32.Mat4) {
	sizeOfMat4 := reflect.TypeOf(mgl32.Mat4{}).Size()
	sizeOfVec4 := reflect.TypeOf(mgl32.Vec4{}).Size()

	// generate data for instanced drawing
	var buffer uint32
	gl.GenBuffers(1, &buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(transforms)*int(sizeOfMat4), gl.Ptr(transforms), gl.STATIC_DRAW)

	mesh.Bind()

	// set up data for instance array
	for j := 0; j < 4; j++ {
		i := uint32(j)
		gl.EnableVertexAttribArray(2 + i)
		gl.VertexAttribPointerWithOffset(2+i, 4, gl.FLOAT, false, int32(sizeOfMat4), uintptr(uint32(sizeOfVec4)*i))
		gl.VertexAttribDivisor(2+i, 1)
	}

	//

	shader := pass.renderer.GetShaderFor(material)
	shader.Bind()

	material.Prepare(shader)

	shader.SetMVP(&transforms[0], &pass.viewMatrix, &pass.projectionMatrix)

	gl.DrawElementsInstanced(
		gl.TRIANGLES,
		int32(mesh.IndexCount()),
		gl.UNSIGNED_INT,
		nil,
		int32(len(transforms)),
	)

	shader.Unbind()
	mesh.Unbind()
}
