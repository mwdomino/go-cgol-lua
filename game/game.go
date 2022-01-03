package game

import (
	"fmt"
	"github.com/mwdomino/go-cgol-lua/config"
	"math/rand"
	"time"
)

// Game implements ebiten.Game interface.
type Game struct {
	CurrentBoard Board
	nextBoard    Board

	Config *config.Config
}

// Init initializes a board of sizeX by sizeY dimensions
func (g *Game) Init() {
	g.CurrentBoard = *initBoard(g.Config.Rows, g.Config.Cols)
	g.nextBoard = *initBoard(g.Config.Rows, g.Config.Cols)
	g.generateRandomBoard(12650)
}

// Tick iterates and pushes the game forward one cycle
// TODO run all calculations in goroutines and collect results
func (g *Game) Tick() {
	c := g.CurrentBoard
	n := g.nextBoard
	for row := 0; row < len(c); row++ {
		for col := 0; col < len(c[row]); col++ {
			n[row][col] = g.calculateCellUpdate(row, col)
		}
	}

	g.CurrentBoard = n
}

// calculateCellUpdate() determines if a cell should be alive or dead on the next round
// returns 0 if cell is dead next round
// returns 1 if cell is alive next round
func (g *Game) calculateCellUpdate(row int, col int) int {
	b := g.CurrentBoard
	var liveNeighbors int
	dead := b[row][col] == 0
	directions := [3]int{-1, 0, 1}

	for x := range directions {
		for y := range directions {
			if x == 0 && y == 0 {
				continue
			}
			if v := b.GetValueAt(row+y, col+x); v == 1 {
				liveNeighbors++
			}
		}

	}

	/*
	   Any live cell with two or three live neighbours survives.
	   Any dead cell with three live neighbours becomes a live cell.
	   All other live cells die in the next generation. Similarly,
	   all other dead cells stay dead.
	*/
	if !dead && liveNeighbors >= 2 && liveNeighbors <= 3 {
		return 1
	}

	if dead && liveNeighbors == 3 {
		return 1
	}

	return 0
}

// TODO - reallocate array here for efficiency
func (g *Game) clearBoard(board *Board) {
	b := *board
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			b[i][j] = 0
		}
	}
}

func (g *Game) generateRandomBoard(iter int) {
	rand.NewSource(time.Now().UnixNano())
	b := g.CurrentBoard
	for i := 0; i < iter; i++ {
		col := rand.Intn(len(b))
		row := rand.Intn(len(b))
		b[row][col] = 1
	}
}

// TODO - should this receive Game?
// InitBoard initializes a new board of x by y dimensions
func initBoard(rows, cols int) *Board {
	var b Board
	b = make([][]int, rows)
	for i := 0; i < cols; i++ {
		b[i] = make([]int, cols)
	}

	return &b
}

func (g Game) DumpBoard() {
	b := g.CurrentBoard
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
	fmt.Println("")
}
