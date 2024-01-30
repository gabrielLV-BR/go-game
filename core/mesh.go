package core

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Mesh struct {
	vao uint32
	vertices []float32
	indices []uint32
}

func NewMesh(vertices []float32, indices []uint32) Mesh {
	var vao uint32
	var vbo uint32
	var ebo uint32

	vertices_size := len(vertices) * 4
	indices_size := len(indices) * 4

	// vao
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)	
	gl.BufferData(gl.ARRAY_BUFFER, vertices_size, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, indices_size, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, gl.Ptr(0))
	gl.EnableVertexAttribArray(0)

	return Mesh {
		vao: vao,
		vertices: vertices,
		indices: indices,
	}
}

//

func (mesh* Mesh) Bind() {
	gl.BindVertexArray(mesh.vao)
}
