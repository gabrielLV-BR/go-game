package main

import (
	"gabriellv/game/controllers"
	"gabriellv/game/core"
	"gabriellv/game/structs"
	"gabriellv/game/structs/color"
	"gabriellv/game/systems"
	"math/rand"
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

	// transform := structs.NewTransform()

	diffuse, err := core.LoadTexture("assets/textures/smile.png")

	if err != nil {
		panic(err)
	}

	// render renderer mf
	material := core.Material{
		Color:    color.Colors.White(),
		Textures: []core.Texture{diffuse},
	}

	transforms := make([]mgl32.Mat4, 100)

	for i := range transforms {
		x := rand.Float32()
		y := rand.Float32()
		z := rand.Float32()

		transform := structs.Transform{}

		transform.Position = mgl32.Vec3{x * 10.0, y * 10.0, z * 10.0}
		transform.Scale = mgl32.Vec3{1, 1, 1}
		transform.Rotation = mgl32.QuatIdent()
		// transform.Rotation = mgl32.AnglesToQuat(x, y, z, mgl32.XYZ)

		transforms[i] = transform.GetModelMatrix()
	}

	// transform := transforms[0]

	camera := structs.NewCamera()
	camera.UsePerspectiveProjection(80.0, float32(window.GetAspectRatio()), 0.1, 1000.0)
	camera.SetPosition(mgl32.Vec3{0.0, 1.0, -1.0})

	fpsController := controllers.FPSCameraController{}
	fpsController.Speed = 20.0
	fpsController.MouseSensitivity = 1

	for !window.ShouldClose() {
		fpsController.Update(&camera, 0.016)

		renderer.Clear()
		pass := renderer.BeginDraw(&camera, core.RENDER_FEATURE_INSTANCED)

		// pass.DrawMesh(mesh, transform, material)

		pass.DrawMeshInstanced(mesh, material, transforms)

		pass.EndDraw()

		systems.InputSystem.Update()

		window.PollAndSwap()
	}
}
