package structs

import (
	"gabriellv/game/utils"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	position mgl32.Vec3
	target   mgl32.Vec3

	// to minimize recalculation
	dirty            bool
	viewMatrix       mgl32.Mat4
	projectionMatrix mgl32.Mat4
}

func NewCamera() Camera {
	camera := Camera{
		position: mgl32.Vec3{},
		target:   mgl32.Vec3{},
		// try to save a little bit on unnecessary calculations
		dirty:            true,
		projectionMatrix: mgl32.Ident4(),
		viewMatrix:       mgl32.Ident4(),
	}

	return camera
}

func (camera *Camera) GetProjectionMatrix() mgl32.Mat4 {
	return camera.projectionMatrix
}

func (camera *Camera) GetViewMatrix() mgl32.Mat4 {
	if camera.dirty {
		camera.recalculateViewMatrix()
	}
	return camera.viewMatrix
}

func (camera *Camera) Translate(velocity mgl32.Vec3) {
	camera.position = camera.position.Add(velocity)
	camera.dirty = true
}

func (camera *Camera) SetPosition(position mgl32.Vec3) {
	camera.position = position
	camera.dirty = true
}

func (camera *Camera) Rotate(yawSpeed, pitchSpeed float32) {
	// veeery convoluted, seek better ways to do this
	direction := camera.target.Sub(camera.position)
	directionLen := direction.Len()

	direction[0] /= directionLen
	direction[1] /= directionLen
	direction[2] /= directionLen

	yaw := float32(math.Atan2(float64(direction.X()), float64(direction.Z())))
	pitch := float32(math.Asin(float64(-direction.Y())))

	yaw += yawSpeed
	pitch += pitchSpeed

	direction = utils.AnglesToVector(yaw, pitch)

	camera.target = camera.position.Add(direction)

	camera.LookAt(camera.position.Add(direction))
}

func (camera *Camera) recalculateViewMatrix() {
	up := mgl32.Vec3{0.0, 1.0, 0.0}
	direction := camera.target.Sub(camera.position).Normalize()

	right := up.Cross(direction)
	up = direction.Cross(right)

	target := camera.position.Add(direction.Mul(10.0))

	camera.viewMatrix = mgl32.LookAtV(
		camera.position, target, up,
	)

	camera.dirty = false
}

func (camera *Camera) LookAt(target mgl32.Vec3) {
	camera.target = target
	camera.dirty = true
}

func (camera *Camera) UsePerspectiveProjection(fov, aspect, near, far float32) {
	camera.projectionMatrix = mgl32.Perspective(fov, aspect, near, far)
}

func (camera *Camera) UseOrthogonalProjection(width, height, near, far float32) {
	camera.projectionMatrix = mgl32.Ortho(0.0, width, height, 0.0, near, far)
}
