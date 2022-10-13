package main

import "fmt"

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
	print(board)
}

func solveSudoku(board [][]byte) {
	solveSudokuLine(0, 0, board)
}

func solveSudokuLine(row, col int, board [][]byte) bool {
	if row >= len(board) {
		fmt.Println("fin")
		return true
	}

	for board[row][col] != '.' {
		if col == len(board)-1 {
			row += 1
			col = 0
		} else {
			col += 1
		}
		if row >= len(board) {
			fmt.Println("fin")
			return true
		}
	}

	if row >= len(board) {
		fmt.Println("fin")
		return true
	}

	min, r, c := countMinPossibleValues(board)
	fmt.Printf("%d %d: %d\n", r, c, min)
	//try values
	possibleVals := getPossibleVals(r, c, board)
	if len(possibleVals) == 0 {
		fmt.Println("error!")
		return false
	}
	isValid := false
	for k := range possibleVals {
		board[r][c] = k
		isValid = solveSudokuLine(0, 0, board)
		if isValid {
			return true
		} else {
			board[r][c] = '.'
		}
	}
	return isValid
}

func getPossibleVals(row, col int, board [][]byte) map[byte]struct{} {
	if board[row][col] != '.' {
		panic("not supported")
	}
	possibleVals := map[byte]struct{}{'1': {}, '2': {}, '3': {}, '4': {}, '5': {}, '6': {}, '7': {}, '8': {}, '9': {}}
	for i := 0; i < len(board); i++ {
		delete(possibleVals, board[row][i])
		delete(possibleVals, board[i][col])
	}
	rowOffset := (row / 3) * 3
	colOffset := (col / 3) * 3
	for i := rowOffset; i < rowOffset+3; i++ {
		for j := colOffset; j < colOffset+3; j++ {
			delete(possibleVals, board[i][j])
		}
	}
	return possibleVals
}

func print(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func countMinPossibleValues(board [][]byte) (int, int, int) {
	min, row, col := 9, 0, 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == '.' {
				//try values
				possibleVals := getPossibleVals(i, j, board)
				if min > len(possibleVals) {
					min, row, col = len(possibleVals), i, j

				}
			}
		}
	}
	return min, row, col
}
