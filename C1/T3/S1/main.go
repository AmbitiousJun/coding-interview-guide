// 解法1
// 通过递归实现两个函数 getAndRemoveLastElm() 和 reverse()
// 作用分别是 “取出并移除栈底元素” 以及 “反转栈”
package main

import (
	"coding_interview_guide/common/stack"
	"fmt"
)

func main() {
	s := NewStack()
	for i := 1; i <= 5; i++ {
		s.Push(i)
	}
	fmt.Println("栈初始状态: ", s.S)
	s.Reverse()
	fmt.Println("栈倒置后的状态: ", s.S)
}

type Stack struct {
	*stack.S[int]
}

func NewStack() *Stack {
	return &Stack{stack.New[int]()}
}

func (s *Stack) Reverse() {
	if s.Empty() {
		return
	}
	last := s.GetAndRemoveLastElm()
	s.Reverse()
	s.Push(last)
}

func (s *Stack) GetAndRemoveLastElm() int {
	res := s.Pop()
	if s.Empty() {
		// 取完一个数后, 栈空, 说明已经取到栈底元素
		return res
	}
	// 栈不为空, 递归调用, 取出栈底, 再将当前数重新压回栈中
	last := s.GetAndRemoveLastElm()
	s.Push(res)
	return last
}
