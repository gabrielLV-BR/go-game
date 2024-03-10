package events

import (
	"gabriellv/game/gamedata"
)

type InputEvent struct {
	Key string
}

type InputListener struct{}

func (listener InputListener) OnEvent(event gamedata.Event) {
	e := event.(InputEvent)

	println(e.Key)
}

func (inputEvent InputEvent) EventType() gamedata.EventType {
	return gamedata.INPUT_EVENT
}
