package main

import (
	"gabriellv/game/core"
	"gabriellv/game/structs/color"
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

	window.SetupInputSystem()

	if err != nil {
		panic(err)
	}

	renderer, err := core.NewRenderer(window)

	if err != nil {
		panic(err)
	}

	if err := renderer.LoadDefaultMaterials(); err != nil {
		panic(err)
	}

	vertices := []float32{
		-0.5, -0.5, 0.0, 0.0, 0.0,
		-0.5, 0.5, 0.0, 0.0, 1.0,
		0.5, 0.5, 0.0, 1.0, 1.0,
		0.5, -0.5, 0.0, 1.0, 0.0,
	}

	indices := []uint32{
		0, 1, 2,
		2, 3, 0,
	}

	mesh := core.NewMesh(vertices, indices)
	mesh.SetAttributes(
		core.MeshAttributes.Position(),
		core.MeshAttributes.UV(),
	)

	diffuse, err := core.LoadTexture("assets/textures/smile.png")

	if err != nil {
		panic(err)
	}

	// render renderer mf
	material := core.Material{
		Color:    color.Colors.White(),
		Textures: []core.Texture{diffuse},
	}

	for !window.ShouldClose() {
		renderer.Clear()

		renderer.DrawMesh(mesh, material)

		window.PollAndSwap()
	}
}
