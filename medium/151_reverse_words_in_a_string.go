package medium

func reverseWords(s string) string {
	var res []byte

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		}
		j := i
		for j-1 >= 0 && s[j-1] != ' ' {
			j--
		}
		if len(res) != 0 {
			res = append(res, ' ')
		}
		start := j
		for ; j <= i; j++ {
			res = append(res, s[j])
		}
		i = start
	}

	return string(res)
}
