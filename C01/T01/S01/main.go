// 解法1
// 维护两个栈 stackData 和 stackMin, 分别用于存储栈数据和栈中当前的最小数据
// push: num 入栈 stackData
//
//	① stackMin 为空, num 入栈 stackMin
//	② stackMin 不为空, 但 num <= (stackMin 中的最小值), num 入栈 stackMin
//
// pop: num 从 stackData 出栈
//
//	① 如果 num == (stackMin 中的最小值), 从 stackMin 中出栈
//
// getMin: 获取 stackMin 的栈顶元素
package main

import (
	"coding_interview_guide/common/stack"
	"fmt"
)

func main() {
	s := NewMyStack()
	pushArr := []int{3, 4, 5, 1, 2, 1}
	for _, num := range pushArr {
		s.Push(num)
		fmt.Println("push ", num, ", current min: ", s.Min())
	}
	for range pushArr {
		num := s.Pop()
		fmt.Println("pop ", num, ", current min: ", s.Min())
	}
}

type MyStack struct {
	stackData *stack.S[int]
	stackMin  *stack.S[int]
}

func NewMyStack() *MyStack {
	return &MyStack{
		stackData: stack.New[int](),
		stackMin:  stack.New[int](),
	}
}

func (ms *MyStack) Push(num int) {
	ms.stackData.Push(num)
	if ms.stackMin.Empty() || num <= ms.Min() {
		ms.stackMin.Push(num)
	}
}

func (ms *MyStack) Pop() int {
	if ms.stackData.Empty() {
		panic("Empty stack")
	}
	num := ms.stackData.Pop()
	if num == ms.Min() {
		ms.stackMin.Pop()
	}
	return num
}

func (ms *MyStack) Min() int {
	if ms.stackMin.Empty() {
		panic("Empty stack")
	}
	return ms.stackMin.Peek()
}
