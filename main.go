package main

import (
	"gabriellv/game/algorithms/procedural"
	"gabriellv/game/controllers"
	"gabriellv/game/core"
	"gabriellv/game/core/materials"
	"gabriellv/game/core/rendering"
	"gabriellv/game/core/systems"
	"gabriellv/game/gamedata"
	"gabriellv/game/gamedata/entities"
	"gabriellv/game/loaders"
	"gabriellv/game/physics"
	"gabriellv/game/physics/shapes"
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
	gameState.Camera = &structs.Camera{}
	gameState.Camera.New()
	gameState.Camera.UsePerspectiveProjection(80.0, window.AspectRatio(), 0.1, 1000.0)
	gameState.Camera.SetPosition(mgl32.Vec3{0.0, 0.0, -1.0})

	{ // Player creation
		//TODO load this kind of stuff from a file

		player := entities.Player{}
		player.Controller = controllers.FPSCameraController{
			Speed:            10.0,
			MouseSensitivity: 1.0,
		}

		body := physics.PhysicsBody{
			Transform: structs.NewTransform(),
			Shape: &shapes.SphereShape{
				Radius: 10.0,
			},
			Type:   physics.PHYSICS_BODY_TYPE_DYNAMIC,
			Weight: 15.0,
		}

		player.BodyId = gameState.PhysicsWorld.AddPhysicsBody(body)

		gameState.Scene.AddEntity(&player)
	}

	{ // Spining object entity
		ent := entities.SpinningEntity{
			Velocity: 1.0,
		}

		mesh := core.Mesh{
			Vertices: obj.Vertices,
			Indices:  obj.Indices,
		}

		meshHandle := mesh.Bind(
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
			MeshHandle: meshHandle,
			Material:   &material,
			Transform:  transform,
		}

		ent.ModelId = gameState.Scene.AddModel(model)
		gameState.Scene.AddEntity(&ent)
	}

	{ // 3D map to test marching cubes

		ent := entities.SpinningEntity{}

		grid := structs.Grid3D[float32]{}
		grid.New(5, 5, 5)

		grid.Place(1.0, 2, 2, 2)
		grid.Place(1.0, 2, 3, 2)
		grid.Place(1.0, 3, 2, 2)
		grid.Place(1.0, 3, 3, 2)

		meshBuilder := procedural.MarchingCubes(grid)
		meshBuilder.IncludeId = true
		meshBuilder.Scale = mgl32.Vec3{1, 1, 1}

		mesh := meshBuilder.Build(false)

		meshHandle := mesh.Bind(core.MeshAttributes.Position(), core.MeshAttributes.Float())

		material := materials.ColorMaterial{
			Color: structs.RGB(1.0, 0.0, 0.0),
		}

		transform := structs.NewTransform()

		model := core.Model{
			MeshHandle: meshHandle,
			Material:   &material,
			Transform:  transform,
		}

		ent.ModelId = gameState.Scene.AddModel(model)
		gameState.Scene.AddEntity(&ent)
	}

	lastTime, now := 0.0, 0.0

	for !window.ShouldClose() {
		delta := now - lastTime

		// Update fase
		for _, ent := range gameState.Scene.Entities {
			ent.Update(&gameState, float32(delta))
		}

		gameState.PhysicsWorld.Update(float32(delta))

		// Rendering fase
		renderer.Clear()
		pass := renderer.BeginDraw(gameState.Camera)

		for _, model := range gameState.Scene.Models {
			pass.DrawMesh(model.MeshHandle, model.Transform, model.Material)
		}

		pass.EndDraw()

		//TODO figure out a way to make this suck less
		systems.InputSystem.Update()

		window.PollAndSwap()

		lastTime = now
		now = core.GetTime()
	}
}
