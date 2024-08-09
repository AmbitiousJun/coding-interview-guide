// 解法1
// 准备两个栈 stackPush 和 stackPop, 一个只负责压入新数据, 另一个只负责弹出数据
// 在 add、poll、peek 操作前, 都尝试将 stackPush 中的数据倒入到 stackPop 中
package main

import (
	"coding_interview_guide/common/stack"
	"fmt"
)

func main() {
	q := NewTwoStacksQueue()
	pushArr := []int{1, 2, 3, 4, 5}
	for _, num := range pushArr {
		q.Add(num)
		fmt.Println("add ", num)
	}
	fmt.Println("poll ", q.Poll())
	fmt.Println("poll ", q.Poll())
	pushArr = []int{6, 7, 8, 9, 10}
	for _, num := range pushArr {
		q.Add(num)
		fmt.Println("add ", num)
	}
	for i := 0; i < 8; i++ {
		fmt.Println("poll ", q.Poll())
	}
}

type TwoStacksQueue struct {
	stackPush *stack.S[int]
	stackPop  *stack.S[int]
}

func NewTwoStacksQueue() *TwoStacksQueue {
	return &TwoStacksQueue{
		stackPush: stack.New[int](),
		stackPop:  stack.New[int](),
	}
}

func (tsq *TwoStacksQueue) push2Pop() {
	if !tsq.stackPop.Empty() {
		// pop 栈不为空, 不能倒入新数据
		return
	}
	for !tsq.stackPush.Empty() {
		tsq.stackPop.Push(tsq.stackPush.Pop())
	}
}

func (tsq *TwoStacksQueue) Add(num int) {
	tsq.stackPush.Push(num)
	tsq.push2Pop()
}

func (tsq *TwoStacksQueue) Poll() int {
	tsq.push2Pop()
	if tsq.stackPop.Empty() {
		panic("Empty queue!")
	}
	return tsq.stackPop.Pop()
}

func (tsq *TwoStacksQueue) Peek() int {
	tsq.push2Pop()
	if tsq.stackPop.Empty() {
		panic("Empty queue!")
	}
	return tsq.stackPop.Peek()
}
