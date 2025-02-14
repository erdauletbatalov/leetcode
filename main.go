package main

import (
	"fmt"
)

func main() {
	// fmt.Println(canPlaceFlowers([]int{0, 0, 0, 0, 1}, 3))
	// fmt.Println("Check")
	// fmt.Println("Second check")
	// fmt.Println(getConcatenation([]int{1, 2, 1}))
	MinStack := Constructor()
	MinStack.Push(3)
	MinStack.Push(5)
	MinStack.Push(8)
	MinStack.Push(2)
	MinStack.Push(9)
	MinStack.Push(7)
	MinStack.Push(1)
	MinStack.Push(0)
	fmt.Println(MinStack)
	fmt.Println()

	fmt.Println(MinStack.GetMin())
	MinStack.Pop()
	fmt.Println(MinStack.GetMin())
	MinStack.Pop()
	fmt.Println(MinStack.GetMin())
	fmt.Println(MinStack.Top())
	fmt.Println(MinStack)

	fmt.Println()
}
