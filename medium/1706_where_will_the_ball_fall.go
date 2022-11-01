package medium

func findBall(grid [][]int) []int {
	if len(grid) == 0 {
		return nil
	}
	res := make([]int, len(grid[0]))
	for idx := range res {
		res[idx] = idx
	}

	for rowIdx := range grid {
		for posIdx := range grid[rowIdx] {
			if res[posIdx] == -1 {
				continue
			}
			if ok, pos := canMove(grid[rowIdx], res[posIdx]); ok {
				res[posIdx] = pos
			} else {
				res[posIdx] = -1
			}
		}
	}
	return res
}

func canMove(row []int, pos int) (bool, int) {
	if row[pos] == -1 && pos-1 >= 0 && row[pos-1] == -1 { // |/|/|
		return true, pos - 1
	}
	if row[pos] == 1 && pos+1 < len(row) && row[pos+1] == 1 { // |\|\|
		return true, pos + 1
	}
	return false, 0
}
