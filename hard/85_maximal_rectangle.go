package hard

type position struct {
	end   int
	start int
	rows  int
	r     int
}

func maximalRectangle(matrix [][]byte) int {
	var stack []*position
	res := 0
	for r := 0; r < len(matrix); r++ {
		count := 0
		for c := 0; c <= len(matrix[0]); c++ {
			if count > 0 && (c == len(matrix[0]) || matrix[r][c] == '0') {
				p := &position{start: c - count, end: c, rows: 1, r: r}
				skipAdding := false
				var nearest *position
				for _, item := range stack {
					if item.r == r-1 {
						if item.start == p.start && item.end == p.end {
							item.rows += 1
							item.r = r
							skipAdding = true
							break
						}
						tmp := (item.end - item.start) * item.rows
						if res < tmp {
							res = tmp
						}
						if p.start >= item.start && p.end <= item.end {
							if nearest == nil || nearest.rows < item.rows {
								nearest = item
							}
						} else if item.start >= p.start && item.end <= p.end {
							item.rows += 1
							item.r = r
						} else {
							start := maxCoord(item.start, p.start)
							end := minCoord(item.end, p.end)
							if end-start > 0 {
								stack = append(stack, &position{start: start, end: end, rows: item.rows + 1, r: r})
							}
						}
					}
				}
				if nearest != nil {
					nearest.start = p.start
					nearest.end = p.end
					nearest.rows += 1
					nearest.r = r
					skipAdding = true
				}
				if !skipAdding {
					stack = append(stack, p)
				}
				count = 0
			} else if c < len(matrix[0]) && matrix[r][c] == '1' {
				count++
			}
		}
		var updated []*position
		for _, item := range stack {
			if item.r == r {
				updated = append(updated, item)
			} else {
				tmp := (item.end - item.start) * item.rows
				if res < tmp {
					res = tmp
				}
			}
		}
		stack = updated
	}
	for _, item := range stack {
		tmp := (item.end - item.start) * item.rows
		if res < tmp {
			res = tmp
		}
	}
	return res
}

func minCoord(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxCoord(x, y int) int {
	if x < y {
		return y
	}
	return x
}
