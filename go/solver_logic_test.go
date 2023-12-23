package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var emptyBoard = [][]int{
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

func copyEmptyBoard() [][]int {
	modifiedBoard := make([][]int, len(emptyBoard))
	for i := range emptyBoard {
		modifiedBoard[i] = make([]int, len(emptyBoard[i]))
		copy(modifiedBoard[i], emptyBoard[i])
	}
	return modifiedBoard

}

func TestBoxIndex1(t *testing.T) {
	board := newBoard(copyEmptyBoard())
	group := board.GetBox(1)

	assert.NotNil(t, group.Cells[Coordinate{Row: 0, Column: 3}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 0, Column: 4}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 0, Column: 5}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 1, Column: 3}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 1, Column: 4}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 1, Column: 5}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 2, Column: 3}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 2, Column: 4}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 2, Column: 5}])
}

func TestBoxIndex5(t *testing.T) {
	board := newBoard(copyEmptyBoard())
	group := board.GetBox(5)

	assert.NotNil(t, group.Cells[Coordinate{Row: 3, Column: 6}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 3, Column: 7}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 3, Column: 8}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 4, Column: 6}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 4, Column: 7}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 4, Column: 8}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 5, Column: 6}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 5, Column: 7}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 5, Column: 8}])
}

func TestBoxIndex6(t *testing.T) {
	board := newBoard(copyEmptyBoard())
	group := board.GetBox(6)

	assert.NotNil(t, group.Cells[Coordinate{Row: 6, Column: 0}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 6, Column: 1}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 6, Column: 2}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 7, Column: 0}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 7, Column: 1}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 7, Column: 2}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 8, Column: 0}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 8, Column: 1}])
	assert.NotNil(t, group.Cells[Coordinate{Row: 8, Column: 2}])
}

func TestBoardInitializesWithEverythingAPossibleSolutionIfCellValueIsZero(t *testing.T) {
	board := newBoard(copyEmptyBoard())
	middleCell := board.Cells[Coordinate{Row: 4, Column: 4}]
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, middleCell.PossibleSolutions)

}

func TestBoardInitializesWithOnlyOnePossibleSolutionIfCellValueIsNotZero(t *testing.T) {
	modifiedBoard := copyEmptyBoard()
	modifiedBoard[4][4] = 5
	board := newBoard(modifiedBoard)
	modifiedCell := board.Cells[Coordinate{Row: 4, Column: 4}]
	assert.Equal(t, []int{5}, modifiedCell.PossibleSolutions)
}

func TestEliminatesPossibleSolutionsAlreadyClaimedByAnotherCell(t *testing.T) {
	modifiedBoard := copyEmptyBoard()
	modifiedBoard[0][0] = 1
	modifiedBoard[0][1] = 2

	board := newBoard(modifiedBoard)
	row := board.GetRow(0)

	eliminateTakenSolutions(row)

	firstRowThirdColumn := board.Cells[Coordinate{Row: 0, Column: 2}]
	assert.Equal(t, []int{3, 4, 5, 6, 7, 8, 9}, firstRowThirdColumn.PossibleSolutions)
	assert.Equal(t, 0, firstRowThirdColumn.Solution)
}

func TestWhenElminatingTakenSolutionIfOnlyOneSolutionLeftThenSetSolution(t *testing.T) {
	modifiedBoard := copyEmptyBoard()
	modifiedBoard[0][0] = 1
	modifiedBoard[0][1] = 2
	modifiedBoard[0][2] = 3
	modifiedBoard[0][3] = 4
	modifiedBoard[0][4] = 5
	modifiedBoard[0][5] = 6
	modifiedBoard[0][6] = 7
	modifiedBoard[0][7] = 8

	board := newBoard(modifiedBoard)
	row := board.GetRow(0)

	eliminateTakenSolutions(row)

	firstRowNinthColumn := board.Cells[Coordinate{Row: 0, Column: 8}]
	assert.Equal(t, []int{9}, firstRowNinthColumn.PossibleSolutions)
	assert.Equal(t, 9, firstRowNinthColumn.Solution)
}
