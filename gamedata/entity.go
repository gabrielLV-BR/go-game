package gamedata

type Entity interface {
	Update(state *State, delta float32)
}

type EntityStruct struct {
	ModelId int
	Updater func(*State, float32)
}
