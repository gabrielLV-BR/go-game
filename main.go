package main

import (
	"gabriellv/game/core"
	"gabriellv/game/core/controllers"
	"gabriellv/game/core/materials"
	"gabriellv/game/core/rendering"
	"gabriellv/game/core/systems"
	"gabriellv/game/loaders"
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
	if err != nil {
		panic(err)
	}

	defer window.Destroy()

	window.SetupInputSystem()

	renderer, err := rendering.NewRenderer(window)
	if err != nil {
		panic(err)
	}

	err = renderer.LoadDefaultMaterials()
	if err != nil {
		panic(err)
	}

	shapeRenderer, err := rendering.NewShapeRenderer()
	if err != nil {
		panic(err)
	}

	diffuse, err := core.LoadTexture("assets/textures/smile.png")
	if err != nil {
		panic(err)
	}

	obj, err := loaders.LoadObj("assets/models/cone.obj")
	if err != nil {
		panic(err)
	}

	mesh := core.NewMesh(obj.Vertices, obj.Indices)
	mesh.SetAttributes(
		core.MeshAttributes.Position(),
		core.MeshAttributes.Normal(),
		core.MeshAttributes.UV(),
	)

	material := materials.TextureMaterial{
		Color:   structs.Colors.White(),
		Texture: diffuse,
	}

	transform := structs.NewTransform()

	camera := structs.NewCamera()
	camera.UsePerspectiveProjection(80.0, window.AspectRatio(), 0.1, 1000.0)
	camera.SetPosition(mgl32.Vec3{0.0, 0.0, -1.0})

	width, height := window.Size()
	ortho := structs.NewCamera()
	ortho.UseOrthogonalProjection(float32(width), float32(height), 0.0, 1000.0)
	ortho.SetPosition(mgl32.Vec3{0.0, 0.0, -1.0})
	ortho.LookAt(mgl32.Vec3{}, mgl32.Vec3{0, 1, 0})

	fpsController := controllers.FPSCameraController{}
	fpsController.Speed = 20.0
	fpsController.MouseSensitivity = 1

	for !window.ShouldClose() {
		fpsController.Update(&camera, 0.016)

		renderer.Clear()
		pass := renderer.BeginDraw(&camera)

		pass.DrawMesh(mesh, transform, &material)

		pass.EndDraw()

		pass = renderer.BeginDraw(&camera)

		shapeRenderer.DrawQuad(
			&pass,
			mgl32.Vec3{1.5, 1.2, 0.0},
			mgl32.Vec3{1.0, 1.0, 1.0},
			structs.Colors.White(),
		)

		pass.EndDraw()

		systems.InputSystem.Update()

		window.PollAndSwap()
	}
}
