package medium

type point [2]int

func removeStones(stones [][]int) int {
	left, right := map[int][]point{}, map[int][]point{}
	v := map[point]bool{}
	heads := 0

	for _, s := range stones {
		x, y := s[0], s[1]
		left[x] = append(left[x], point{x, y})
		right[y] = append(right[y], point{x, y})
	}
	for _, s := range stones {
		x, y := s[0], s[1]
		if !v[point{x, y}] {
			dfs(point{x, y}, v, left, right)
			heads += 1
		}
	}
	return len(stones) - heads
}

func dfs(p point, v map[point]bool, left, right map[int][]point) {
	v[p] = true
	for _, np := range left[p[0]] {
		if !v[np] {
			dfs(np, v, left, right)
		}
	}
	for _, np := range right[p[1]] {
		if !v[np] {
			dfs(np, v, left, right)
		}
	}
}
