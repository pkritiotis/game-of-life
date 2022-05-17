package gameoflife

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//CellState represents the state of the cell
type CellState int

const (
	//Unknown represents an uninitialized state
	Unknown CellState = iota
	//Still represents the state of the Cell that will remain in its current health state (dead->dead or alive->alive)
	Still
	//Reproduction represents the state of a dead Cell that will become alive in the next step
	Reproduction
	//OverPopulated represents the state of a live Cell that will die in the next step because of overpopulation
	OverPopulated
	//UnderPopulated represents the state of a live Cell that will die in the next step because of underpopulation
	UnderPopulated
)

//Cell encapsulates the liveness and state of a cell
type Cell struct {
	IsAlive bool
	State   CellState
}

//GaemOfLife contains the grid of the game of life
type GameOfLife struct {
	Grid    [][]Cell
	rows    int
	columns int
}

//New GameOfLife constructor
func New(rows, columns int) GameOfLife {
	gol := GameOfLife{
		rows:    rows,
		columns: columns,
	}
	gol.Grid = gol.populateCellStates(newRandomGrid(rows, columns))
	return gol
}

func newRandomGrid(rows, columns int) [][]Cell {
	grid := make([][]Cell, rows)
	for i := range grid {
		grid[i] = make([]Cell, columns)
		for j := range grid[i] {
			grid[i][j] = Cell{
				IsAlive: rand.Intn(2) == 1,
				State:   Unknown,
			}
		}
	}
	return grid
}

func (gol GameOfLife) populateCellStates(grid [][]Cell) [][]Cell {
	populatedGrid := make([][]Cell, len(grid))
	for i := range grid {
		populatedGrid[i] = make([]Cell, len(grid[i]))
		for j := range grid[i] {
			activeNeighbors := gol.getActiveNeighbors(i, j, grid)
			cell := grid[i][j]
			if cell.IsAlive {
				if activeNeighbors == 2 || activeNeighbors == 3 {
					cell.State = Still
				} else if activeNeighbors > 3 {
					cell.State = OverPopulated
				} else {
					cell.State = UnderPopulated
				}
			} else {
				if activeNeighbors == 3 {
					cell.State = Reproduction
				} else {
					cell.State = Still
				}
			}
			populatedGrid[i][j] = cell
		}
	}
	return populatedGrid
}

func newGrid(rows, columns int) [][]Cell {
	grid := make([][]Cell, rows)
	for i := range grid {
		grid[i] = make([]Cell, columns)
	}
	return grid
}

//Next advances the GameOfLife to the next step/state
func (gol *GameOfLife) Next() {
	newStateGrid := newGrid(gol.rows, gol.columns)
	for i := range gol.Grid {
		for j := range gol.Grid[i] {
			activeNeighbors := gol.getActiveNeighbors(i, j, gol.Grid)
			if gol.Grid[i][j].IsAlive {
				newStateGrid[i][j].IsAlive = activeNeighbors == 2 || activeNeighbors == 3
			} else {
				newStateGrid[i][j].IsAlive = activeNeighbors == 3
			}
			newStateGrid[i][j].State = Unknown
		}
	}
	gol.Grid = gol.populateCellStates(newStateGrid)

}

func (gol GameOfLife) getActiveNeighbors(i int, j int, grid [][]Cell) int {
	activeNeighbors := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if x+i < 0 || x+i >= gol.rows ||
				y+j < 0 || y+j >= gol.columns {
				continue
			}
			if grid[x+i][y+j].IsAlive {
				activeNeighbors++
			}
		}
	}
	return activeNeighbors
}
