package main

import (
	"gabriellv/game/controllers"
	"gabriellv/game/core"
	"gabriellv/game/core/materials"
	"gabriellv/game/core/rendering"
	"gabriellv/game/core/systems"
	"gabriellv/game/gamedata"
	"gabriellv/game/gamedata/entities"
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

	diffuse, err := core.LoadTexture("assets/textures/smile.png")
	if err != nil {
		panic(err)
	}

	obj, err := loaders.LoadFromFile("assets/models/cone.obj")
	if err != nil {
		panic(err)
	}

	gameState := gamedata.State{}
	gameState.Camera = structs.NewCamera()
	gameState.Camera.UsePerspectiveProjection(80.0, window.AspectRatio(), 0.1, 1000.0)
	gameState.Camera.SetPosition(mgl32.Vec3{0.0, 0.0, -1.0})

	{ // Player creation
		//TODO load this kind of stuff from a file

		player := gamedata.Player{}
		player.Controller = controllers.FPSCameraController{
			Speed:            10.0,
			MouseSensitivity: 1.0,
		}

		gameState.AddEntity(&player)
	}

	{ // Spining object entity
		ent := entities.SpinningEntity{
			Velocity: 10.0,
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

		model := core.Model{
			Mesh:      mesh,
			Material:  &material,
			Transform: transform,
		}

		ent.ModelId = gameState.AddModel(model)
		gameState.AddEntity(&ent)
	}

	for !window.ShouldClose() {

		// Update fase
		for _, ent := range gameState.Entities {
			ent.Update(&gameState, 0.16)
		}

		// Rendering fase
		renderer.Clear()
		pass := renderer.BeginDraw(&gameState.Camera)

		for _, model := range gameState.Models {
			pass.DrawMesh(model.Mesh, model.Transform, model.Material)
		}

		pass.EndDraw()

		//TODO figure out a way to make this suck less
		systems.InputSystem.Update()

		window.PollAndSwap()
	}
}
