package physics

import (
	"gabriellv/game/structs"

	"github.com/go-gl/mathgl/mgl32"
)

type PhysicsBodyType int

// pretty blatantly copying Unity/Godot on this
const (
	PHYSICS_BODY_TYPE_STATIC PhysicsBodyType = 0
	PHYSICS_BODY_TYPE_DYNAMIC
	PHYSICS_BODY_TYPE_KINEMATIC
)

type PhysicsBody struct {
	Transform   structs.Transform
	Shape       PhysicsShape
	BoundingBox BoundingBox

	// for now, I'll stick to Euler's integration for movement
	// I do know about the supposed superior numerical stability
	// of Verlet's integration, however, since I've never actually
	// implemented a physics system from scratch, I'll stick to the
	// simpler methods and eventually do some benchmarks and test
	// if the change is worthwhile
	LinearVelocity     mgl32.Vec3
	LinearAcceleration mgl32.Vec3

	Type PhysicsBodyType

	Weight float32

	// not going to do this for now as I have absolutely no idea
	// on how these are implemented LMAO

	// AngularVelocity mgl32.Vec3
	// AngularAcceleration mgl32.Vec3
	// FrictionCoefficient float32
}

type PhysicsShape interface {
	SupportPoint(dir mgl32.Vec3) mgl32.Vec3
}
