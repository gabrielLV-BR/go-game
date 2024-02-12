package gamedata

type Entity interface {
	Update(state *State, delta float32)
}
