package main

func Solver(inputBoard [BOARD_SIZE][BOARD_SIZE]int) [BOARD_SIZE][BOARD_SIZE]int {
	board := InitializeBoard(inputBoard)

	return board.solutions()
}

func eliminateExclusiveSolutions(cells [9]*Cell, setSize int) bool {
	continueProcessing := false

	cellsWithoutSolutions := filterCellsThatDoNotHaveASolution(cells)
	uniqueCombinationsOfCells := generateUniqueCombinationsOfCells(cellsWithoutSolutions, setSize)

	return continueProcessing
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
