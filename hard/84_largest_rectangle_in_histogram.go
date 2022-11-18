package hard

import (
	str "leet-code/data_structure"
)

func largestRectangleArea(heights []int) int {
	type position struct {
		val int
		idx int
	}
	var stack str.Stack[position]
	stack.Push(position{idx: -1, val: 0})
	res := 0
	heights = append(heights, 0)
	for i := 0; i < len(heights); i++ {
		for stack.Len() > 0 && heights[i] < stack.Peek().val {
			val := stack.Pop().val
			tmp := val * (i - stack.Peek().idx - 1)
			if res < tmp {
				res = tmp
			}
		}
		if stack.Len() > 0 && stack.Peek().val == heights[i] {
			stack.Pop()
		}
		stack.Push(position{idx: i, val: heights[i]})
	}
	return res
}
