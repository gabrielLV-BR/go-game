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

	mousePosition mgl32.Vec2
	mouseButtons  map[glfw.MouseButton]float32
}

var InputSystem inputSystem

func (input *inputSystem) Init() {
	input.keys = make(map[glfw.Key]float32)
	input.mousePosition = mgl32.Vec2{}
	input.mouseButtons = make(map[glfw.MouseButton]float32)
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

func InputKeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		InputSystem.keys[key] = 1.0
	} else {
		InputSystem.keys[key] = 0.0
	}

	// make gofmt shut up
	_, _, _ = window, scancode, mods
}

func InputMouseMotionCallback(window *glfw.Window, x, y float64) {
	InputSystem.mousePosition[0] = float32(x)
	InputSystem.mousePosition[1] = float32(y)

	// make gofmt shut up
	_ = window
}

func InputMouseButtonCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		InputSystem.mouseButtons[button] = 1.0
	} else {
		InputSystem.mouseButtons[button] = 0.0
	}

	// make gofmt shut up
	_, _ = window, mods
}
