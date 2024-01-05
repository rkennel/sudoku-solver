package main

import (
	"fmt"
	"github.com/golang-collections/collections/set"
	"github.com/stretchr/testify/assert"
	"testing"
)

func emptyBoardArray() [BOARD_SIZE][BOARD_SIZE]int {
	empty := [BOARD_SIZE][BOARD_SIZE]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	return empty
}

func TestEmptyBoardPopulatesPossibleSolutions(t *testing.T) {
	empty := emptyBoardArray()
	board := InitializeBoard(empty)

	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			cell := board.Cells[row][col]
			if cell.PossibleSolutions.Len() != 9 {
				t.Error("Expected 9 possible solutions, got ", cell.PossibleSolutions.Len())
			}
		}
	}
}

func TestEmptyBoardPopulatesAllSolutionsAsZero(t *testing.T) {
	empty := emptyBoardArray()
	board := InitializeBoard(empty)

	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			cell := board.Cells[row][col]
			if cell.Solution != 0 {
				t.Error("Expected 0 solutions, got ", cell.Solution)
			}
		}
	}
}

func TestInitializeBoardEliminatesTakenSolutions(t *testing.T) {
	simpleBoard := emptyBoardArray()
	simpleBoard[0][0] = 1
	board := InitializeBoard(simpleBoard)

	assert.Equal(t, 1, board.Cells[0][0].PossibleSolutions.Len())
	assert.Equal(t, 1, board.Cells[0][0].Solution)

	assert.Equal(t, 8, board.Cells[0][1].PossibleSolutions.Len())
	for i := 2; i <= BOARD_SIZE; i++ {
		assert.True(t, board.Cells[0][1].PossibleSolutions.Has(i), "Expected %d to be a possible solutions", i)
	}
}

func TestBoxUpperLeft(t *testing.T) {

	board := InitializeBoard(emptyBoardArray())

	box := board.getBox(board.Cells[0][0])

	assert.True(t, cellArrayContains(box, board.Cells[0][0]))
	assert.True(t, cellArrayContains(box, board.Cells[0][1]))
	assert.True(t, cellArrayContains(box, board.Cells[0][2]))
	assert.True(t, cellArrayContains(box, board.Cells[1][0]))
	assert.True(t, cellArrayContains(box, board.Cells[1][1]))
	assert.True(t, cellArrayContains(box, board.Cells[1][2]))
	assert.True(t, cellArrayContains(box, board.Cells[2][0]))
	assert.True(t, cellArrayContains(box, board.Cells[2][1]))
	assert.True(t, cellArrayContains(box, board.Cells[2][2]))
}

func TestBoxCenter(t *testing.T) {

	board := InitializeBoard(emptyBoardArray())

	box := board.getBox(board.Cells[4][4])

	assert.True(t, cellArrayContains(box, board.Cells[3][3]))
	assert.True(t, cellArrayContains(box, board.Cells[3][4]))
	assert.True(t, cellArrayContains(box, board.Cells[3][5]))
	assert.True(t, cellArrayContains(box, board.Cells[4][3]))
	assert.True(t, cellArrayContains(box, board.Cells[4][4]))
	assert.True(t, cellArrayContains(box, board.Cells[4][5]))
	assert.True(t, cellArrayContains(box, board.Cells[5][3]))
	assert.True(t, cellArrayContains(box, board.Cells[5][4]))
	assert.True(t, cellArrayContains(box, board.Cells[5][5]))
}

func TestBoxLowerRight(t *testing.T) {

	board := InitializeBoard(emptyBoardArray())

	box := board.getBox(board.Cells[8][8])

	assert.True(t, cellArrayContains(box, board.Cells[6][6]))
	assert.True(t, cellArrayContains(box, board.Cells[6][7]))
	assert.True(t, cellArrayContains(box, board.Cells[6][8]))
	assert.True(t, cellArrayContains(box, board.Cells[7][6]))
	assert.True(t, cellArrayContains(box, board.Cells[7][7]))
	assert.True(t, cellArrayContains(box, board.Cells[7][8]))
	assert.True(t, cellArrayContains(box, board.Cells[8][6]))
	assert.True(t, cellArrayContains(box, board.Cells[8][7]))
	assert.True(t, cellArrayContains(box, board.Cells[8][8]))
}

