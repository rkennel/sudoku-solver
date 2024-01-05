package main

import (
	"github.com/golang-collections/collections/set"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoardInitializationIfArrayIsNotProperBoardSize(t *testing.T) {
	_, err := NewBoard([][]int{{1, 2, 3}})
	assert.Error(t, err)
}

func TestBoardInitializationAndToArray(t *testing.T) {
	inputBoardArray := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{9, 1, 2, 3, 4, 5, 6, 7, 8},
		{8, 9, 1, 2, 3, 4, 5, 6, 7},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9, 1, 2, 3, 4},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{2, 3, 4, 5, 6, 7, 8, 9, 1},
	}

	board, err := NewBoard(inputBoardArray)
	assert.NoErrorf(t, err, "Error creating board: %v", err)
	assert.Equal(t, inputBoardArray, board.AsArray())
}

func TestBoardInitializationIdentifiesPossibleSolutions(t *testing.T) {
	inputBoardArray := [][]int{
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

	board, _ := NewBoard(inputBoardArray)
	assert.Equal(t, board.GetCell(1, 1).PossibleSolutions, set.New(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

func TestBoardInitializationOnlyOnePossibleSolutionIfTheSolutionIsKnown(t *testing.T) {
	inputBoardArray := [][]int{
		{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	board, _ := NewBoard(inputBoardArray)
	assert.Equal(t, board.GetCell(1, 1).PossibleSolutions, set.New(8))
}

//func TestBoardInitializationEliminatesPossibleSolutionsInRow(t *testing.T) {
//	inputBoardArray := [][]int{
//		{0, 7, 8, 0, 0, 0, 0, 0, 0},
//		{0, 0, 0, 0, 0, 0, 0, 0, 0},
//		{0, 0, 0, 0, 0, 0, 0, 0, 0},
//		{0, 0, 0, 0, 0, 0, 0, 0, 0},
//		{0, 0, 0, 0, 0, 0, 0, 0, 0},
//		{0, 0, 0, 0, 0, 0, 0, 0, 0},
//		{0, 0, 0, 0, 0, 0, 0, 0, 0},
//		{0, 0, 0, 0, 0, 0, 0, 0, 0},
//		{0, 0, 0, 0, 0, 0, 0, 0, 0},
//	}
//
//	board, _ := NewBoard(inputBoardArray)
//	assert.Equal(t, board.GetCell(1, 1).PossibleSolutions, []int{1, 2, 3, 4, 5, 6, 9})
//}
