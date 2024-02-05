package main

import (
	"gabriellv/game/core"
	"gabriellv/game/core/controllers"
	"gabriellv/game/core/materials"
	"gabriellv/game/core/renderer"
	"gabriellv/game/core/systems"
	"gabriellv/game/structs"
	"runtime"

	"github.com/go-gl/mathgl/mgl32"
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

func init() {
	runtime.LockOSThread()
}

func main() {
	window, err := core.NewWindow(WIDTH, HEIGHT, TITLE)
	defer window.Destroy()

	window.SetupInputSystem()

	if err != nil {
		panic(err)
	}

	renderer, err := renderer.NewRenderer(window)

	if err != nil {
		panic(err)
	}

	if err := renderer.LoadDefaultMaterials(); err != nil {
		panic(err)
	}

	vertices := []float32{
		// POSITION        // NORMAL         // UV
		-0.5, -0.5,  0.0,  1.0,  0.0,  0.0,  0.0,  0.0,
		-0.5,  0.5,  0.0,  0.0,  1.0,  0.0,  0.0,  1.0,
		 0.5,  0.5,  0.0,  0.0,  0.0,  1.0,  1.0,  1.0,
		 0.5, -0.5,  0.0,  1.0,  0.0,  1.0,  1.0,  0.0,
	}

	indices := []uint32{
		0, 1, 2,
		2, 3, 0,
	}

	mesh := core.NewMesh(vertices, indices)
	mesh.SetAttributes(
		core.MeshAttributes.Position(),
		core.MeshAttributes.Normal(),
		core.MeshAttributes.UV(),
	)

	phongMaterial := materials.PhongMaterial {
		Color: structs.Colors.White(),
		AmbientLight: core.AmbientLight{
			Color: structs.NewColor(0.7, 0.5, 0.3, 1.0),
		},
		PointLights: []core.PointLight{
			{
				Position: mgl32.Vec3{0.4, 1.0, 0.3},
				Color: structs.NewColor(1.0, 0.3, 0.3, 1.0),
			},
		},
	}

	transform := structs.NewTransform()

	camera := structs.NewCamera()
	camera.UsePerspectiveProjection(80.0, window.AspectRatio(), 0.1, 1000.0)
	camera.SetPosition(mgl32.Vec3{0.0, 0.0, -1.0})

	fpsController := controllers.FPSCameraController{}
	fpsController.Speed = 20.0
	fpsController.MouseSensitivity = 1

	for !window.ShouldClose() {
		fpsController.Update(&camera, 0.016)

		renderer.Clear()
		pass := renderer.BeginDraw(&camera)

		pass.DrawMesh(mesh, transform, &phongMaterial)

		pass.EndDraw()

		systems.InputSystem.Update()

		window.PollAndSwap()
	}
}
