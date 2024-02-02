package controllers

import (
	"gabriellv/game/structs"
	"gabriellv/game/systems"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type FPSCameraController struct {
	Speed            float32
	MouseSensitivity float32
}

// TODO decouple this function from the input system and allow some sort of DI
func (controller *FPSCameraController) Update(camera *structs.Camera, delta float32) {
	linearVelocity := mgl32.Vec3{0.0, 0.0, 0.0}
	angularVelocity := mgl32.Vec2{}

	//TODO add remapping hability
	linearVelocity[0] = systems.InputSystem.GetAxis(glfw.KeyA, glfw.KeyD)
	linearVelocity[2] = systems.InputSystem.GetAxis(glfw.KeyS, glfw.KeyW)

	velLen := linearVelocity.Len()

	if !mgl32.FloatEqual(velLen, 0.0) {
		linearVelocity = linearVelocity.Mul(1.0 / velLen)
	}

	linearVelocity = linearVelocity.Mul(controller.Speed * delta)

	angularVelocity = systems.InputSystem.GetMouseDelta().Mul(controller.MouseSensitivity * delta)

	camera.Translate(linearVelocity)
	camera.Rotate(angularVelocity.X(), angularVelocity.Y())
}
