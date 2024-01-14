package main

import (
	"gabriellv/game/core"

	"github.com/go-gl/glfw/v3.0/glfw"
)

/*
	GO GAME
	-------
	Check README.md for more information
*/

const (
	WIDTH  = 500
	HEIGHT = 500
	TITLE  = "Window"
)

func main() {
	window, err := core.NewWindow(WIDTH, HEIGHT, TITLE)

	if err != nil {
		panic(err)
	}

	renderer, err := core.NewRenderer(&window)

	if err != nil {
		panic(err)
	}

	lastTime := core.GetTime()
	delta := 0.0

	for !window.ShouldClose() {
		delta = core.GetTime() - lastTime

		window.PollAndSwap()

		lastTime = glfw.GetTime()
	}
}
