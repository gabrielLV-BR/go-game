package physics

import (
	"gabriellv/game/physics/shapes"

	"github.com/go-gl/mathgl/mgl32"
)

var (
	GRAVITY = mgl32.Vec3{0, -9.8, 0}
)

type PhysicsWorld struct {
	bodies []PhysicsBody
}

func (world *PhysicsWorld) AddPhysicsBody(body PhysicsBody) int {
	index := len(world.bodies)

	world.bodies = append(world.bodies, body)

	return index
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

			if !shapes.BoundingBoxCollide(
				body1.TransformedBoundingBox(),
				body2.TransformedBoundingBox(),
			) {
				continue
			}

			// for now, will only check on spheres to
			// make sure collision response is working

			var sphere1, sphere2 *shapes.SphereShape
			var ok bool

			sphere1, ok = body1.Shape.(*shapes.SphereShape)
			if !ok {
				continue
			}

			sphere2, ok = body2.Shape.(*shapes.SphereShape)
			if !ok {
				continue
			}

			penetration := shapes.SpheresPenetrationVector(
				*sphere1, body1.Transform.Position,
				*sphere2, body2.Transform.Position,
			)

			if penetration.Len() > 0 {
				// dumb static resolve
				penetration = penetration.Mul(0.5)

				body1.Transform.Position.Sub(penetration)
				body2.Transform.Position.Sub(penetration)
			}
		}
	}

}
