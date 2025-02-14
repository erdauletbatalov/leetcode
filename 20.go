package main

func isValid(s string) bool {
	a, b, c := 0, 0, 0
	for _, val := range s {
		switch val {
		case '{':
			a++
		case '(':
			b++
		case '[':
			c++
		case '}':
			a--
		case ')':
			b--
		case ']':
			c--
		}
		if a < 0 || b < 0 || c < 0 {
			return false
		}
	}

	if a == 0 && b == 0 && c == 0 {
		return true
	}
	return false
}
