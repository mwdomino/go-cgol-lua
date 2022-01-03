package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	util "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mwdomino/go-cgol-lua/game"
	"image/color"
	"log"
	"os"
)

// Gui is one of the available output types
// Specifically, it runs an `ebiten` based GUI
//
// Gui implements the ebiten.Game interface
type Gui struct {
	WindowSizeX     int
	WindowSizeY     int
	GridSpacing     int
	GridColor       color.Color
	BackgroundColor color.Color
	WindowTitle     string
	Game            *game.Game
}

func (g *Gui) Init() {
	ebiten.SetWindowSize(g.WindowSizeX, g.WindowSizeY)
	ebiten.SetWindowTitle(g.WindowTitle)
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetMaxTPS(2)
}

func (g *Gui) Run() {
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}

func (g *Gui) Update() error {
	g.Game.Tick()
	if g.Game.CurrentBoard.IsEmpty() {
		os.Exit(0)
	}
	return nil
}

// Draw draws the game screen.
func (g *Gui) Draw(screen *ebiten.Image) {
	g.DrawCurrentBoard(screen)
}

func (g *Gui) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.WindowSizeX, g.WindowSizeY
}

func (g *Gui) DrawBackground(screen *ebiten.Image) {
	screen.Fill(g.BackgroundColor)
}

// DrawGrid draws a grid of gridSpacing sized squares
func (g *Gui) DrawGrid(screen *ebiten.Image) {
	g.DrawBackground(screen)
	x := float64(g.WindowSizeX)
	y := float64(g.WindowSizeY)
	gs := g.GridSpacing

	// horizontal
	for i := float64(gs); i < x; i += float64(gs) {
		util.DrawLine(screen, 0, i, x, i, g.GridColor)
		util.DrawLine(screen, i, 0, i, y, g.GridColor)
	}

	// vertical
	for i := float64(gs); i < y; i += float64(gs) {
		util.DrawLine(screen, 0, i, x, i, g.GridColor)
		util.DrawLine(screen, i, 0, i, y, g.GridColor)
	}
}

func (g *Gui) DrawCurrentBoard(screen *ebiten.Image) {
	g.DrawGrid(screen)
	gs := float64(g.GridSpacing)
	cb := g.Game.CurrentBoard

	// TODO - efficiency?
	for i := 0; i < len(cb); i++ {
		for j := 0; j < len(cb[i]); j++ {
			if cb[i][j] == 1 {
				x := float64(i)
				y := float64(j)
				util.DrawRect(screen, x*gs, y*gs, gs-1, gs, g.GridColor)
			}
		}
	}
}
