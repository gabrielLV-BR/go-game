package gamedata

import "gabriellv/game/core"

// a Scene contains the map and objects in a level,
// as well as some other data to allow interaction with it

type Scene struct {
	Entities []Entity
	Models   []core.Model
}

func (scene *Scene) AddModel(model core.Model) int {
	index := len(scene.Models)

	scene.Models = append(scene.Models, model)

	return index
}

func (scene *Scene) AddEntity(entity Entity) int {
	index := len(scene.Entities)

	scene.Entities = append(scene.Entities, entity)

	return index
}
