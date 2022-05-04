package main

import (
	"flag"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/pkritiotis/game-of-life/gameoflife"
	"golang.org/x/image/colornames"
	"math/rand"
	"time"
)

var (
	size       *int
	windowSize *float64
	frameRate  *time.Duration
)

func init() {
	rand.Seed(time.Now().UnixNano())
	size = flag.Int("size", 20, "The size of each cell")
	windowSize = flag.Float64("windowSize", 800, "The pixel size of one side of the grid")
	frameRate = flag.Duration("frameRate", 200*time.Millisecond, "The framerate in milliseconds")
	flag.Parse()
}

func run() {

	cfg := pixelgl.WindowConfig{
		Title:  "Game of Life",
		Bounds: pixel.R(0, 0, *windowSize, *windowSize),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.White)

	// since the game board is square, rows and cols will be the same
	rows := int(*windowSize) / *size

	gridDraw := imdraw.New(nil)
	game := gameoflife.New(rows, rows)
	tick := time.Tick(*frameRate)
	for !win.Closed() {
		// game loop
		select {
		case <-tick:
			gridDraw.Clear()
			printBoard(game.Board, gridDraw, *size)
			gridDraw.Draw(win)
			game.Next()
		}
		win.Update()
	}
}

func printBoard(board [][]gameoflife.CellState, imd *imdraw.IMDraw, cellSize int) {
	for i := range board {
		for j := range board[i] {
			switch board[i][j] {
			case gameoflife.Dead_WillRemainDead:
				imd.Color = colornames.White
				break
			case gameoflife.Dead_WillBeBorn:
				imd.Color = colornames.Green
				break
			case gameoflife.Alive_UnderPopulated:
				imd.Color = colornames.Yellow
				break
			case gameoflife.Alive_OverPopulated:
				imd.Color = colornames.Red
				break
			case gameoflife.Alive_WillSurvive:
				imd.Color = colornames.Blue
			}

			imd.Push(
				pixel.V(float64(i*cellSize), float64(j*cellSize)),
				pixel.V(float64(i*cellSize+cellSize), float64(j*cellSize+cellSize)),
			)
			imd.Rectangle(0)
		}
	}
}

func main() {
	pixelgl.Run(run)
}
