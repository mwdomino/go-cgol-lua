package main

import (
	"github.com/mwdomino/go-cgol-lua/config"
	"github.com/mwdomino/go-cgol-lua/game"
	"github.com/mwdomino/go-cgol-lua/output/gui"
	"image/color"
)

func main() {
	config := config.Config{
		Rows: 32,
		Cols: 32,
	}
	game := &game.Game{Config: &config}
	out := gui.Gui{
		WindowSizeX:     320,
		WindowSizeY:     320,
		GridSpacing:     10,
		GridColor:       color.Black,
		BackgroundColor: color.White,
		WindowTitle:     "Game of Life",
		Game:            game,
	}

	game.Init()
	out.Init()
	out.Run()
}
