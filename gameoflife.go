package main

import "fmt"
import "time"
import "math/rand"

func main() {
	gol := New(Settings{
		x:     10,
		y:     10,
		speed: 0,
	})
	gol.Start()
}

type Settings struct {
	x     int
	y     int
	speed int
}

type GameOfLife struct {
	board    [][]bool
	settings Settings
}

func New(settings Settings) GameOfLife {
	board := NewRandomBoard(settings)
	return GameOfLife{settings: settings, board: board}
}

func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func NewRandomBoard(settings Settings) [][]bool {
	board := make([][]bool, settings.x)
	for i := range board {
		board[i] = make([]bool, settings.y)
		for j := range board[i] {
			board[i][j] = RandBool()
		}
	}
	return board
}

func NewBoard(settings Settings) [][]bool {
	board := make([][]bool, settings.x)
	for i := range board {
		board[i] = make([]bool, settings.y)
	}
	return board
}

func (gol GameOfLife) Start() error {
	for !gol.IsDead() {
		gol.NextState()
		gol.Print()
		fmt.Println("----------")
		time.Sleep(time.Millisecond * 500)
	}
	return nil
}

func (gol GameOfLife) Print() {
	for i := range gol.board {
		for j := range gol.board[i] {
			if gol.board[i][j] {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}

func (gol *GameOfLife) NextState() {
	newStateBoard := NewBoard(gol.settings)
	for i := range gol.board {
		for j := range gol.board[i] {
			activeNeighbors := gol.GetActiveNeighbors(i, j)
			if gol.board[i][j] {
				newStateBoard[i][j] = activeNeighbors == 2 || activeNeighbors == 3
			} else {
				newStateBoard[i][j] = activeNeighbors == 3
			}
		}
	}
	gol.board = newStateBoard
}

func (gol GameOfLife) GetActiveNeighbors(i int, j int) int {
	activeNeighbors := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if x+i < 0 || x+i >= gol.settings.x ||
				y+j < 0 || y+j >= gol.settings.y {
				continue
			}
			if gol.board[x+i][y+j] {
				activeNeighbors++
			}
		}
	}
	return activeNeighbors
}

func (gol GameOfLife) IsDead() bool {
	return false
}
