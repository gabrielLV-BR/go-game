package gamedata

import (
	"slices"
)

type EventType int

const (
	INPUT_EVENT   EventType = iota
	COMMAND_EVENT           = iota
)

type EventMap struct {
	mapping map[EventType][]EventListener

	// reverse mapping to speed up removal
	listenerToTypes map[EventListener][]EventType
}

func (eventMap *EventMap) New() {
	eventMap.mapping = make(map[EventType][]EventListener)
	eventMap.listenerToTypes = make(map[EventListener][]EventType)
}

func (eventMap *EventMap) Broadcast(event Event) {
	listeners, ok := eventMap.mapping[event.EventType()]

	if !ok {
		return
	}

	for _, listener := range listeners {
		listener.OnEvent(event)
	}
}

func (eventMap *EventMap) Subscribe(listener EventListener, interestedEventTypes ...EventType) {
	for _, eventType := range interestedEventTypes {
		if eventMap.mapping[eventType] == nil {
			eventMap.mapping[eventType] = []EventListener{}
		}

		eventMap.mapping[eventType] = append(eventMap.mapping[eventType], listener)
	}

	if _, ok := eventMap.listenerToTypes[listener]; !ok {
		eventMap.listenerToTypes[listener] = interestedEventTypes
	} else {
		eventMap.listenerToTypes[listener] = append(eventMap.listenerToTypes[listener], interestedEventTypes...)
	}
}

func (eventMap *EventMap) Unsubscribe(listener EventListener) {
	for _, interestedType := range eventMap.listenerToTypes[listener] {
		eventMap.mapping[interestedType] = slices.DeleteFunc(
			eventMap.mapping[interestedType],
			func(s EventListener) bool {
				return listener == s
			},
		)
	}

	delete(eventMap.listenerToTypes, listener)
}

type EventListener interface {
	OnEvent(Event)
}

type Event interface {
	EventType() EventType
}
