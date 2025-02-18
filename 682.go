package main

import (
	"strconv"
)

func calPoints(operations []string) int {
	stack := []int{}
	total := 0

	for _, val := range operations {
		switch val {
		case "D":
			if len(stack) == 0 {
				continue // Игнорируем некорректную операцию
			}
			double := stack[len(stack)-1] * 2
			stack = append(stack, double)
			total += double
		case "C":
			if len(stack) == 0 {
				continue // Игнорируем, если нечего удалять
			}
			total -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "+":
			if len(stack) < 2 {
				continue // Игнорируем, если недостаточно элементов
			}
			sum := stack[len(stack)-1] + stack[len(stack)-2]
			stack = append(stack, sum)
			total += sum
		default:
			digit, err := strconv.Atoi(val)
			if err != nil {
				continue // Игнорируем некорректные строки
			}
			stack = append(stack, digit)
			total += digit
		}
	}

	return total
}
