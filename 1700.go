package main

func countStudents(students []int, sandwiches []int) int {
	studentsQueue := queue{
		head: &qnode{},
		tail: &qnode{},
	}
	curr := studentsQueue.head
	for i := range students {
		curr.val = students[i]
		if i == 0 {
			studentsQueue.head = curr
		}
		if i == len(students)-1 {
			studentsQueue.tail = curr
		}
		curr.next = &qnode{}
		curr = curr.next
	}

	sandwichesQueue := queue{
		head: &qnode{},
		tail: &qnode{},
	}
	curr = sandwichesQueue.head
	for i := range sandwiches {
		curr.val = sandwiches[i]
		if i == 0 {
			sandwichesQueue.head = curr
		}
		if i == len(sandwiches)-1 {
			sandwichesQueue.tail = curr
		}
		curr.next = &qnode{}
		curr = curr.next
	}

	len := len(students)
	i := len
	var result int
	for len > 0 {
		if studentsQueue.head.val == sandwichesQueue.head.val {
			studentsQueue.RemoveFromHead()
			sandwichesQueue.RemoveFromHead()
			len--
			i = len
			continue
		} else {
			node := studentsQueue.RemoveFromHead()
			studentsQueue.AddToTail(node)
			if i--; i == 0 && len != 0 {
				result = len
				break
			}
		}

	}
	return result
}

type qnode struct {
	next *qnode
	val  int
}

type queue struct {
	head *qnode
	tail *qnode
}

func (q *queue) AddToTail(n *qnode) {
	q.tail.next = n
	q.tail = q.tail.next
}

func (q *queue) RemoveFromHead() *qnode {
	temp := q.head
	q.head = q.head.next
	return temp
}
