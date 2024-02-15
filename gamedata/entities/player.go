package entities

import (
	"gabriellv/game/controllers"
	"gabriellv/game/gamedata"
)

type Player struct {
	// this gets inserted when the player is constructed and allows us to
	// get the model and do stuff with it :)
	Controller controllers.FPSCameraController
	//TODO when implementing physics
	BodyId int
}

func (player *Player) Update(state *gamedata.State, delta float32) {
	player.Controller.Update(state.Camera, delta)
}
