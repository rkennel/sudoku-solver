package main

import (
	"errors"
	"fmt"
	"github.com/golang-collections/collections/set"
)

var BOARD_SIZE = 9

type Coordinate struct {
	Row    int
	Column int
}

type Cell struct {
	Solution          int
	PossibleSolutions *set.Set
	Coordinate        Coordinate
}

type SudokuBoard struct {
	Cells map[Coordinate]*Cell
}

func (board *SudokuBoard) GetCell(row int, column int) *Cell {
	return board.Cells[Coordinate{row, column}]
}

func (board *SudokuBoard) AsArray() [][]int {
	rows := [][]int{}
	for i := 0; i < BOARD_SIZE; i++ {
		rownum := i + 1
		columns := []int{}
		for j := 0; j < BOARD_SIZE; j++ {
			colnum := j + 1
			columns = append(columns, board.Cells[Coordinate{rownum, colnum}].Solution)
		}
		rows = append(rows, columns)
	}
	return rows
}

func NewBoard(inputBoard [][]int) (*SudokuBoard, error) {
	if !verifyBoardSize(inputBoard) {
		return nil, errors.New(fmt.Sprintf("Board is not the proper size [%v,%v]", BOARD_SIZE, BOARD_SIZE))
	}

	board := SudokuBoard{
		Cells: make(map[Coordinate]*Cell),
	}

	for i, row := range inputBoard {
		rownum := i + 1
		for j, value := range row {
			colnum := j + 1
			var possibleSolutions *set.Set
			if value == 0 {
				possibleSolutions = set.New(1, 2, 3, 4, 5, 6, 7, 8, 9)
			} else {
				possibleSolutions = set.New(value)
			}
			board.Cells[Coordinate{rownum, colnum}] = &Cell{
				Solution:          value,
				PossibleSolutions: possibleSolutions,
				Coordinate:        Coordinate{rownum, colnum},
			}
		}
	}

	//eliminatePossibleSolutions(&board)

	return &board, nil
}

func verifyBoardSize(board [][]int) bool {
	if len(board) != BOARD_SIZE {
		return false
	}

	for _, row := range board {
		if len(row) != BOARD_SIZE {
			return false
		}
	}

	return true
}

func (board *SudokuBoard) Solve() bool {
	return false
}

//
//func eliminatePossibleSolutions(board *SudokuBoard) {
//	for row := 1; row <= BOARD_SIZE; row++ {
//		for column := 1; column <= BOARD_SIZE; column++ {
//			cell := board.GetCell(row, column)
//			if cell.Solution != 0 {
//				eliminatePossibleSolutionsInRow(board, cell)
//			}
//		}
//	}
//}
//
//func eliminatePossibleSolutionsInRow(board *SudokuBoard, cell *Cell) {
//	for column := 1; column <= BOARD_SIZE; column++ {
//		if column != cell.Coordinate.Column {
//			takenSolution := board.GetCell(cell.Coordinate.Row, column).Solution
//			if takenSolution != 0 {
//				cell.PossibleSolutions = removePossibleSolution(cell.PossibleSolutions, takenSolution)
//			}
//		}
//	}
//}
//
//func removePossibleSolution(possibleSolutions []int, takenSolution int) []int {
//	newPossibleSolutions := []int{}
//	for _, possibleSolution := range possibleSolutions {
//		if possibleSolution != takenSolution {
//			newPossibleSolutions = append(newPossibleSolutions, possibleSolution)
//		}
//	}
//	return newPossibleSolutions
//}
