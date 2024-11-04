package sdlmgr

import (
	"errors"
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type SDLManager struct {
	Window   *sdl.Window
	Surface  *sdl.Surface
	Renderer *sdl.Renderer
}

/*
Returns a new SDLManager instance. All members are nil pointers at creation,
so Init method must be called to initialize them.
*/
func New() *SDLManager {
	return &SDLManager{
		nil,
		nil,
		nil,
	}
}

/*
Initializes the SDL subsystem and populates the calling instance with its
corresponding window, surface and renderer pointers. Must be called for
the instance to be used. Error is returned on any mishap
*/
func (s *SDLManager) Init() error {
	err := sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to initialize SDL Video subsystem")
		return errors.New("SDL_VIDEO_INIT_FAILURE")
	}

	window, err := sdl.CreateWindow("CI6450 Proyecto",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		ScreenWidth,
		ScreenHeight,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "SDL Window creation failed.")
		return errors.New("SDL_WINDOW_CREATE_FAIL")
	}
	s.Window = window

	surface, err := s.Window.GetSurface()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get surface from window")
		return errors.New("SDL_SURFACE_FETCH_FAIL")
	}
	s.Surface = surface

	renderer, err := s.Window.GetRenderer()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get renderer from window")
		return errors.New("SDL_RENDERER_FETCH_FAIL")
	}
	s.Renderer = renderer

	return nil
}

/*
Renders the current window surface
*/
func (s *SDLManager) Render() {
	s.Renderer.Present()
}

/*
Clears the screen
*/
func (s *SDLManager) Clear() {
	s.Renderer.SetDrawColor(0, 0, 0, 1)
	s.Renderer.Clear()
}

/*
Gets all input state from the SDL subsystem and returns it
*/
func (s *SDLManager) GetInput() []uint8 {
	sdl.PumpEvents()
	return sdl.GetKeyboardState()
}

/*
Polls SDL events
*/
func (s *SDLManager) PollEvents() sdl.Event {
	return sdl.PollEvent()
}
