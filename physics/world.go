package physics

import (
	"github.com/go-gl/mathgl/mgl32"
)

var (
	GRAVITY = mgl32.Vec3{0, -9.8, 0}
)

type PhysicsWorld struct {
	bodies []PhysicsBody
}

func (world *PhysicsWorld) Update(delta float32) {

	// add forces to the body
	//TODO make this more complex
	for _, body := range world.bodies {
		body.LinearAcceleration = body.LinearAcceleration.Add(GRAVITY)

		body.LinearVelocity = body.LinearVelocity.Add(body.LinearAcceleration)
		body.Transform.Position = body.Transform.Position.Add(body.LinearVelocity)
	}

	// collision checks, horribly inneficient
	//TODO make this suck less
	for _, body1 := range world.bodies {
		for _, body2 := range world.bodies {
			if body1 == body2 {
				continue
			}

			if !BoundingBoxCollide(body1.BoundingBox, body2.BoundingBox) {
				continue
			}

			if GJKCollide(body1.Shape, body2.Shape) {
				//TODO resolve collision
			}
		}
	}

}
