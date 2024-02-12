package gamedata

import (
	"gabriellv/game/controllers"
)

type Player struct {
	// this gets inserted when the player is constructed and allows us to
	// get the model and do stuff with it :)
	Controller controllers.FPSCameraController
	//TODO when implementing physics
	//BodyId uint32
}

func (player *Player) Update(state *State, delta float32) {
	player.Controller.Update(&state.Camera, delta)
}
