/*
Design a stack that supports increment operations on its elements.

Implement the CustomStack class:

CustomStack(int maxSize) Initializes the object with maxSize which is the maximum number of elements in the stack.
void push(int x) Adds x to the top of the stack if the stack has not reached the maxSize.
int pop() Pops and returns the top of the stack or -1 if the stack is empty.
void inc(int k, int val) Increments the bottom k elements of the stack by val. If there are less than k elements in the stack, increment all the elements in the stack.

*/

package main

import "fmt"

type CustomStack struct {
	s       []int
	maxSize int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{s: make([]int, 0, maxSize), maxSize: maxSize}
}

func (this *CustomStack) Push(x int) {
	if len(this.s) < this.maxSize {
		this.s = append(this.s, x)
	}
}

func (this *CustomStack) Pop() int {
	if len(this.s) == 0 {
		return -1
	}
	ret := this.s[len(this.s)-1]
	this.s = this.s[:len(this.s)-1]
	return ret
}

func (this *CustomStack) Increment(k int, val int) {
	if k > len(this.s) {
		k = len(this.s)
	}
	for i := 0; i < k; i++ {
		this.s[i] += val
	}
}

func main() {
	obj := Constructor(3)
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Increment(1, 100)
	obj.Increment(2, 100)
	param_3 := obj.Pop()
	param_4 := obj.Pop()
	param_5 := obj.Pop()
	param_6 := obj.Pop()
	fmt.Println(param_3, param_4, param_5, param_6)

}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
