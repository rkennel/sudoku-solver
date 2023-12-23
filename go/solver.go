package main

const BOARD_SIZE = 9

const GROUP_TYPE_ROW = "ROW"
const GROUP_TYPE_COLUMN = "COLUMN"
const GROUP_TYPE_BOX = "BOX"

func Solver(inputBoard [][]int) [][]int {

	board := newBoard(inputBoard)

	continueLoop := true

	for continueLoop {
		//eliminate possible solutions across rows
		continueLoop = evaluateRows(board)

		//eliminate possible solutions across columns
		continueLoop = continueLoop || evaluateColumns(board)

		//eliminate possible solutions across boxes
		continueLoop = continueLoop || evaluateBoxes(board)
	}

	return convertBoardToSolution(board)
}

func convertBoardToSolution(board *Board) [][]int {
	solution := make([][]int, BOARD_SIZE)
	for i := 0; i < BOARD_SIZE; i++ {
		solution[i] = make([]int, BOARD_SIZE)
		for j := 0; j < BOARD_SIZE; j++ {
			solution[i][j] = board.Cells[Coordinate{Row: i, Column: j}].Solution
		}
	}
	return solution
}

/**
 * A sudoku Board is made up of 9 rows, each containing 9 columns.
 * Each row and column contains 9 cells.
 * We will index them 0-8.
 * Row 0 is the top most row
 * Column 0 is the left most column
 * The boxes go from left to right, top to bottom.  Upper left box is 0, upper right box is 2, lower left box is 6, lower right box is 8.
 */
type Coordinate struct {
	Row    int
	Column int
}

type BoardCell struct {
	Coordinates       Coordinate
	Solution          int
	PossibleSolutions []int
}

type Board struct {
	Cells map[Coordinate]*BoardCell
}

type Group struct {
	GroupType string
	Index     int
	Cells     map[Coordinate]*BoardCell
}

func evaluateRows(board *Board) bool {
	return evaluateGroup(board.GetRow)
}

func evaluateColumns(board *Board) bool {
	return evaluateGroup(board.GetColumn)
}

func evaluateBoxes(board *Board) bool {
	return evaluateGroup(board.GetBox)
}

func evaluateGroup(getGroup func(int) *Group) bool {
	modified := false
	for i := 0; i < BOARD_SIZE; i++ {
		group := getGroup(i)
		modified = modified || eliminateTakenSolutions(group)
	}
	return modified
}

func (b *Board) GetRow(index int) *Group {
	if index < 0 || index >= BOARD_SIZE {
		return nil // return nil if index is out of bounds
	}
	group := &Group{GROUP_TYPE_ROW, index, make(map[Coordinate]*BoardCell)}
	for i := 0; i < BOARD_SIZE; i++ {
		coordinate := Coordinate{Row: index, Column: i}
		group.Cells[coordinate] = b.Cells[coordinate]
	}
	return group
}

func (b *Board) GetColumn(index int) *Group {
	if index < 0 || index >= BOARD_SIZE {
		return nil // return nil if index is out of bounds
	}

	group := &Group{GROUP_TYPE_COLUMN, index, make(map[Coordinate]*BoardCell)}
	for i := 0; i < BOARD_SIZE; i++ {
		coordinate := Coordinate{Row: i, Column: index}
		group.Cells[coordinate] = b.Cells[coordinate]
	}
	return group
}

func (b *Board) GetBox(index int) *Group {
	if index < 0 || index >= BOARD_SIZE {
		return nil // return nil if index is out of bounds
	}

	group := &Group{GROUP_TYPE_BOX, index, make(map[Coordinate]*BoardCell)}
	for i := 0; i < BOARD_SIZE; i++ {
		/*
			5=> row
		*/
		row := index/3*3 + i/3
		column := index%3*3 + i%3
		coordinate := Coordinate{Row: row, Column: column}
		group.Cells[coordinate] = b.Cells[coordinate]
	}
	return group
}

func newBoard(board [][]int) *Board {
	newBoard := &Board{make(map[Coordinate]*BoardCell)}
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			coordinate := Coordinate{Row: i, Column: j}
			solution := board[i][j]
			possibleSolutions := []int{solution}
			if solution == 0 {
				possibleSolutions = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			}
			boardCell := BoardCell{coordinate, solution, possibleSolutions}
			newBoard.Cells[coordinate] = &boardCell
		}
	}

	return newBoard
}

func eliminateTakenSolutions(row *Group) bool {
	modified := false
	for _, cell := range row.Cells {
		if cell.Solution == 0 {
			for _, possibleSolution := range cell.PossibleSolutions {
				if anotherCellHasThatSolution(row, possibleSolution) {
					cell.PossibleSolutions = removePossibleSolution(cell.PossibleSolutions, possibleSolution)
					modified = true
				}
			}
		}
	}
	if modified {
		SetSolutionIfOnlyOnePossibleSolution(row)
	}

	return modified
}

func anotherCellHasThatSolution(row *Group, solution int) bool {
	for _, cell := range row.Cells {
		if cell.Solution == solution {
			return true
		}
	}
	return false
}

func removePossibleSolution(possibleSolutions []int, solution int) []int {
	var newPossibleSolutions []int
	for _, possibleSolution := range possibleSolutions {
		if possibleSolution != solution {
			newPossibleSolutions = append(newPossibleSolutions, possibleSolution)
		}
	}
	return newPossibleSolutions
}

func SetSolutionIfOnlyOnePossibleSolution(row *Group) {
	for _, cell := range row.Cells {
		if cell.Solution == 0 && len(cell.PossibleSolutions) == 1 {
			cell.Solution = cell.PossibleSolutions[0]
		}
	}
}
