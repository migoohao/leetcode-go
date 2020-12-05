package queue

type MyQueue struct {
	stack  []int
	assist []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack:  make([]int, 0),
		assist: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	pop := this.Peek()
	if len(this.assist) > 0 {
		this.assist = this.assist[:len(this.assist)-1]
	}
	return pop
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.assist) == 0 {
		for len(this.stack) > 0 {
			this.assist = append(this.assist, this.stack[len(this.stack)-1])
			this.stack = this.stack[:len(this.stack)-1]
		}
	}
	if len(this.assist) > 0 {
		return this.assist[len(this.assist)-1]
	}
	return -1
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.assist) == 0 && len(this.stack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
