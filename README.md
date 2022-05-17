# Game of life

This is a sample implementation of the game of life in go using `faiface/pixel` 2D game library (https://github.com/faiface/pixel).

![](./docs/sample.gif)

Features:
- Graphical User Interface with different colors representing the cell state and condition
  - ![#f5f5f5ff](http://via.placeholder.com/15/f5f5f5ff/000000?text=+): Dead - the conditions don't allow a reproduction of a new cell in the next state
  - ![#bddfbdff](http://via.placeholder.com/15/bddfbddf/000000?text=+): Dead - a new cell will be reproduced in the next state
  - ![#228b22ff](http://via.placeholder.com/15/228b22ff/000000?text=+): Live Cell - not in danger
  - ![#ff4500ff](http://via.placeholder.com/15/ff4500ff/000000?text=+): Live Cell - Overpopulated - will die in the next state
  - ![#ff8c00ff](http://via.placeholder.com/15/ff8c00ff/000000?text=+): Live Cell - Underpopulated - will die in the next state
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
