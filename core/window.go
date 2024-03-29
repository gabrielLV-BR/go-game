package core

import (
	"gabriellv/game/core/systems"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
	handle        *glfw.Window
	width, height int
	title         string
}

func NewWindow(width, height int, title string) (Window, error) {
	window := Window{
		handle: nil,
		width:  width,
		height: height,
		title:  title,
	}

	err := glfw.Init()

	if err != nil {
		return window, err
	}

	// set variables

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window.handle, err = glfw.CreateWindow(width, height, title, nil, nil)

	if err != nil {
		return window, err
	}

	// make context
	window.handle.MakeContextCurrent()

	return window, nil
}

func (window *Window) AspectRatio() float32 {
	return float32(window.width) / float32(window.height)
}

func (window *Window) SetupInputSystem() {
	systems.InputSystem.Init()

	window.handle.SetKeyCallback(systems.InputKeyCallback)
	window.handle.SetCursorPosCallback(systems.InputMouseMotionCallback)
	window.handle.SetMouseButtonCallback(systems.InputMouseButtonCallback)
}

func (window *Window) Destroy() {
	glfw.Terminate()
}

func (window *Window) Size() (int, int) {
	return window.width, window.height
}

func (window *Window) Resize(width, height int) {
	window.handle.SetSize(width, height)
	window.width = width
	window.height = height
}

func GetTime() float64 {
	return glfw.GetTime()
}

func (window *Window) ShouldClose() bool {
	return window.handle.ShouldClose()
}

func (window *Window) PollAndSwap() {
	glfw.PollEvents()
	window.handle.SwapBuffers()
}
