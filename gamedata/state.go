package gamedata

import (
	"gabriellv/game/core"
	"gabriellv/game/physics"
	"gabriellv/game/structs"
)

type State struct {
	Camera structs.Camera

	Entities []Entity
	Models   []core.Model

	PhysicsWorld physics.PhysicsWorld
}

func (state *State) AddModel(model core.Model) int {
	index := len(state.Models)

	state.Models = append(state.Models, model)

	return index
}

func (state *State) AddEntity(entity Entity) int {
	index := len(state.Entities)

	state.Entities = append(state.Entities, entity)

	return index
}
