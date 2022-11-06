package hard

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var res []string
	r := &row{Max: maxWidth}
	for _, word := range words {
		if r.CanAdd(word) {
			r.Add(word)
		} else {
			res = append(res, r.String(false))
			r = &row{Max: maxWidth}
			r.Add(word)
		}
	}
	res = append(res, r.String(true))
	return res
}

type row struct {
	Max   int
	words []string
	l     int
}

func (r *row) String(isLast bool) string {
	if r.len() == 0 {
		return ""
	}
	res := ""
	if isLast || len(r.words) == 1 {
		for _, w := range r.words {
			res += w + " "
		}
		return res[:len(res)-1] + strings.Repeat(" ", r.Max-r.minLen())
	}
	//s := make([]string, r.spaces())
	//for i, _ := range s {
	//	s[i] = " "
	//}
	//for i := 0; leave > 0; leave-- {
	//	if i == len(s) {
	//		i = 0
	//	}
	//	s[i] += " "
	//	i++
	//}
	spacesCount := r.spaces()              // 5
	spacesToAdd := r.Max - r.minLen()      // 12
	div := spacesToAdd / spacesCount       // 2
	remainder := spacesToAdd % spacesCount // 2
	for idx, _ := range r.words {
		if idx == len(r.words)-1 {
			res += r.words[idx]
		} else {
			space := strings.Repeat(" ", 1+div)
			if remainder > 0 {
				space += " "
				remainder--
			}
			res += r.words[idx] + space
		}
	}
	return res
}

func (r *row) Add(word string) {
	r.words = append(r.words, word)
	r.l += len(word)
}

func (r *row) CanAdd(word string) bool {
	if len(r.words) == 0 {
		return len(word) <= r.Max
	}
	return (r.minLen() + 1 + len(word)) <= r.Max
}

func (r *row) minLen() int {
	return r.len() + r.spaces()
}

func (r *row) len() int {
	return r.l
}

func (r *row) spaces() int {
	if len(r.words) < 2 {
		return 0
	}
	return len(r.words) - 1
}
