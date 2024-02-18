package core

import (
	"reflect"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type MeshHandle struct {
	VAO        uint32
	IndexCount int32
}

type Mesh struct {
	Vertices []float32
	Indices  []uint32
}

type MeshAttribute struct {
	count int32
	xsize int32
	xtype uint32
}

func (mesh *Mesh) Bind(attributes ...MeshAttribute) MeshHandle {
	var vao uint32
	var vbo uint32
	var ebo uint32

	vertices_size := len(mesh.Vertices) * int(reflect.TypeOf(mesh.Vertices).Elem().Size())
	indices_size := len(mesh.Indices) * int(reflect.TypeOf(mesh.Indices).Elem().Size())

	// vao
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, vertices_size, gl.Ptr(mesh.Vertices), gl.STATIC_DRAW)

	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, indices_size, gl.Ptr(mesh.Indices), gl.STATIC_DRAW)

	setAttributes(attributes...)

	gl.BindVertexArray(0)

	return MeshHandle{
		VAO:        vao,
		IndexCount: int32(len(mesh.Indices)),
	}
}

func (handle MeshHandle) Bind() {
	gl.BindVertexArray(uint32(handle.VAO))
}

func (handle MeshHandle) Unbind() {
	gl.BindVertexArray(0)
}

func setAttributes(attributes ...MeshAttribute) {
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
