package gameoflife

import (
	"time"
)
import "math/rand"

func init() {
	rand.Seed(time.Now().UnixNano())
}

type CellState int

const (
	Dead_WillRemainDead CellState = iota
	Dead_WillBeBorn
	Alive_WillSurvive
	Alive_OverPopulated
	Alive_UnderPopulated
)

type GameOfLife struct {
	Board   [][]CellState
	rows    int
	columns int
}

func New(rows, columns int) GameOfLife {
	board := getBoardWithCellStates(newRandomBoard(rows, columns))

	return GameOfLife{
		Board:   board,
		rows:    rows,
		columns: columns,
	}
}

func getBoardWithCellStates(board [][]bool) [][]CellState {
	newStateBoard := newCellStateBoard(len(board), len(board[0]))
	for i := range board {
		for j := range board[i] {
			activeNeighbors := getActiveNeighborsBool(i, j, board)
			if board[i][j] {
				if activeNeighbors == 2 || activeNeighbors == 3 {
					newStateBoard[i][j] = Alive_WillSurvive
					continue
				}
				if activeNeighbors > 3 {
					newStateBoard[i][j] = Alive_OverPopulated
					continue
				}
				newStateBoard[i][j] = Alive_UnderPopulated
			} else {
				if activeNeighbors == 3 {
					newStateBoard[i][j] = Dead_WillBeBorn
					continue
				}
				newStateBoard[i][j] = Dead_WillRemainDead
			}
		}
	}
	return newStateBoard
}

func randBool() bool {
	return rand.Intn(2) == 1
}

func newRandomBoard(rows, columns int) [][]bool {
	board := make([][]bool, rows)
	for i := range board {
		board[i] = make([]bool, columns)
		for j := range board[i] {
			board[i][j] = randBool()
		}
	}
	return board
}

func newCellStateBoard(rows, columns int) [][]CellState {
	board := make([][]CellState, rows)
	for i := range board {
		board[i] = make([]CellState, columns)
	}
	return board
}

func newBoard(rows, columns int) [][]bool {
	board := make([][]bool, rows)
	for i := range board {
		board[i] = make([]bool, columns)
	}
	return board
}

func (gol *GameOfLife) Next() {
	newStateBoard := newBoard(gol.rows, gol.columns)
	for i := range gol.Board {
		for j := range gol.Board[i] {
			activeNeighbors := getActiveNeighbors(i, j, gol.Board)
			if gol.Board[i][j] > 1 {
				newStateBoard[i][j] = activeNeighbors == 2 || activeNeighbors == 3
			} else {
				newStateBoard[i][j] = activeNeighbors == 3
			}
		}
	}
	gol.Board = getBoardWithCellStates(newStateBoard)
}

func getActiveNeighborsBool(i int, j int, board [][]bool) int {
	activeNeighbors := 0
	for x := -1; x <= 1; x++ {
		rows := len(board)
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			columns := len(board[0])
			if x+i < 0 || x+i >= rows ||
				y+j < 0 || y+j >= columns {
				continue
			}
			if board[x+i][y+j] {
				activeNeighbors++
			}
		}
	}
	return activeNeighbors
}

func getActiveNeighbors(i int, j int, board [][]CellState) int {
	activeNeighbors := 0
	for x := -1; x <= 1; x++ {
		rows := len(board)
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			columns := len(board[0])
			if x+i < 0 || x+i >= rows ||
				y+j < 0 || y+j >= columns {
				continue
			}
			if board[x+i][y+j] > 1 {
				activeNeighbors++
			}
		}
	}
	return activeNeighbors
}

func (gol GameOfLife) IsDead() bool {
	return false
}
