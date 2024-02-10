package rendering

import (
	"gabriellv/game/core"
	"gabriellv/game/structs"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type ShapeRenderer struct {
	quadShape core.Mesh

	quadShader core.Shader
}

func NewShapeRenderer() (ShapeRenderer, error) {
	shapeRenderer := ShapeRenderer{}

	quadMesh := newQuadMesh()

	quadShader, err := loadShader("assets/shaders/shapes/quad")
	if err != nil {
		return shapeRenderer, err
	}

	shapeRenderer.quadShape = quadMesh
	shapeRenderer.quadShader = quadShader

	return shapeRenderer, nil
}

func (shape *ShapeRenderer) DrawQuad(pass *RenderPass, position, size mgl32.Vec3, color structs.Color) {
	model := mgl32.Translate3D(position.X(), position.Y(), position.Z())
	// model = model.Mul4(mgl32.Scale3D(size.X(), size.Y(), size.Z()))

	shape.quadShape.Bind()

	shape.quadShader.Bind()
	shape.quadShader.SetColor(core.MATERIAL_COLOR_UNIFORM, color)
	shape.quadShader.SetMVP(&model, &pass.viewMatrix, &pass.projectionMatrix)

	gl.DrawElements(gl.TRIANGLES, int32(shape.quadShape.IndexCount()), gl.UNSIGNED_INT, nil)

	shape.quadShader.Unbind()
	shape.quadShape.Unbind()
}

func loadShader(path string) (core.Shader, error) {
	shader := core.NewShader()

	vertShaderSource, err := os.ReadFile(path + ".vert.glsl")
	if err != nil {
		return shader, err
	}

	fragShaderSource, err := os.ReadFile(path + ".frag.glsl")
	if err != nil {
		return shader, err
	}

	shader.LoadStageSource(string(vertShaderSource), gl.VERTEX_SHADER)
	shader.LoadStageSource(string(fragShaderSource), gl.FRAGMENT_SHADER)

	err = shader.Link()

	return shader, err
}

func newQuadMesh() core.Mesh {
	vertices := []float32{
		-1.0, -1.0, 0.0, // 0.0, 0.0,
		-1.0, 1.0, 0.0, // 0.0, 1.0,
		1.0, 1.0, 0.0, // 1.0, 1.0,
		1.0, -1.0, 0.0, // 0.0, 1.0,
	}

	indices := []uint32{
		0, 1, 2,
		2, 3, 0,
	}

	mesh := core.NewMesh(vertices, indices)

	mesh.SetAttributes(
		core.MeshAttributes.Position(),
		// core.MeshAttributes.UV()
	)

	return mesh
}
