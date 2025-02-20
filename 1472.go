package main

type BrowserHistory struct {
	homepage string
	head     *BrowserPage
	curr     *BrowserPage
	currSize int
	size     int
}

type BrowserPage struct {
	val  string
	next *BrowserPage
	prev *BrowserPage
}

func Constructor1472(homepage string) BrowserHistory {
	return BrowserHistory{homepage: homepage}
}

func (this *BrowserHistory) Visit(url string) {
	if this.size == 0 || this.currSize == 0 {
		this.head = &BrowserPage{val: url}
		this.curr = this.head
		this.currSize++
		this.size = this.currSize
		return
	}
	this.curr.next = &BrowserPage{val: url, next: nil, prev: this.curr}
	this.curr = this.curr.next
	this.currSize++
	this.size = this.currSize
}

func (this *BrowserHistory) Back(steps int) string {
	if steps >= this.size || steps >= this.currSize {
		this.curr = nil
		this.currSize = 0
		return this.homepage
	}
	for i := 0; i < steps; i++ {
		if this.size == 0 || this.currSize == 0 {
			return this.homepage
		}
		this.curr = this.curr.prev
		this.currSize--
	}
	if this.curr == nil {
		return this.homepage
	}

	return this.curr.val
}

func (this *BrowserHistory) Forward(steps int) string {

	for i := 0; i < steps; i++ {
		if this.size == 0 {
			return this.homepage
		}
		if this.currSize >= this.size {
			return this.curr.val
		}
		if this.curr == nil {
			this.curr = this.head
			this.currSize++
			continue
		}
		this.curr = this.curr.next
		this.currSize++
	}
	return this.curr.val
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
