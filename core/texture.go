package core

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Texture struct {
	handle uint32
	target uint32
	unit   uint32
}

func (texture *Texture) Bind(unit uint32) {
	gl.ActiveTexture(unit)
	gl.BindTexture(texture.target, texture.handle)
	texture.unit = unit
}

func (texture *Texture) Unbind() {
	gl.BindTexture(texture.target, 0)
	texture.unit = 0
}

func (texture *Texture) Unit() uint32 {
	return texture.unit
}

func NewTexture(img image.Image) Texture {
	// store image in RGBA format
	//TODO read img and figure out format used
	rgba := image.NewRGBA(img.Bounds())

	draw.Draw(rgba, rgba.Bounds(), img, image.Pt(0, 0), draw.Src)

	var id uint32 = 0
	gl.GenTextures(1, &id)

	target      := uint32(gl.TEXTURE_2D)
	internalFmt := int32(gl.SRGB_ALPHA)
	format      := uint32(gl.RGBA)
	width       := int32(rgba.Rect.Size().X)
	height      := int32(rgba.Rect.Size().Y)
	pixType     := uint32(gl.UNSIGNED_BYTE)
	dataPtr     := gl.Ptr(rgba.Pix)

	texture := Texture{
		handle: id,
		target: target,
	}

	texture.Bind(gl.TEXTURE0)

	// set the texture wrapping/filtering options (applies to current bound texture obj)
	// TODO-cs
	gl.TexParameteri(texture.target, gl.TEXTURE_WRAP_R, gl.MIRRORED_REPEAT)
	gl.TexParameteri(texture.target, gl.TEXTURE_WRAP_S, gl.MIRRORED_REPEAT)
	gl.TexParameteri(texture.target, gl.TEXTURE_MIN_FILTER, gl.LINEAR)  // minification filter
	gl.TexParameteri(texture.target, gl.TEXTURE_MAG_FILTER, gl.LINEAR)  // magnification filter

	gl.TexImage2D(target, 0, internalFmt, width, height, 0, format, pixType, dataPtr)

	gl.GenerateMipmap(texture.handle)

	texture.Unbind()

	return texture
}

func LoadTexture(path string) (Texture, error) {
	file, err := os.Open(path)

	if err != nil {
		return Texture{}, err
	}

	defer file.Close()

	// image.Decode figure out the extension automatically
	pixels, _, err := image.Decode(file)

	if err != nil {
		return Texture{}, err
	}

	return NewTexture(pixels), nil
}
