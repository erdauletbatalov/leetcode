package main

import (
	"fmt"
)

func main() {
	ll := Constructor707()
	fmt.Println(ll)

	ll.AddAtTail(10)
	// ll.AddAtTail(20)
	fmt.Println(ll)

	// ll.AddAtIndex(0, 10)
	// ll.AddAtIndex(0, 20)
	// ll.AddAtIndex(1, 30)

	// fmt.Println(ll.Get(0))
	fmt.Println(ll.Get(1))

	// ll.AddAtHead(5)
	// fmt.Println(ll.head, ll.head.next)

	// ll.DeleteAtIndex(0)
	// fmt.Println(ll)
}
