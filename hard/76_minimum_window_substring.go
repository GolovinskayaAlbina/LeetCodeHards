package hard

type pos struct {
	char byte
	p    int
}

func minWindow(s string, t string) string {
	dictT := make(map[byte]int)
	for _, char := range []byte(t) {
		dictT[char] += 1
	}
	var filteredS []*pos
	for i, char := range []byte(s) {
		if _, ok := dictT[char]; ok {
			filteredS = append(filteredS, &pos{char: char, p: i})
		}
	}
	l, r, formed := 0, 0, 0
	windowCounts := make(map[byte]int)
	ans := []int{-1, 0, 0}
	required := len(dictT)
	for r < len(filteredS) {
		c := filteredS[r].char
		count := windowCounts[c]
		windowCounts[c] = count + 1

		if res, ok := dictT[c]; ok && windowCounts[c] == res {
			formed++
		}

		for l <= r && formed == required {
			c = filteredS[l].char

			end := filteredS[r].p
			start := filteredS[l].p
			if ans[0] == -1 || end-start+1 < ans[0] {
				ans[0] = end - start + 1
				ans[1] = start
				ans[2] = end
			}

			windowCounts[c] = windowCounts[c] - 1
			if res, ok := dictT[c]; ok && windowCounts[c] < res {
				formed--
			}
			l++
		}
		r++
	}
	if ans[0] == -1 {
		return ""
	}
	return s[ans[1] : ans[2]+1]
}
