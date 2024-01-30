package core

import "github.com/go-gl/glfw/v3.3/glfw"

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
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window.handle, err = glfw.CreateWindow(width, height, title, nil, nil)

	if err != nil {
		return window, err
	}

	// make context
	window.handle.MakeContextCurrent()

	return window, nil
}

func (window *Window) Destroy() {
	glfw.Terminate()
}

func (window *Window) Resize(width, height int) {
	window.handle.SetSize(int(width), int(height))
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
