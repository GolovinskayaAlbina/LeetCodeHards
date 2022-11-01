package hard

import (
	"errors"
	"strings"
)

func solveNQueens(n int) [][]string {
	var res [][]string
	if boards := existsNQueensArrangement(0, createEmptyBoard(n)); len(boards) > 0 {
		for _, board := range boards {
			for idx := range board {
				board[idx] = strings.Replace(board[idx], "x", ".", -1)
			}
		}
		res = append(res, boards...)
	}
	return res
}

func totalNQueens(n int) int {
	res := 0
	if boards := existsNQueensArrangement(0, createEmptyBoard(n)); len(boards) > 0 {
		res += len(boards)
	}
	return res
}

func existsNQueensArrangement(row int, board []string) [][]string {
	if row == len(board) {
		return [][]string{board}
	}
	var res [][]string
	for col := 0; col < len(board); col++ {
		if board[row][col] == '.' {
			clone := deepClone(board)
			clone[row] = createNQueensRow(col, len(clone), "x")
			if err := setCellsUnderAttack(row, col, clone); err != nil {
				continue
			}
			boards := existsNQueensArrangement(row+1, clone)
			res = append(res, boards...)
		}
	}
	return res
}

func deepClone(board []string) []string {
	clone := []string{}
	for idx := range board {
		clone = append(clone, board[idx])
	}
	return clone
}

func createEmptyBoard(n int) []string {
	res := make([]string, n)
	for idx := range res {
		res[idx] = createNQueensRow(-1, n, ".")
	}
	return res
}

func createNQueensRow(posQ int, n int, symb string) string {
	var b strings.Builder
	b.Grow(n)
	idx := 0
	for idx < n {
		if idx == posQ {
			b.WriteString("Q")
		} else {
			b.WriteString(symb)
		}
		idx++
	}
	return b.String()
}

func strReplace(init string, pos int, symb string) string {
	var b strings.Builder
	b.Grow(len(init))
	if pos == 0 {
		b.WriteString(symb)
		b.WriteString(init[1:])
	} else if pos == len(init)-1 {
		b.WriteString(init[:len(init)-1])
		b.WriteString(symb)
	} else {
		b.WriteString(init[:pos])
		b.WriteString(symb)
		b.WriteString(init[pos+1:])
	}
	return b.String()
}

func setCellsUnderAttack(qRow, qCol int, board []string) error {
	//update col
	for idx := range board {
		if idx <= qRow {
			continue
		}
		if board[idx][qCol] == 'Q' {
			return errors.New("not valid board")
		}
		board[idx] = strReplace(board[idx], qCol, "x")
		diff := qRow - idx
		idxC := qCol - diff
		if idxC >= 0 && idxC < len(board) {
			if board[idx][idxC] == 'Q' {
				return errors.New("not valid board")
			}
			board[idx] = strReplace(board[idx], idxC, "x")
		}
		idxC = qCol + diff
		if idxC >= 0 && idxC < len(board) {
			if board[idx][idxC] == 'Q' {
				return errors.New("not valid board")
			}
			board[idx] = strReplace(board[idx], idxC, "x")
		}
	}
	return nil
}
