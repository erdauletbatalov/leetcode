package main

type BrowserHistory struct {
	list *node
}

type node struct {
	val  string
	next *node
	prev *node
}

func Constructor1472(homepage string) BrowserHistory {
	return BrowserHistory{list: &node{val: homepage}}
}

func (this *BrowserHistory) Visit(url string) {
	this.list.next = &node{val: url, next: nil, prev: this.list}
	this.list = this.list.next
}

func (this *BrowserHistory) Back(steps int) string {
	for this.list.prev != nil {
		this.list = this.list.prev
		if steps--; steps == 0 {
			break
		}
	}
	return this.list.val
}

func (this *BrowserHistory) Forward(steps int) string {

	for this.list.next != nil {
		this.list = this.list.next
		if steps--; steps == 0 {
			break
		}
	}
	return this.list.val
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
