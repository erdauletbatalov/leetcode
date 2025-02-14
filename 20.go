package main

func isValid(s string) bool {
	stack := []rune{}
	for _, val := range s {
		stack = append(stack, val)
		if len(stack) >= 2 {
			if string(stack[len(stack)-2:]) == "{}" || string(stack[len(stack)-2:]) == "()" || string(stack[len(stack)-2:]) == "[]" {
				stack = stack[:len(stack)-2]
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
