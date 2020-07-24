package main

import (
	"fmt"
	"image/color"

	"github.com/Terisback/ludok"
)

var milliseconds, fps string

var tooltip *Tooltip
var drawingTooltip *Tooltip

var red bool = false

var drawingCanvas *ludok.Canvas

var mouseX, mouseY int
var pressed bool

// Loading all the things
func Load() {
	tooltip = NewTooltip(6, ludok.White, color.RGBA{R: 0, G: 0, B: 0, A: 80})
	x, y := ludok.Graphics.GetSize()
	drawingTooltip = NewTooltip(6, ludok.White, color.RGBA{R: 0, G: 0, B: 0, A: 80})
	drawingCanvas = ludok.NewCanvas(x, y)

	drawingTooltip.Update("Space - Clear canvas")
}

func Update(dt float64) {
	// Updating tooltip with fps info
	milliseconds = fmt.Sprintf("%.2f %s", ludok.GetAvgDeltaTime(), "ms")
	fps = fmt.Sprintf("%d %s", ludok.GetFPS(), "FPS")
	tooltip.Update(milliseconds + "\n" + fps)
}

func Draw() {
	// Clearing screen
	ludok.Graphics.Clear(color.RGBA{R: 100, G: 100, B: 140, A: 255})

	// Displaying indicator that spacebar is pressed
	if red {
		x, y := ludok.Graphics.GetSize()
		ludok.Graphics.Box(0, 0, x, y, ludok.Red)

		// Clearing drawing canvas
		ludok.Graphics.SetCanvas(drawingCanvas)
		ludok.Graphics.Clear(color.RGBA{R: 0, G: 0, B: 0, A: 0})
		ludok.Graphics.ResetCanvas()
	}

	// Drawing into canvas
	ludok.Graphics.SetCanvas(drawingCanvas)
	if pressed {
		ludok.Graphics.Box(mouseX-5, mouseY-5, mouseX+6, mouseY+6, ludok.Blue)
	}
	ludok.Graphics.ResetCanvas()

	// Drawing canvas on screen
	x, y := ludok.Graphics.GetSize()
	ludok.Graphics.DrawCanvas(drawingCanvas, 0, 0, x, y)
	ludok.Graphics.Box(mouseX-1, mouseY-1, mouseX+1, mouseY+2, ludok.Yellow)

	// Draw fps tooltip
	tooltip.Draw(5, 5)

	// Draw help tooltip
	drawingTooltip.Draw(x/2-drawingTooltip.Width/2, y-y/10-drawingTooltip.Height/2)
}

func Keyboard(event ludok.KeyboardEvent) {
	if event.Symcode == ludok.K_SPACE {
		if event.State == ludok.Pressed {
			red = true
		} else {
			red = false
		}
	}
}

func MouseMoved(event ludok.MouseMovedEvent) {
	mouseX, mouseY = event.X, event.Y
}

func MouseButton(event ludok.MouseButtonEvent) {
	if event.Button == ludok.BUTTON_LEFT {
		if event.State == ludok.Pressed {
			pressed = true
		} else {
			pressed = false
		}
	}
}

func main() {
	// That's shitty idea, I don't like that (but it's looking like Love2D and I like it)
	ludok.Load = Load
	ludok.Update = Update
	ludok.Draw = Draw
	ludok.Keyboard = Keyboard
	ludok.MouseMoved = MouseMoved
	ludok.MouseButton = MouseButton

	ludok.Run(ludok.Config{
		WindowWidth:  1280,
		WindowHeight: 720,
		WindowTitle:  "Ludok Test",
		VSync:        false,
	})
}
