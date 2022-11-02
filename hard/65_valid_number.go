package hard

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	containsDig := false
	for idx := range s {
		containsDig = s[idx] >= '0' && s[idx] <= '9'
		if containsDig {
			break
		}
	}
	if !containsDig {
		return false
	}
	dot, e := false, false
	for idx := range s {
		if s[idx] == '.' {
			if dot || e || (idx != 0 && s[idx-1] != '-' && s[idx-1] != '+' && (s[idx-1] < '0' || s[idx-1] > '9')) {
				return false
			}
			dot = true
		} else if s[idx] == 'e' || s[idx] == 'E' {
			if e || idx == 0 || ((s[idx-1] < '0' || s[idx-1] > '9') && s[idx-1] != '.') || idx == len(s)-1 {
				return false
			}
			mantissa := false
			for i := 0; i < idx && !mantissa; i++ {
				mantissa = s[i] >= '0' && s[i] <= '9'
			}
			if !mantissa {
				return false
			}
			e = true
		} else if s[idx] == '-' || s[idx] == '+' {
			if idx == len(s)-1 || (idx > 0 && (s[idx-1] != 'e' && s[idx-1] != 'E')) {
				return false
			}
		} else if s[idx] < '0' || s[idx] > '9' {
			return false
		}
	}
	return true
}