func cellArrayContains(box [BOARD_SIZE]*Cell, cell *Cell) bool {
	for _, c := range box {
		if c.Row == cell.Row && c.Col == cell.Col {
			return true
		}
	}

	return false
}

func TestGenerateUniqueCombinationOfCellsSize2(t *testing.T) {
	cells := []*Cell{
		&Cell{0, 0, 0, set.New()},
		&Cell{0, 1, 0, set.New()},
		&Cell{0, 2, 0, set.New()},
		&Cell{0, 3, 0, set.New()},
		&Cell{0, 4, 0, set.New()},
		&Cell{0, 5, 0, set.New()},
		&Cell{0, 6, 0, set.New()},
		&Cell{0, 7, 0, set.New()},
		&Cell{0, 8, 0, set.New()},
	}

	uniqueCombinations := generateUniqueCombinationsOfCells(cells, 2)

	assert.Equal(t, 36, len(uniqueCombinations))
}

func TestGenerateUniqueCombinationOfCellsSize3(t *testing.T) {
	cells := []*Cell{
		&Cell{0, 0, 0, set.New()},
		&Cell{0, 1, 0, set.New()},
		&Cell{0, 2, 0, set.New()},
		&Cell{0, 3, 0, set.New()},
		&Cell{0, 4, 0, set.New()},
		&Cell{0, 5, 0, set.New()},
		&Cell{0, 6, 0, set.New()},
		&Cell{0, 7, 0, set.New()},
		&Cell{0, 8, 0, set.New()},
	}

	uniqueCombinations := generateUniqueCombinationsOfCells(cells, 3)

	assert.Equal(t, 84, len(uniqueCombinations))
}

func TestGenerateUniqueCombinationOfCellsSize4(t *testing.T) {
	cells := []*Cell{
		&Cell{0, 0, 0, set.New()},
		&Cell{0, 1, 0, set.New()},
		&Cell{0, 2, 0, set.New()},
		&Cell{0, 3, 0, set.New()},
		&Cell{0, 4, 0, set.New()},
		&Cell{0, 5, 0, set.New()},
		&Cell{0, 6, 0, set.New()},
		&Cell{0, 7, 0, set.New()},
		&Cell{0, 8, 0, set.New()},
	}

	uniqueCombinations := generateUniqueCombinationsOfCells(cells, 4)

	assert.Equal(t, 126, len(uniqueCombinations))
}

func TestEliminateExclusivePairSolution(t *testing.T) {
	cells := [BOARD_SIZE]*Cell{
		&Cell{0, 0, 0, set.New(1, 2, 3)},
		&Cell{0, 1, 0, set.New(2, 3)},
		&Cell{0, 2, 0, set.New(2, 3)},
		&Cell{0, 3, 4, set.New(4)},
		&Cell{0, 4, 5, set.New(5)},
		&Cell{0, 5, 6, set.New(6)},
		&Cell{0, 6, 7, set.New(7)},
		&Cell{0, 7, 8, set.New(8)},
		&Cell{0, 8, 9, set.New(9)},
	}

	modified := eliminateExclusiveSolutions(cells, 2)
	assert.True(t, modified)
	assert.Equal(t, 1, cells[0].PossibleSolutions.Len())
}

func TestMyUnderstanding(t *testing.T) {
	nums := []int{1, 2, 3}
	fmt.Println(nums)
	modifyNums(&nums)
	fmt.Println(nums)
}

func modifyNums(nums *[]int) {
	*nums = append(*nums, 4)
}
