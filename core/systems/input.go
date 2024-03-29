package systems

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// unfortunately the input system will have to be a global instance :'(
// really couldn't get the window user pointer to work
// also, it kinda makes sense for the input system to be global
// i can't think of a situation where multiple input systems would be helpful

type inputSystem struct {
	keys map[glfw.Key]float32

	mousePosition     mgl32.Vec2
	lastMousePosition mgl32.Vec2
	mouseButtons      map[glfw.MouseButton]float32
}

var InputSystem inputSystem

func (input *inputSystem) Init() {
	input.keys = make(map[glfw.Key]float32)
	input.mousePosition = mgl32.Vec2{}
	input.lastMousePosition = mgl32.Vec2{}
	input.mouseButtons = make(map[glfw.MouseButton]float32)

	// register wanted keys

	ensuredKeys := []glfw.Key{
		glfw.KeyW, glfw.KeyA, glfw.KeyS, glfw.KeyD,
	}

	for _, key := range ensuredKeys {
		input.keys[key] = 0.0
	}
}

func (input *inputSystem) Update() {
	input.lastMousePosition = input.mousePosition
}

func (input *inputSystem) IsKeyDown(key glfw.Key) bool {
	val, ok := input.keys[key]

	return ok && val > 0.5
}

func (input *inputSystem) GetAxis(neg, pos glfw.Key) float32 {
	negative, ok := input.keys[neg]

	if !ok {
		negative = 0.0
	}

	positive, ok := input.keys[pos]

	if !ok {
		positive = 0.0
	}

	return positive - negative
}

func (input *inputSystem) GetMouseDelta() mgl32.Vec2 {
	return input.mousePosition.Sub(input.lastMousePosition)
}

func InputKeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		if window.GetInputMode(glfw.CursorMode) != glfw.CursorNormal && key == glfw.KeyEscape {
			window.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
		}

		InputSystem.keys[key] = 1.0
	} else if action == glfw.Release {
		InputSystem.keys[key] = 0.0
	}

	// make gofmt shut up
	_, _, _ = window, scancode, mods
}

func InputMouseMotionCallback(window *glfw.Window, x, y float64) {
	InputSystem.lastMousePosition = InputSystem.mousePosition

	InputSystem.mousePosition[0] = float32(x)
	InputSystem.mousePosition[1] = float32(y)

	// make gofmt shut up
	_ = window
}

func InputMouseButtonCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		if window.GetInputMode(glfw.CursorMode) != glfw.CursorDisabled {
			window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
		}

		InputSystem.mouseButtons[button] = 1.0
	} else if action == glfw.Release {
		InputSystem.mouseButtons[button] = 0.0
	}

	// make gofmt shut up
	_, _ = window, mods
}
