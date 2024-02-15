package entities

import (
	"gabriellv/game/gamedata"

	"github.com/go-gl/mathgl/mgl32"
)

type SpinningEntity struct {
	ModelId  int
	Velocity float32

	rotation float32
}

func (ent *SpinningEntity) Update(state *gamedata.State, delta float32) {
	model := &state.Scene.Models[ent.ModelId]

	model.Transform.Rotation = mgl32.EulerToQuat(0.0, ent.rotation, 0.0)

	ent.rotation += delta * ent.Velocity
}
