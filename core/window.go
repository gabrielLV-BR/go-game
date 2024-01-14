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

func (this *Window) Resize(width, height int) {
	this.handle.SetSize(int(width), int(height))
	this.width = width
	this.height = height
}

func GetTime() float64 {
	return glfw.GetTime()
}

func (this *Window) ShouldClose() bool {
	return this.handle.ShouldClose()
}

func (this *Window) PollAndSwap() {
	glfw.PollEvents()
	this.handle.SwapBuffers()
}
