// 解法1
// 准备两个栈 stackPush 和 stackPop, 一个只负责压入新数据, 另一个只负责弹出数据
// 在 add、poll、peek 操作前, 都尝试将 stackPush 中的数据倒入到 stackPop 中
package main

import "fmt"

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
	stackPush []int
	stackPop  []int
}

func NewTwoStacksQueue() *TwoStacksQueue {
	return &TwoStacksQueue{
		stackPush: []int{},
		stackPop:  []int{},
	}
}

func (tsq *TwoStacksQueue) push2Pop() {
	if len(tsq.stackPop) != 0 {
		// pop 栈不为空, 不能倒入新数据
		return
	}
	for len(tsq.stackPush) != 0 {
		num := tsq.stackPush[len(tsq.stackPush)-1]
		tsq.stackPush = tsq.stackPush[:len(tsq.stackPush)-1]
		tsq.stackPop = append(tsq.stackPop, num)
	}
}

func (tsq *TwoStacksQueue) Add(num int) {
	tsq.stackPush = append(tsq.stackPush, num)
	tsq.push2Pop()
}

func (tsq *TwoStacksQueue) Poll() int {
	tsq.push2Pop()
	if len(tsq.stackPop) == 0 {
		panic("Empty queue!")
	}
	num := tsq.stackPop[len(tsq.stackPop)-1]
	tsq.stackPop = tsq.stackPop[:len(tsq.stackPop)-1]
	return num
}

func (tsq *TwoStacksQueue) Peek() int {
	tsq.push2Pop()
	if len(tsq.stackPop) == 0 {
		panic("Empty queue!")
	}
	return tsq.stackPop[len(tsq.stackPop)-1]
}
