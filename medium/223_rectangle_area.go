package medium

func computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int) int {
	area := getIntersectionArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
	areaA := countArea(ax1, ay1, ax2, ay2)
	areaB := countArea(bx1, by1, bx2, by2)
	return areaA + areaB - area
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func getIntersectionArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int) int {
	aLen := min(bx2, ax2) - max(bx1, ax1)
	if aLen < 1 {
		return 0
	}
	aHight := min(by2, ay2) - max(by1, ay1)
	if aHight < 1 {
		return 0
	}
	return aLen * aHight
}

func countArea(x1, y1, x2, y2 int) int {
	return (x2 - x1) * (y2 - y1)
}
