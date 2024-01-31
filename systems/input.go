package systems

import "github.com/go-gl/glfw/v3.3/glfw"

type InputSystem struct{}

func NewInputSystem() InputSystem {
	return InputSystem{}
}

func InputKeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	inputSystem := (*InputSystem)(window.GetUserPointer())
}

func InputMouseMotionCallback(window *glfw.Window, x, y float64) {
	inputSystem := (*InputSystem)(window.GetUserPointer())
}

func InputMouseButtonCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	inputSystem := (*InputSystem)(window.GetUserPointer())
}
