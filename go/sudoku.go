package main

import "github.com/golang-collections/collections/set"

const BOARD_SIZE = 9

type Cell struct {
	Row               int
	Col               int
	Solution          int
	PossibleSolutions *set.Set
}

func (c *Cell) PossibleSolutionsAsArray() []int {
	var possibleSolutions []int
	c.PossibleSolutions.Do(func(solution interface{}) {
		possibleSolutions = append(possibleSolutions, solution.(int))
	})

	return possibleSolutions
}

type Board struct {
	Cells [BOARD_SIZE][BOARD_SIZE]*Cell
}

func (board *Board) solutions() [BOARD_SIZE][BOARD_SIZE]int {
	var solutions [BOARD_SIZE][BOARD_SIZE]int

	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			solutions[i][j] = board.Cells[i][j].Solution
		}
	}

	return solutions
}

func (board *Board) calculateTakenSolutions(cell *Cell) *set.Set {

	row := board.getRow(cell)
	column := board.getColumn(cell)
	box := board.getBox(cell)

	takenSolutions := set.New()
	addKnownSolutions(takenSolutions, row)
	addKnownSolutions(takenSolutions, column)
	addKnownSolutions(takenSolutions, box)

	return takenSolutions
}

func addKnownSolutions(solutions *set.Set, cells [BOARD_SIZE]*Cell) {
	for _, cell := range cells {
		if cell.Solution != 0 {
			solutions.Insert(cell.Solution)
		}
	}
}

func (board *Board) getRow(cell *Cell) [BOARD_SIZE]*Cell {
	return board.Cells[cell.Row]
}

func (board *Board) getColumn(cell *Cell) [BOARD_SIZE]*Cell {
	var column [BOARD_SIZE]*Cell
	for row := 0; row < BOARD_SIZE; row++ {
		column[row] = board.Cells[row][cell.Col]
	}
	return column
}

func (board *Board) getBox(cell *Cell) [BOARD_SIZE]*Cell {
	rowStart := cell.Row / 3 * 3
	colStart := cell.Col / 3 * 3

	var box [BOARD_SIZE]*Cell
	boxIndex := 0
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			box[boxIndex] = board.Cells[rowStart+row][colStart+col]
			boxIndex++
		}
	}

	return box
}

func (board *Board) updatePossibleSolutions() {
	continueProcessing := true

	for continueProcessing {
		removeTakenSolutions(board)
		continueProcessing = setSolutionWhereOnlyOnePossibleSolutionsRemains(board)
	}
}

func removeTakenSolutions(board *Board) {
	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			cell := board.Cells[row][col]
			takenSolutions := board.calculateTakenSolutions(cell)
			cell.PossibleSolutions = cell.PossibleSolutions.Difference(takenSolutions)
			if cell.Solution != 0 {
				cell.PossibleSolutions.Insert(cell.Solution)
			}
		}
	}
}

func setSolutionWhereOnlyOnePossibleSolutionsRemains(board *Board) bool {
	continueProcessing := false

	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			cell := board.Cells[row][col]
			if cell.Solution == 0 && cell.PossibleSolutions.Len() == 1 {
				cell.Solution = cell.PossibleSolutionsAsArray()[0]
				continueProcessing = true
			}
		}
	}

	return continueProcessing
}

func InitializeBoard(inputBoard [BOARD_SIZE][BOARD_SIZE]int) Board {
	var board Board
	board.Cells = [BOARD_SIZE][BOARD_SIZE]*Cell{}

	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			board.Cells[i][j] = &Cell{i, j, inputBoard[i][j], set.New()}
			if board.Cells[i][j].Solution == 0 {
				for k := 1; k <= BOARD_SIZE; k++ {
					board.Cells[i][j].PossibleSolutions.Insert(k)
				}
			}
		}
	}

	board.updatePossibleSolutions()

	return board
}
