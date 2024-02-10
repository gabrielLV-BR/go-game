package core

import (
	"reflect"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Mesh struct {
	vao      uint32
	vertices []float32
	indices  []uint32
}

type MeshAttribute struct {
	count int32
	xsize int32
	xtype uint32
}

func NewMesh(vertices []float32, indices []uint32) Mesh {
	var vao uint32
	var vbo uint32
	var ebo uint32

	vertices_size := len(vertices) * int(reflect.TypeOf(vertices).Elem().Size())
	indices_size := len(indices) * int(reflect.TypeOf(indices).Elem().Size())

	// vao
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, vertices_size, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, indices_size, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.BindVertexArray(0)

	return Mesh{
		vao:      vao,
		vertices: vertices,
		indices:  indices,
	}
}

func (mesh *Mesh) SetAttributes(attributes ...MeshAttribute) {
	mesh.Bind()

	var stride int32

	for _, attrib := range attributes {
		stride += attrib.Stride()
	}

	var offset uint32
	for i, attribute := range attributes {
		gl.VertexAttribPointerWithOffset(
			uint32(i),
			attribute.count,
			attribute.xtype,
			false,
			stride,
			uintptr(offset),
		)
		gl.EnableVertexAttribArray(uint32(i))
		offset += (uint32)(attribute.Stride())
	}

	mesh.Unbind()
}

func (mesh *Mesh) IndexCount() int {
	return len(mesh.indices)
}

//

func (mesh Mesh) Bind() {
	gl.BindVertexArray(mesh.vao)
}

func (mesh Mesh) Unbind() {
	gl.BindVertexArray(0)
}

//

var MeshAttributes MeshAttribute

func (m MeshAttribute) Position() MeshAttribute {
	return MeshAttribute{
		count: 3,
		xsize: 4,
		xtype: gl.FLOAT,
	}
}

func (m MeshAttribute) Normal() MeshAttribute {
	return m.Position()
}

func (m MeshAttribute) UV() MeshAttribute {
	return MeshAttribute{
		count: 2,
		xsize: 4,
		xtype: gl.FLOAT,
	}
}

func (m MeshAttribute) Stride() int32 {
	return m.count * m.xsize
}
