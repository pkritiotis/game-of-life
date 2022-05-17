# Game of life

This is a sample implementation of the game of life in `go` using the `faiface/pixel` game library

![](./docs/sample.gif)

Features:
- Graphical User Interface with different colors representing the cell state and condition
  - Grey: Dead - the conditions don't allow a reproduction of a new cell in the next state
  - Light Green: Dead - a new cell will be reproduced in the next state
  - Green: Live Cell - not in danger
  - Orange: Live Cell - Overpopulated - will die in the next state
  - Red: Live Cell - Underpopulated - will die in the next state
- Customizable width, height, cell size, and framerate (see Usage)
- A new random board is populated at every run

# Running the game

To run the game run `make run` in the root directory of the repo.

## Usage
```
  -cellSize int
    	The pixel size of each cell (default 30)
  -frameRate duration
    	The framerate in milliseconds (default 500ms)
  -height float
    	The pixel size of the height of the grid (default 300)
  -width float
    	The pixel size of the width of the grid (default 600)
```