package main

import (
	"github.com/mwdomino/go-cgol-lua/config"
	"github.com/mwdomino/go-cgol-lua/game"
	"github.com/mwdomino/go-cgol-lua/output/gui"
	"image/color"
)

func main() {
	config := config.Config{
		Rows: 100,
		Cols: 100,
	}
	game := &game.Game{Config: &config}
	out := gui.Gui{
		WindowSizeX:     config.Cols * 10,
		WindowSizeY:     config.Rows * 10,
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
