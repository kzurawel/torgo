package display

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func New(title string, width float64, height float64, bgColor color.RGBA) *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  title,
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(bgColor)

	return win
}
