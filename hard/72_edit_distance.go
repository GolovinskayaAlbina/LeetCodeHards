package hard

func minDistance(from string, to string) int {
	if to == "" {
		return len(from)
	}
	if from == "" {
		return len(to)
	}
	if from == to {
		return 0
	}
	levenshteinDistance := make([][]int, len(to)+1)
	levenshteinDistance[0] = make([]int, len(from)+1)
	levenshteinDistance[0][0] = 0
	for r := 1; r <= len(to); r++ {
		levenshteinDistance[r] = make([]int, len(from)+1)
		levenshteinDistance[r][0] = r
	}
	for c := 1; c <= len(from); c++ {
		levenshteinDistance[0][c] = c
	}
	for r := 1; r <= len(to); r++ {
		for c := 1; c <= len(from); c++ {
			k := 1
			if to[r-1] == from[c-1] {
				k = 0
			}
			levenshteinDistance[r][c] = min(levenshteinDistance[r-1][c]+1, levenshteinDistance[r][c-1]+1, levenshteinDistance[r-1][c-1]+k)
		}
	}
	return levenshteinDistance[len(to)][len(from)]
}

func min(x, y, z int) int {
	m := x
	if m > y {
		m = y
	}
	if m > z {
		return z
	}
	return m
}
