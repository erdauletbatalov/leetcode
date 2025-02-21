package main

func main() {
	mystack := Constructor225()
	mystack.Push(5)
	mystack.Push(4)
	mystack.Push(3)
	mystack.Push(2)
	mystack.Pop()
	mystack.Pop()
	mystack.Pop()
	mystack.Push(2)
	mystack.Push(2)
	mystack.Empty()
}
