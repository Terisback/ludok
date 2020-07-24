package ludok

import (
	"errors"
	"fmt"
	"image/color"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type GraphicsHandler struct {
	renderer   *sdl.Renderer
	currCanvas *Canvas
	mainCanvas *Canvas
}

var (
	Graphics *GraphicsHandler = &GraphicsHandler{}
)

func (g *GraphicsHandler) Clear(clearColor color.RGBA) {
	cr, cg, cb, ca, err := g.renderer.GetDrawColor()
	if err != nil {
		panic(errors.New(fmt.Sprintf("Failed to get draw color: %s\n", err)))
	}

	g.SetDrawColor(clearColor)
	err = g.renderer.Clear()
	if err != nil {
		panic(errors.New(fmt.Sprintf("Failed to clear with renderer: %s\n", err)))
	}

	oldColor := color.RGBA{R: cr, G: cg, B: cb, A: ca}
	g.SetDrawColor(oldColor)
}

func (g *GraphicsHandler) SetDrawColor(color color.RGBA) {
	err := g.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	if err != nil {
		panic(errors.New(fmt.Sprintf("Failed to set draw color: %s\n", err)))
	}
}

func (g *GraphicsHandler) GetCanvas() *Canvas {
	if g.currCanvas != nil {
		return g.currCanvas
	} else {
		return g.mainCanvas
	}
}

func (g *GraphicsHandler) SetCanvas(canvas *Canvas) {
	g.currCanvas = canvas
	err := g.renderer.SetRenderTarget(g.currCanvas.canvas)
	if err != nil {
		panic(errors.New(fmt.Sprintf("Failed to set render target: %s\n", err)))
	}
}

func (g *GraphicsHandler) DrawCanvas(canvas *Canvas, x, y, w, h int) {
	g.renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	ow, oh := canvas.GetSize()
	g.renderer.Copy(canvas.canvas,
		&sdl.Rect{X: 0, Y: 0, W: int32(ow), H: int32(oh)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)})
}

func (g *GraphicsHandler) ResetCanvas() {
	g.currCanvas = nil
	err := g.renderer.SetRenderTarget(g.mainCanvas.canvas)
	if err != nil {
		panic(errors.New(fmt.Sprintf("Failed to reset render target: %s\n", err)))
	}
}

func (g *GraphicsHandler) GetSize() (width, height int) {
	x, y := window.GetSize()
	return int(x), int(y)
}

func (g *GraphicsHandler) Present() {
	g.renderer.Present()
}

func (g *GraphicsHandler) Print(x, y int, text string, color color.RGBA) {
	ok := gfx.StringColor(g.renderer, int32(x), int32(y), text, sdl.Color(color))
	if !ok {
		panic(errors.New("Failed to draw a box.\n"))
	}
}

func (g *GraphicsHandler) Box(x1, y1, x2, y2 int, color color.RGBA) {
	ok := gfx.BoxColor(g.renderer, int32(x1), int32(y1), int32(x2), int32(y2), sdl.Color(color))
	if !ok {
		panic(errors.New("Failed to draw a box.\n"))
	}
}

type Canvas struct {
	canvas *sdl.Texture
}

func NewCanvas(width, height int) *Canvas {
	texture, err := Graphics.renderer.CreateTexture(uint32(sdl.PIXELFORMAT_RGBA32), sdl.TEXTUREACCESS_TARGET, int32(width), int32(height))
	if err != nil {
		panic(errors.New(fmt.Sprintf("Failed to create texture: %s\n", err)))
	}
	err = texture.SetBlendMode(sdl.BLENDMODE_BLEND)
	if err != nil {
		panic(errors.New(fmt.Sprintf("Failed to set blend mode for texture: %s\n", err)))
	}
	return &Canvas{canvas: texture}
}

func (c *Canvas) GetSize() (width, height int) {
	_, _, x, y, err := c.canvas.Query()
	if err != nil {
		// TODO: Do this thing with event system (call OnError)
		panic(errors.New(fmt.Sprintf("Failed to get size of the canvas: %s\n", err)))
		return 0, 0
	}
	return int(x), int(y)
}
