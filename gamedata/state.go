package gamedata

import (
	"gabriellv/game/physics"
	"gabriellv/game/structs"
)

type State struct {
	Camera *structs.Camera

	Scene Scene

	PhysicsWorld physics.PhysicsWorld
}
