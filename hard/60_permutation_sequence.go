package hard

import "strconv"

func getPermutation(n int, k int) string {
	str := ""
	fact := 1
	for i := 1; i <= n; i++ {
		fact = fact * i
		str += strconv.Itoa(i)
	}
	k = k - 1
	block := fact / n
	res := getPermutationCombination(str, k, block, 0)
	k = k % block
	return res[k]
}

func getPermutationCombination(cur string, k int, block, level int) []string {
	if len(cur) == 1 {
		return []string{cur}
	}
	if len(cur) == 2 {
		res := []string{
			string(cur[0]) + string(cur[1]),
			string(cur[1]) + string(cur[0]),
		}
		if level == 0 && k < len(res) {
			return []string{res[k]}
		}
	}
	var res []string
	for i := 0; i < len(cur); i++ {
		if level == 0 {
			if k < block*i || k >= block*(i+1) {
				continue
			}
		}
		head := string(cur[i])
		var next string
		if i == 0 {
			next = cur[i+1:]
		} else if i == len(cur)-1 {
			next = cur[:len(cur)-1]
		} else {
			next = cur[:i] + cur[i+1:]
		}
		combo := getPermutationCombination(next, k, block, level+1)
		for _, s := range combo {
			res = append(res, head+s)
		}
	}
	return res
}
