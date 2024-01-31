package main

import (
	"gabriellv/game/core"
	"gabriellv/game/structs/color"

	"github.com/go-gl/gl/v4.1-core/gl"
)

/*
	GO GAME
	-------
	Check README.md for more information
*/

const (
	WIDTH  = 500
	HEIGHT = 500
	TITLE  = "Window"
)

func main() {
	window, err := core.NewWindow(WIDTH, HEIGHT, TITLE)
	defer window.Destroy()

	if err != nil {
		panic(err)
	}

	renderer, err := core.NewRenderer(&window)

	if err != nil {
		panic(err)
	}

	if err := renderer.LoadDefaultMaterials() ; err != nil {
		panic(err)
	}

	//lastTime := core.GetTime()
	//delta := 0.0

	vertices := []float32{
		-0.2, -0.5, 0.0,
		-0.5, 0.5, 0.0,
		0.5, 0.5, 0.0,
		0.5, -0.5, 0.0,
	}

	indices := []uint32{
		0, 1, 2,
		2, 3, 0,
	}

	mesh := core.NewMesh(vertices, indices)

	// render renderer mf
	material := core.Material {
		Color: color.Color{R: 1.0, G: 1.0, B: 1.0, A: 1.0},
	}

	gl.ClearColor(1.0, 0.8, 0.6, 1.0)

	for !window.ShouldClose() {
		//delta = core.GetTime() - lastTime
		gl.Clear(gl.COLOR_BUFFER_BIT)

		renderer.DrawMesh(mesh, material)

		window.PollAndSwap()

		//lastTime = glfw.GetTime()
	}
}
