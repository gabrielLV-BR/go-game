package rendering

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

	boundShader     core.Shader
	boundViewMatrix mgl32.Mat4
	boundProjMatrix mgl32.Mat4
}

func (renderer *Renderer) New(window core.Window) error {
	if err := gl.Init(); err != nil {
		return err
	}

	width, height := window.Size()

	gl.Viewport(0, 0, int32(width), int32(height))
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Enable(gl.DEPTH_TEST)

	renderer.shaderMap = make(map[core.MaterialId]core.Shader)

	return nil
}

func (renderer *Renderer) UseCamera(camera *structs.Camera) {
	renderer.boundProjMatrix = camera.ProjectionMatrix()
	renderer.boundViewMatrix = camera.ViewMatrix()
}

func (renderer *Renderer) UseMaterial(material core.Material) {
	shader, ok := renderer.shaderMap[material.Id()]

	if !ok {
		panic("No shader found for material")
	}

	shader.Bind()

	material.Prepare(shader)

	renderer.boundShader = shader
}

func (renderer *Renderer) DrawMesh(meshHandle core.MeshHandle, transform structs.Transform) {
	model := transform.ModelMatrix()
	renderer.boundShader.SetMVP(&model, &renderer.boundViewMatrix, &renderer.boundProjMatrix)

	meshHandle.Bind()

	gl.DrawElements(
		gl.TRIANGLES,
		meshHandle.IndexCount,
		gl.UNSIGNED_INT,
		nil,
	)
}

func (renderer *Renderer) DrawMeshInstanced(meshHandle core.MeshHandle, material core.Material, transforms []mgl32.Mat4) {
	sizeOfMat4 := reflect.TypeOf(mgl32.Mat4{}).Size()
	sizeOfVec4 := reflect.TypeOf(mgl32.Vec4{}).Size()

	// generate data for instanced drawing
	var buffer uint32
	gl.GenBuffers(1, &buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(transforms)*int(sizeOfMat4), gl.Ptr(transforms), gl.STATIC_DRAW)

	meshHandle.Bind()

	// set up data for instance array
	for j := 0; j < 4; j++ {
		i := uint32(j)
		gl.EnableVertexAttribArray(2 + i)
		gl.VertexAttribPointerWithOffset(2+i, 4, gl.FLOAT, false, int32(sizeOfMat4), uintptr(uint32(sizeOfVec4)*i))
		gl.VertexAttribDivisor(2+i, 1)
	}

	//

	renderer.boundShader.SetMVP(&transforms[0], &renderer.boundViewMatrix, &renderer.boundProjMatrix)

	gl.DrawElementsInstanced(
		gl.TRIANGLES,
		meshHandle.IndexCount,
		gl.UNSIGNED_INT,
		nil,
		int32(len(transforms)),
	)
}

func (renderer *Renderer) Resize(width, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}

func (renderer *Renderer) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// TODO put this into config file
func (renderer *Renderer) LoadDefaultMaterials() error {
	defaultMaterials := []string{
		"color",
		"texture",
		"phong",
	}

	const root string = "assets/shaders/"

	for _, material := range defaultMaterials {
		shader := core.Shader{}

		vertexPath := filepath.Join(root, material+".vert.glsl")
		fragPath := filepath.Join(root, material+".frag.glsl")

		vertexSource, err := os.ReadFile(vertexPath)
		if err != nil {
			return err
		}

		fragmentSource, err := os.ReadFile(fragPath)
		if err != nil {
			return err
		}

		if err := shader.LoadStageSource(string(vertexSource), gl.VERTEX_SHADER); err != nil {
			return err
		}

		if err := shader.LoadStageSource(string(fragmentSource), gl.FRAGMENT_SHADER); err != nil {
			return err
		}

		if err := shader.Link(); err != nil {
			return err
		}

		renderer.LoadMaterial(shader, core.MaterialId(material))
	}

	return nil
}

func (renderer *Renderer) LoadMaterial(shader core.Shader, id core.MaterialId) {
	renderer.shaderMap[id] = shader
}
