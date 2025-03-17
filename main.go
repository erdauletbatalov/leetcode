package main

import (
	"container/list"
)

func main() {
	// Создаем новый список
	myList := list.New()
	myList.Init()
	myList.PushBack(3)
	myList.PushFront(4)
	myList.InsertAfter(1, myList.Front())
	myList.InsertBefore(2, myList.Back())

	for e := myList.Front(); e != nil; e = e.Next() {
		print(e.Value.(int), " ")
	}
	println()
	print(myList.Back().Value.(int))
	print(" ")
	print(myList.Front().Value.(int))
	println()
	myList.Remove(myList.Front())

	for e := myList.Front(); e != nil; e = e.Next() {
		print(e.Value.(int), " ")
	}
	println()

	println("Len: ", myList.Len())

	myList.MoveAfter(myList.Front(), myList.Back())
	println()
	for e := myList.Front(); e != nil; e = e.Next() {
		print(e.Value.(int), " ")
	}
	println()

	myList.InsertAfter(5, myList.Back())

	myList.MoveBefore(myList.Front(), myList.Back())
	println()
	for e := myList.Front(); e != nil; e = e.Next() {
		print(e.Value.(int), " ")
	}
	println()

	myList.Init()

	println()
	for e := myList.Front(); e != nil; e = e.Next() {
		print(e.Value.(int), " ")
	}
	println()
	// myList.Remove(myList.Front())
	// myList.Remove(myList.Back())
	// myList.Remove(myList.Front())
	// myList.Remove(myList.Back())
}
