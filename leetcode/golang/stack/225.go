package stack

import "container/list"

type MyStack struct {
	Stack *list.List
}

/** Initialize your data structure here. */
func Constructor1() MyStack {
	return MyStack{
		Stack: list.New(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Stack.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	len := this.Stack.Len()
	for i := 0; i < len; i++ {
		e := this.Stack.Remove(this.Stack.Front())
		if i == len-1 {
			return e.(int)
		}
		this.Stack.PushBack(e)
	}
	return -1
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Stack.Back().Value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Stack.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
