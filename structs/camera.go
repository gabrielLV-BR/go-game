package structs

import (
	math "github.com/chewxy/math32"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	CAMERA_PITCH_MIN = -math.Pi/2.0 + 0.1
	CAMERA_PITCH_MAX = +math.Pi/2.0 - 0.1
)

type Camera struct {
	position mgl32.Vec3

	yaw   float32
	pitch float32

	// to minimize recalculation
	dirty            bool
	viewMatrix       mgl32.Mat4
	projectionMatrix mgl32.Mat4
}

func NewCamera() Camera {
	camera := Camera{
		position: mgl32.Vec3{},
		yaw:      0.0,
		pitch:    0.0,
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

func (camera *Camera) Direction() mgl32.Vec3 {
	direction := mgl32.Vec3{}

	direction[0] = math.Cos(camera.yaw) * math.Cos(camera.pitch)
	direction[1] = math.Sin(camera.pitch)
	direction[2] = math.Sin(camera.yaw) * math.Cos(camera.pitch)

	return direction.Normalize()
}

func (camera *Camera) Translate(velocity mgl32.Vec3) {
	camera.position = camera.position.Add(velocity)
	camera.dirty = true
}

func (camera *Camera) TranslateLocal(velocity mgl32.Vec3) {
	forward := camera.Direction()
	forward[1] = 0.0
	forward = forward.Normalize()

	up := mgl32.Vec3{0, 1, 0}

	right := up.Cross(forward).Normalize()

	forward = forward.Mul(velocity.Z())
	right = right.Mul(velocity.X())

	camera.position = camera.position.Add(forward).Add(right)
	camera.dirty = true
}

func (camera *Camera) SetPosition(position mgl32.Vec3) {
	camera.position = position
	camera.dirty = true
}

func (camera *Camera) Rotate(yawSpeed, pitchSpeed float32) {
	camera.pitch += pitchSpeed
	camera.yaw -= yawSpeed

	camera.pitch = mgl32.Clamp(camera.pitch, CAMERA_PITCH_MIN, CAMERA_PITCH_MAX)

	camera.dirty = true
}

func (camera *Camera) recalculateViewMatrix() {
	up := mgl32.Vec3{0.0, 1.0, 0.0}

	direction := camera.Direction()
	right := up.Cross(direction)
	up = direction.Cross(right)

	target := camera.position.Add(direction.Mul(10.0))

	camera.LookAt(target, up)
}

func (camera *Camera) LookAt(target, up mgl32.Vec3) {
	camera.viewMatrix = mgl32.LookAtV(camera.position, target, up)

	camera.dirty = false
}

func (camera *Camera) UsePerspectiveProjection(fov, aspect, near, far float32) {
	camera.projectionMatrix = mgl32.Perspective(fov, aspect, near, far)
}

func (camera *Camera) UseOrthogonalProjection(width, height, near, far float32) {
	camera.projectionMatrix = mgl32.Ortho(0.0, width, height, 0.0, near, far)
}
