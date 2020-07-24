package ludok

import (
	"errors"
	"fmt"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	// Window can be resized
	WINDOW_RESIZABLE = sdl.WINDOW_RESIZABLE
	// Window is minimized
	WINDOW_MINIMIZED = sdl.WINDOW_MINIMIZED
	// Window is maximized
	WINDOW_MAXIMIZED = sdl.WINDOW_MAXIMIZED
	// No window decoration
	WINDOW_BORDERLESS = sdl.WINDOW_BORDERLESS
)

var (
	Load   func()           = func() {}
	Update func(dt float64) = func(dt float64) {}
	Draw   func()           = func() {}

	Keyboard func(event KeyboardEvent) = func(event KeyboardEvent) {}

	MouseMoved  func(event MouseMovedEvent)  = func(event MouseMovedEvent) {}
	MouseButton func(event MouseButtonEvent) = func(event MouseButtonEvent) {}
)

func Run(config Config) (err error) {
	if (config != Config{}) {
		cfg = config
	}

	sdl.Main(func() {
		err = run()
	})

	return
}

// Do the specified function in the main thread.
// Preferred to use it then you do something directly in SDL or OpenGL
func Do(f func()) {
	sdl.Do(f)
}

func Exit() {
	setRunning(false)
}

var (
	cfg Config = Config{
		WindowWidth:  800,
		WindowHeight: 600,
		WindowTitle:  "Ludok - 2D Game Framework",
		VSync:        true,
	}

	running      bool = true
	runningMutex sync.Mutex

	window *sdl.Window
)

func setRunning(value bool) {
	runningMutex.Lock()
	running = value
	runningMutex.Unlock()
}

func run() (err error) {
	Do(func() {
		window, err = sdl.CreateWindow(cfg.WindowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, cfg.WindowWidth, cfg.WindowHeight,
			sdl.WINDOW_SHOWN|sdl.WINDOW_ALLOW_HIGHDPI|sdl.WINDOW_OPENGL|cfg.Flags)
	})
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to create window: %s\n", err))
	}
	defer Do(func() {
		window.Destroy()
	})
	{
		var vsync uint32
		if cfg.VSync {
			vsync = sdl.RENDERER_PRESENTVSYNC
		}
		Do(func() {
			Graphics.renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|vsync)
		})
	}
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to create renderer: %s\n", err))
	}
	defer Do(func() {
		Graphics.renderer.Destroy()
	})

	Do(func() {
		Graphics.mainCanvas = &Canvas{canvas: Graphics.renderer.GetRenderTarget()}
	})

	Load()

	for running {
		// Handling Events
		Do(func() {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch event.(type) {
				case *sdl.QuitEvent:
					setRunning(false)
				case *sdl.KeyboardEvent:
					e := event.(*sdl.KeyboardEvent)
					Keyboard(KeyboardEvent{
						State:    InputEventType(e.State),
						Scancode: uint(e.Keysym.Scancode),
						Symcode:  KeyboardKeycode(e.Keysym.Sym),
						Mod:      KeyboardEventModifier(e.Keysym.Mod),
						Repeated: e.Repeat,
					})
				case *sdl.MouseMotionEvent:
					e := event.(*sdl.MouseMotionEvent)
					MouseMoved(MouseMovedEvent{
						X:      int(e.X),
						Y:      int(e.Y),
						DeltaX: int(e.XRel),
						DeltaY: int(e.YRel),
						Button: MouseEventButton(e.State),
					})
				case *sdl.MouseButtonEvent:
					e := event.(*sdl.MouseButtonEvent)
					MouseButton(MouseButtonEvent{
						State:  InputEventType(e.State),
						X:      int(e.X),
						Y:      int(e.Y),
						Clicks: int(e.Clicks),
						Button: MouseEventButton(e.Button),
					})
				}
			}
		})

		Do(func() {
			Graphics.Clear(Lime)
		})

		calcFrameTime()
		Update(dt)
		Draw()

		Do(func() {
			Graphics.ResetCanvas()
			Graphics.Present()
		})
	}

	return nil
}
