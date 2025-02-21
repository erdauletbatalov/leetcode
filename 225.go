package main

import "container/list"

type MyStack struct {
	q1 *list.List
	q2 *list.List
}

func Constructor225() MyStack {
	return MyStack{q1: list.New(), q2: list.New()}
}

func (this *MyStack) Push(x int) {
	// Step 1: add x to q2
	this.q2.PushBack(x)

	// Step 2: move q1 elements into q2
	for this.q1.Len() != 0 {
		this.q2.PushBack(this.q1.Remove(this.q1.Front()))
	}

	// Step 3: Swap s1 with s2

	this.q1, this.q2 = this.q2, this.q1
}

func (this *MyStack) Pop() int {
	if this.q1.Len() == 0 {
		return 0
	}
	front := this.q1.Front()
	this.q1.Remove(front)
	return front.Value.(int)
}

func (this *MyStack) Top() int {
	if this.q1.Len() == 0 {
		return 0
	}
	return this.q1.Front().Value.(int)
}

func (this *MyStack) Empty() bool {
	if this.q1.Len() == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
