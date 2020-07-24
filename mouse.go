package ludok

import "github.com/veandco/go-sdl2/sdl"

type MouseMovedEvent struct {
	X      int
	Y      int
	DeltaX int
	DeltaY int

	Button MouseEventButton
}

type MouseButtonEvent struct {
	// Event type
	//
	// Pressed, Released
	State InputEventType
	X     int
	Y     int

	Clicks int

	Button MouseEventButton
}

type MouseEventButton uint32

// Used in masks for mouse buttons
const (
	BUTTON_LEFT   MouseEventButton = sdl.BUTTON_LEFT
	BUTTON_RIGHT  MouseEventButton = sdl.BUTTON_RIGHT
	BUTTON_MIDDLE MouseEventButton = sdl.BUTTON_MIDDLE
	BUTTON_X1     MouseEventButton = sdl.BUTTON_X1
	BUTTON_X2     MouseEventButton = sdl.BUTTON_X2
)
