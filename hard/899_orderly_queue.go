package hard

import "strings"

func orderlyQueue(s string, k int) string {
	if k > 1 {
		letters := make(map[byte]int)
		for i := 0; i < len(s); i++ {
			letters[s[i]] += 1
		}
		var b strings.Builder
		b.Grow(len(s))
		for _, v := range "abcdefghijklmnopqrstuvwxyz" {
			if count, ok := letters[byte(v)]; ok {
				b.WriteString(strings.Repeat(string(byte(v)), count))
			}
		}
		return b.String()
	} else {
		res := s
		for i := 0; i < len(s); i++ {
			s = s[1:] + string(s[0])
			if s < res {
				res = s
			}
		}
		return res
	}
}
