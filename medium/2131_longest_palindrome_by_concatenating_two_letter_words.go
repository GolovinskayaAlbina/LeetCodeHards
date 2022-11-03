package medium

func longestPalindrome(words []string) int {
	if len(words) == 0 {
		return 0
	}
	col := make(map[string]int)
	for _, w := range words {
		col[w] = col[w] + 1
	}
	length := 0
	middle := false
	for w, _ := range col {
		if col[w] == 0 {
			continue
		}
		if w[0] == w[1] {
			if col[w] > 1 { // get only even to left and right side ot word, leave 1 if not even
				even := (col[w] / 2) * 2
				col[w] = col[w] - even
				length += even * 2
			}
			if !middle && col[w] == 1 { //only one 'cc' can be in the middle
				length += 2
				col[w] = 0
				middle = true
			}
		} else {
			palWord := string(w[1]) + string(w[0])
			for _, ok := col[palWord]; ok && col[palWord] != 0 && col[w] != 0; { //few 'ad' 'da' here
				col[palWord] = col[palWord] - 1
				col[w] = col[w] - 1
				length += 4
			}
		}
	}
	return length
}
