// 解法1
// 准备一个 help 栈, 维护使得该栈变成一个从顶到底小到大的顺序
// 维护方法：从 stack 栈中取出一个元素 cur, 并将其与 help 栈的栈顶 peek 进行比较
//
//	① cur <= peek, cur 直接入 help 栈
//	② cur > peek, 重复弹出 help 栈中的元素到 stack 栈中, 直到符合条件 ①
package main

import (
	"coding_interview_guide/common/stack"
	"fmt"
)

func main() {
	nums := []int{1, 8, 8, 1, 1, 0, 1, 2, 1, 3, 8, 1, 5, 3, 5, 8, 2, 0, 7, 7, 1}
	s := stack.New[int]()
	for _, num := range nums {
		s.Push(num)
	}
	fmt.Println("当前栈中元素: ", s)
	SortStackByStack(s)
	fmt.Println("排序后栈中元素: ", s)
}

func SortStackByStack(s *stack.S[int]) {
	help := stack.New[int]()
	for !s.Empty() {
		cur := s.Pop()
		for !help.Empty() && help.Peek() < cur {
			s.Push(help.Pop())
		}
		help.Push(cur)
	}
	for !help.Empty() {
		s.Push(help.Pop())
	}
}
