// 解法1
// 通过递归实现两个函数 getAndRemoveLastElm() 和 reverse()
// 作用分别是 “取出并移除栈底元素” 以及 “反转栈”
package main

import "fmt"

func main() {
	s := NewStack()
	for i := 1; i <= 5; i++ {
		s.S = append(s.S, i)
	}
	fmt.Println("栈初始状态: ", s.S)
	s.Reverse()
	fmt.Println("栈倒置后的状态: ", s.S)
}

type Stack struct {
	S []int
}

func NewStack() *Stack {
	return &Stack{S: []int{}}
}

func (s *Stack) Reverse() {
	if len(s.S) == 0 {
		return
	}
	last := s.GetAndRemoveLastElm()
	s.Reverse()
	s.S = append(s.S, last)
}

func (s *Stack) GetAndRemoveLastElm() int {
	res := s.S[len(s.S)-1]
	s.S = s.S[:len(s.S)-1]
	if len(s.S) == 0 {
		// 取完一个数后, 栈空, 说明已经取到栈底元素
		return res
	}
	// 栈不为空, 递归调用, 取出栈底, 再将当前数重新压回栈中
	last := s.GetAndRemoveLastElm()
	s.S = append(s.S, res)
	return last
}
