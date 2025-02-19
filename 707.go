package main

type MyLinkedList struct {
	head *Node
	size int
}

type Node struct {
	val  int
	next *Node
}

func Constructor707() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	curr := this.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	return curr.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.head = &Node{val: val, next: this.head}
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{val: val}
	if this.head == nil {
		this.head = newNode
	} else {
		curr := this.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = newNode
	}
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	curr := this.head
	for i := 0; i < index-1; i++ {
		curr = curr.next
	}
	curr.next = &Node{val: val, next: curr.next}
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	if index == 0 {
		this.head = this.head.next
	} else {
		curr := this.head
		for i := 0; i < index-1; i++ {
			curr = curr.next
		}
		curr.next = curr.next.next
	}

	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
