package main

import "strconv"

func calPoints(operations []string) int {
	stack := []int{}
	for _, val := range operations {
		if val == "D" {
			stack = append(stack, stack[len(stack)-1]*2)
		} else if val == "C" {
			stack = stack[:len(stack)-1]
		} else if val == "+" {
			stack = append(stack, stack[len(stack)-2]+stack[len(stack)-1])
		} else {
			digit, _ := strconv.Atoi(val)
			stack = append(stack, digit)
		}
	}
	var result int
	for _, val := range stack {
		result += val
	}
	return result
}
