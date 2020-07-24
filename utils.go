package ludok

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func SetFullscreen(fullscreen bool) (err error) {
	var dIndex int
	Do(func() {
		dIndex, err = window.GetDisplayIndex()
	})
	if err != nil {
		log.Printf("Failed to get display index: %s\n", err)
		return err
	}
	var dBounds sdl.Rect
	Do(func() {
		dBounds, err = sdl.GetDisplayBounds(dIndex)
	})
	if err != nil {
		log.Printf("Failed to get display bounds: %s\n", err)
		return err
	}

	sdl.Do(func() {
		window.Restore()

		if fullscreen {
			window.SetSize(cfg.WindowWidth, cfg.WindowHeight)
			window.SetPosition(dBounds.X+dBounds.W/2-cfg.WindowWidth/2, dBounds.Y+dBounds.H/2-cfg.WindowHeight/2)
		} else {
			window.SetSize(dBounds.W, dBounds.H+10)
			window.SetPosition(dBounds.X, dBounds.Y)
		}
	})

	return nil
}
