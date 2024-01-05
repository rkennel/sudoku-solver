package main

import (
	"github.com/golang-collections/collections/set"
)

func Solver(inputBoard [BOARD_SIZE][BOARD_SIZE]int) [BOARD_SIZE][BOARD_SIZE]int {
	board := InitializeBoard(inputBoard)

	keepProcessing := true

	for keepProcessing {
		modified := eliminateExclusiveSolutionsForRows(&board)
		if modified {
			board.updatePossibleSolutions()
			continue
		}

		keepProcessing = eliminateExclusiveSolutionsForColumns(&board)
		if modified {
			board.updatePossibleSolutions()
			continue
		}

		keepProcessing = eliminateExclusiveSolutionsForBoxes(&board)
		if modified {
			board.updatePossibleSolutions()
			continue
		} else {
			keepProcessing = false
		}
	}

	board.debugPossibleSolutionsCSV()

	return board.solutions()
}

func eliminateExclusiveSolutionsForRows(board *Board) bool {
	return eliminateExclusiveSolutionsFromGroup(board, func(board *Board, index int) [BOARD_SIZE]*Cell {
		return board.getRow(board.Cells[index][0])
	})
}

func eliminateExclusiveSolutionsForColumns(board *Board) bool {
	return eliminateExclusiveSolutionsFromGroup(board, func(board *Board, index int) [BOARD_SIZE]*Cell {
		return board.getRow(board.Cells[0][index])
	})
}

func eliminateExclusiveSolutionsForBoxes(board *Board) bool {
	return eliminateExclusiveSolutionsFromGroup(board, func(board *Board, index int) [BOARD_SIZE]*Cell {
		row := index / 3
		column := index % 3
		return board.getRow(board.Cells[row][column])
	})
}

func eliminateExclusiveSolutionsFromGroup(board *Board, getGroup func(board *Board, index int) [BOARD_SIZE]*Cell) bool {
	modified := false

	for setSize := 2; setSize < BOARD_SIZE; setSize++ {
		for groupIndex := 0; groupIndex < BOARD_SIZE; groupIndex++ {
			modified = eliminateExclusiveSolutions(getGroup(board, groupIndex), setSize)

			if modified {
				return modified
			}
		}
	}

	return modified
}

func eliminateExclusiveSolutions(cells [9]*Cell, setSize int) bool {

	cellsWithoutSolutions := filterCellsThatDoNotHaveASolution(cells)

	if len(cellsWithoutSolutions) <= setSize {
		return false
	}

	modified := false
	uniqueCombinationsOfCells := generateUniqueCombinationsOfCells(cellsWithoutSolutions, setSize)

	for _, combination := range uniqueCombinationsOfCells {
		possibleSolutionsForCombination := calculatePossibleSolutionsForCombination(combination)
		if possibleSolutionsForCombination.Len() == setSize {
			modified = removePossibleSolutionsFromOtherCells(cells, possibleSolutionsForCombination, combination)
		}
		if modified {
			break
		}
	}

	return modified
}

func filterCellsThatDoNotHaveASolution(cells [9]*Cell) []*Cell {
	cellsWithoutSolutions := []*Cell{}

	for _, cell := range cells {
		if cell.Solution == 0 {
			cellsWithoutSolutions = append(cellsWithoutSolutions, cell)
		}
	}

	return cellsWithoutSolutions
}

func generateUniqueCombinationsOfCells(cellsWithoutSolutions []*Cell, size int) [][]*Cell {
	uniqueCombinations := [][]*Cell{}

	// Create a helper function to generate combinations recursively
	var generateCombinations func(start int, combination []*Cell)
	generateCombinations = func(start int, uniqueCombination []*Cell) {
		// If the combination has reached the desired size, add it to the combinations slice
		if len(uniqueCombination) == size {
			uniqueCombinations = append(uniqueCombinations, uniqueCombination)
			return
		}

		// Iterate over the remaining elements in the nums array
		for i := start; i < len(cellsWithoutSolutions); i++ {
			// Create a new combination by adding the current element to the existing combination
			newCombination := make([]*Cell, len(uniqueCombination))
			copy(newCombination, uniqueCombination)
			newCombination = append(newCombination, cellsWithoutSolutions[i])

			// Recursively generate combinations starting from the next element
			generateCombinations(i+1, newCombination)
		}
	}

	// Generate combinations starting from the first element
	generateCombinations(0, []*Cell{})

	// Return the combinations slice
	return uniqueCombinations

}

func calculatePossibleSolutionsForCombination(combination []*Cell) *set.Set {
	possibleSolutions := set.New()

	for _, cell := range combination {
		possibleSolutions = possibleSolutions.Union(cell.PossibleSolutions)
	}

	return possibleSolutions
}

func removePossibleSolutionsFromOtherCells(cells [9]*Cell, possibleSolutionsForCombination *set.Set, combination []*Cell) bool {
	modified := false
	for _, cell := range cells {
		if cell.Solution == 0 && !containsCell(combination, cell) {
			before := cell.PossibleSolutions.Len()
			cell.PossibleSolutions = cell.PossibleSolutions.Difference(possibleSolutionsForCombination)
			after := cell.PossibleSolutions.Len()
			if before != after {
				modified = true
			}
		}
	}

	return modified
}

func containsCell(combination []*Cell, cell *Cell) bool {
	for _, c := range combination {
		if c == cell {
			return true
		}
	}

	return false
}
