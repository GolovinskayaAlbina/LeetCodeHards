package hard

import (
	"sort"
	"strings"
)

var ump = make(map[string]bool)

func isScramble1(s1 string, s2 string) bool {
	n := len(s1)
	key := s1 + "_" + s2

	if val, ok := ump[key]; ok {
		return val
	}
	if n != len(s2) {
		return false
	}
	if s1 == s2 {
		ump[key] = true
		return true
	}
	if n < 2 {
		return false
	}
	if sorted(s1) != sorted(s2) {
		ump[key] = false
		return false
	}
	for i := 1; i < n; i++ {
		f := s1[i:n]
		s := s1[0:i]
		if (f + s) == s2 {
			ump[key] = true
			return true
		}
		if isScramble(f, s2[i:n]) && isScramble(s, s2[0:i]) {
			ump[key] = true
			return true
		}
		if isScramble(f, s2[0:n-i]) && isScramble(s, s2[n-i:n]) {
			ump[key] = true
			return true
		}
	}
	ump[key] = false
	return false
}

func sorted(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

var m = map[string]bool{}

func isScramble(s1 string, s2 string) bool {
	if len(s1) == 0 || s1 == s2 {
		return true
	}
	if len(s1) == 1 || len(s1) != len(s2) {
		return false
	}
	key := s1 + s2
	if res, ok := m[key]; ok {
		return res
	}
	if !equalIfSorted(s1, s2) {
		return false
	}

	res := false
	l := len(s1)
	for i := 1; i < l; i++ {
		if (isScramble(s1[:i], s2[l-i:]) && isScramble(s1[i:], s2[:l-i])) ||
			(isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:])) {
			res = true
			break
		}
	}
	m[key] = res
	return res
}

func equalIfSorted(s1 string, s2 string) bool {
	sort1 := []byte(s1)
	sort2 := []byte(s2)
	sort.Slice(sort1, func(i int, j int) bool { return sort1[i] < sort1[j] })
	sort.Slice(sort2, func(i int, j int) bool { return sort2[i] < sort2[j] })
	return string(sort1) == string(sort2)
}
