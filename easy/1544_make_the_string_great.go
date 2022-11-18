package easy

import (
	str "leet-code/data_structure"
)

func makeGood(s string) string {
	var stack str.Stack[uint8]
	diff := uint8('a' - 'A')
	for i, _ := range s {
		if stack.Len() > 0 && mod(stack.Peek(), s[i]) == diff {
			stack.Pop()
		} else {
			stack.Push(s[i])
		}
	}
	reverse := []byte(stack.String())
	i, j := 0, len(reverse)-1
	for i < j {
		reverse[i], reverse[j] = reverse[j], reverse[i]
		i++
		j--
	}
	return string(reverse)
}

func mod(x, y uint8) uint8 {
	if x >= y {
		return x - y
	}
	return y - x
}
