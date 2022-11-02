package medium

import "strings"

func minMutation(start string, end string, bank []string) int {
	if bankContains(end, bank) < 0 {
		return -1
	}
	if start == end {
		return 0
	}
	f, s := minMutationRec(start, end, bank)
	if f {
		return s
	}
	return -1
}

func minMutationRec(start string, end string, bank []string) (bool, int) {
	if start == end {
		return true, 0
	}
	if len(bank) == 0 {
		return false, 0
	}
	found := false
	min := len(bank)
	for i := 0; i < len(start); i++ {
		for _, symb := range []string{"A", "C", "G", "T"} {
			newStart := replaceSymbol(start, i, symb)
			bIbx := bankContains(newStart, bank)
			if bIbx >= 0 {
				f, s := minMutationRec(newStart, end, replaceItemInBank(bank, bIbx))
				if f {
					found = f
					if min > (1 + s) {
						min = 1 + s
					}
				}
			}
		}
	}

	return found, min
}

func replaceItemInBank(bank []string, pos int) []string {
	if pos == 0 {
		return bank[1:]
	} else if pos == len(bank)-1 {
		return bank[:len(bank)-1]
	} else {
		var res []string
		res = append(res, bank[:pos]...)
		res = append(res, bank[pos+1:]...)
		return res
	}
	return []string{}
}

func replaceSymbol(str string, pos int, what string) string {
	var b strings.Builder
	b.Grow(len(str))
	if pos == 0 {
		b.WriteString(what)
		b.WriteString(str[1:])
	} else if pos == len(str)-1 {
		b.WriteString(str[:len(str)-1])
		b.WriteString(what)
	} else {
		b.WriteString(str[:pos])
		b.WriteString(what)
		b.WriteString(str[pos+1:])
	}
	return b.String()
}

func bankContains(str string, bank []string) int {
	for idx, b := range bank {
		if b == str {
			return idx
		}
	}
	return -1
}
