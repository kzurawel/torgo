package main

import (
	"github.com/faiface/pixel/pixelgl"
  "golang.org/x/image/colornames"
  "github.com/kzurawel/torgo/pkg/display"
)

func run() {
  win := display.New("Torgo", 1024, 768, colornames.Green)

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
