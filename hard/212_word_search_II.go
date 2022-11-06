package hard

func findWords(board [][]byte, words []string) []string {
	var res []string
	for _, w := range words {
		desired := &word{symbols: w}
		if desired.IsIn(board) {
			res = append(res, w)
		}
	}
	return res
}

type visited struct {
	r int
	c int
}

type word struct {
	symbols string
}

func (w *word) IsIn(board [][]byte) bool {
	if len(w.symbols) == 0 {
		return false
	}
	if !w.Placed(board) {
		return false
	}
	first := w.symbols[0]
	for r, _ := range board {
		for c, _ := range board[r] {
			if board[r][c] == first {
				if w.CanMoveToTheEnd(r, c, 0, board) {
					return true
				}
			}
		}
	}
	return false
}

func (w *word) alreadyVisit(r, c int, v []visited) bool {
	for _, v := range v {
		if v.r == r && v.c == c {
			return true
		}
	}
	return false
}

// Placed additional check to accelerate execution
// not necessary to use
func (w *word) Placed(board [][]byte) bool {
	wSymbols := make(map[byte]int)
	for _, char := range w.symbols {
		wSymbols[byte(char)] += 1
	}
	bSymbols := make(map[byte]int)
	for r, _ := range board {
		for c, _ := range board[r] {
			bSymbols[board[r][c]] += 1
		}
	}
	for k, v := range wSymbols {
		if val, ok := bSymbols[k]; !ok || val < v {
			return false
		}
	}
	return true
}

func (w *word) canGo(pos, r, c int, board [][]byte, v []visited) bool {
	if w.alreadyVisit(r, c, v) {
		return false
	}
	v = append(v, visited{r, c})
	if pos == len(w.symbols) {
		return true
	}

	// check no ways - additional check to accelerate execution
	// not necessary to use
	if (r == len(board)-1 || board[r+1][c] != w.symbols[pos]) &&
		(r == 0 || board[r-1][c] != w.symbols[pos]) &&
		(c == 0 || board[r][c-1] != w.symbols[pos]) &&
		(c == len(board[0])-1 || board[r][c+1] != w.symbols[pos]) {
		return false
	}

	if r+1 < len(board) && board[r+1][c] == w.symbols[pos] { //bottom
		if w.canGo(pos+1, r+1, c, board, v) {
			return true
		}
	}
	if r-1 >= 0 && board[r-1][c] == w.symbols[pos] { //top
		if w.canGo(pos+1, r-1, c, board, v) {
			return true
		}
	}
	if c-1 >= 0 && board[r][c-1] == w.symbols[pos] { //left
		if w.canGo(pos+1, r, c-1, board, v) {
			return true
		}
	}
	if c+1 < len(board[0]) && board[r][c+1] == w.symbols[pos] { //right
		if w.canGo(pos+1, r, c+1, board, v) {
			return true
		}
	}
	return false
}

func (w *word) CanMoveToTheEnd(r, c, pos int, board [][]byte) bool {
	if pos == len(w.symbols)-1 {
		return true
	}
	return w.canGo(pos+1, r, c, board, []visited{})
}
