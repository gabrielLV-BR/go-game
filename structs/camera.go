package structs

import (
	"gabriellv/game/utils"
	math "github.com/chewxy/math32"

	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	position  mgl32.Vec3
	direction mgl32.Vec3

	// to minimize recalculation
	dirty            bool
	viewMatrix       mgl32.Mat4
	projectionMatrix mgl32.Mat4
}

func NewCamera() Camera {
	camera := Camera{
		position:  mgl32.Vec3{},
		direction: mgl32.Vec3{},
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

func (camera *Camera) LookAt(target mgl32.Vec3) {
	if target.ApproxEqual(camera.position) {
		panic("Camera's TARGET and POSITION must not be equal")
	}

	camera.direction = target.Sub(camera.position).Normalize()
	camera.dirty = true
}

func (camera *Camera) Rotate(yawSpeed, pitchSpeed float32) {
	if mgl32.FloatEqual(yawSpeed, pitchSpeed) && mgl32.FloatEqual(yawSpeed, 0.0) {
		return
	}

	// veeery convoluted, seek better ways to do this
	direction := camera.direction

	yaw := math.Atan2(direction.X(), direction.Z())
	pitch := math.Asin(-direction.Y())

	yaw += yawSpeed
	pitch += pitchSpeed

	direction = utils.AnglesToVector(yaw, pitch)

	camera.LookAt(camera.position.Add(direction))
}

func (camera *Camera) recalculateViewMatrix() {
	up := mgl32.Vec3{0.0, 1.0, 0.0}

	right := up.Cross(camera.direction)
	up = camera.direction.Cross(right)

	target := camera.position.Add(camera.direction.Mul(10.0))

	camera.viewMatrix = mgl32.LookAtV(
		camera.position, target, up,
	).Inv()

	camera.dirty = false
}

func (camera *Camera) UsePerspectiveProjection(fov, aspect, near, far float32) {
	camera.projectionMatrix = mgl32.Perspective(fov, aspect, near, far)
}

func (camera *Camera) UseOrthogonalProjection(width, height, near, far float32) {
	camera.projectionMatrix = mgl32.Ortho(0.0, width, height, 0.0, near, far)
}
