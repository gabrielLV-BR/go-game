package main

import (
	"gabriellv/game/core"
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

	err = renderer.LoadDefaultMaterials()

	if err != nil {
		panic(err)
	}

	//lastTime := core.GetTime()
	//delta := 0.0

	vertices := []float32 {
		-0.5, -0.5, 0.0,
		-0.5, 0.5, 0.0,
		0.5, 0.5, 0.0,
		0.5, -0.5, 0.0,
	}

	indices := []uint32 {
		0, 1, 2,
		2, 3, 0,
	}

	mesh := core.NewMesh(vertices, indices)
	mesh.Bind()

	// render renderer mf

	for !window.ShouldClose() {
		//delta = core.GetTime() - lastTime

		renderer.DrawMesh(mesh, core.Material {})

		window.PollAndSwap()

		//lastTime = glfw.GetTime()
	}
}
