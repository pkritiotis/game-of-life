package main

import (
	"flag"
	"fmt"
	"image/color"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/pkritiotis/game-of-life/gameoflife"
	"golang.org/x/image/colornames"
)

var (
	cellSize     *int
	windowWidth  *float64
	windowHeight *float64
	frameRate    *time.Duration
)

func init() {
	cellSize = flag.Int("cellSize", 30, "The pixel size of each cell")
	windowWidth = flag.Float64("width", 600, "The pixel size of the width of the grid")
	windowHeight = flag.Float64("height", 500, "The pixel size of the height of the grid")
	frameRate = flag.Duration("frameRate", 300*time.Millisecond, "The framerate in milliseconds")
	flag.Parse()
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Game of Life",
		Bounds: pixel.R(0, 0, *windowWidth, *windowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		fmt.Printf("Could not create a new window: %v", err)
		os.Exit(1)
	}
	win.Clear(colornames.White)

	rows := int(*windowWidth) / *cellSize
	columns := int(*windowHeight) / *cellSize

	gridDraw := imdraw.New(nil)
	game := gameoflife.New(rows, columns)
	tick := time.Tick(*frameRate)
	for !win.Closed() {
		select {
		case <-tick:
			gridDraw.Clear()
			printBoard(game.Grid, gridDraw, *cellSize)
			gridDraw.Draw(win)
			game.Next()
		}
		win.Update()
	}
}

func printBoard(board [][]gameoflife.Cell, imd *imdraw.IMDraw, cellSize int) {
	for i := range board {
		for j := range board[i] {
			switch board[i][j].State {
			case gameoflife.Still:
				if !board[i][j].IsAlive {
					imd.Color = colornames.Whitesmoke
				} else {
					imd.Color = colornames.Forestgreen
				}
				break
			case gameoflife.Reproduction:
				imd.Color = color.RGBA{0xbd, 0xdf, 0xbd, 0xdf}
				break
			case gameoflife.UnderPopulated:
				imd.Color = colornames.Orangered
				break
			case gameoflife.OverPopulated:
				imd.Color = colornames.Darkorange
				break
			}

			imd.Push(
				pixel.V(float64(i*cellSize)+2, float64(j*cellSize+2)),
				pixel.V(float64(i*cellSize+cellSize)-2, float64(j*cellSize+cellSize-2)),
			)
			// imd.Circle(float64(cellSize/2), 0)
			imd.Rectangle(0)
		}
	}
}

func main() {
	pixelgl.Run(run)
}
