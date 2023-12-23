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

func copyEmptyBoard() [][]int {
	modifiedBoard := make([][]int, len(emptyBoard))
	for i := range emptyBoard {
		modifiedBoard[i] = make([]int, len(emptyBoard[i]))
		copy(modifiedBoard[i], emptyBoard[i])
	}
	return modifiedBoard

}
