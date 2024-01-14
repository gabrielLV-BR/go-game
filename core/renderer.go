package core

import "github.com/go-gl/gl/v3.3-core/gl"

type Renderer struct{}

func NewRenderer(window *Window) (Renderer, error) {
	renderer := Renderer{}

	err := gl.Init()

	if err != nil {
		return renderer, err
	}

	gl.Viewport(0, 0, int32(window.width), int32(window.height))

	return renderer, nil
}

func (this *Renderer) Resize(width, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}
