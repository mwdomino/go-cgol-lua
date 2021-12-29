package game

import (
	"fmt"
	"github.com/mwdomino/go-cgol-lua/config"
	"math/rand"
)

type Board [][]int

func (b Board) IsInBounds(x, y int) bool {
	return (y > 0 &&
		y < len(b) &&
		x > 0 &&
		x < len(b[y]))
}

// IsEmpty returns a boolean representing whether the game has completed
// and all cells have died
func (b Board) IsEmpty() bool {
	for y := 0; y < len(b); y++ {
		for x := 0; x < len(b[y]); x++ {
			if b[y][x] == 1 {
				return false
			}
		}
	}
	return true
}

// Game implements ebiten.Game interface.
type Game struct {
	CurrentBoard Board
	nextBoard    Board

	Config *config.Config
}

// Init initializes a board of sizeX by sizeY dimensions
func (g *Game) Init() {
	x := g.Config.GridSizeX
	y := g.Config.GridSizeY

	g.CurrentBoard = initBoard(x, y)
	g.nextBoard = initBoard(x, y)
	g.generateRandomBoard(850)
}

// Tick iterates and pushes the game forward one cycle
// TODO run all calculations in goroutines and collect results
func (g *Game) Tick() {
	c := g.CurrentBoard
	n := g.nextBoard
	for x := 0; x < len(c); x++ {
		for y := 0; y < len(c[x]); y++ {
			n[x][y] = g.calculateCellUpdate(x, y)
		}
	}
	g.CurrentBoard = n
}

// calculateCellUpdate() determines if a cell should be alive or dead on the next round
// returns 0 if cell is dead next round
// returns 1 if cell is alive next round
func (g *Game) calculateCellUpdate(col int, row int) int {
	/*
		(-1,1)(0,1)(1,1)
		(-1,0)(0,0)(1,0)
		(-1,-1)(0,-1)(1,-1)
	*/
	// inner index is row
	// outer index is col
	// row is y
	// col is x
	b := g.CurrentBoard
	liveNeighbors := 0
	currentStatus := b[row][col]
	for x := -1; x < 1; x++ {
		for y := -1; y < 1; y++ {
			if b.IsInBounds(row+y, col+x) {
				if b[row+y][col+x] == 1 {
					liveNeighbors++
				}
			}
		}
	}

	/*
	   Any live cell with two or three live neighbours survives.
	   Any dead cell with three live neighbours becomes a live cell.
	   All other live cells die in the next generation. Similarly, all other dead cells stay dead.
	*/
	// dead with 3 live neighbors
	ret := 0
	if currentStatus == 0 && liveNeighbors == 3 {
		ret = 1
	}
	// live with 2 or 3 live neighbors
	if currentStatus == 1 && liveNeighbors >= 2 && liveNeighbors <= 3 {
		ret = 1
	}

	return ret
}

func (g *Game) clearBoard(board *Board) {
	b := *board
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			b[i][j] = 0
		}
	}
}

func (g *Game) generateRandomBoard(iter int) {
	b := g.CurrentBoard
	for i := 0; i < iter; i++ {
		x := rand.Intn(len(b))
		y := rand.Intn(len(b))
		b[x][y] = 1
	}
}

// InitBoard initializes a new board of x by y dimensions
func initBoard(x, y int) Board {
	var b Board
	b = make([][]int, y)
	for i := 0; i < x; i++ {
		b[i] = make([]int, x)
	}

	return b
}

func (g Game) DumpBoard() {
	b := g.CurrentBoard
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
	fmt.Println("")
}
