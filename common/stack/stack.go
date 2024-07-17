// 通用栈结构实现
package stack

import "fmt"

type S[T any] struct {
	arr []T
}

// New 初始化一个通用栈结构
// 传入指定类型的任意值确定泛型
func New[T any](typeElm T) *S[T] {
	return &S[T]{arr: []T{}}
}

// Empty 检查是否栈空
func (s *S[T]) Empty() bool {
	return s.Len() == 0
}

// Len 获取栈大小
func (s *S[T]) Len() int {
	return len(s.arr)
}

// Push 将元素压入栈中
func (s *S[T]) Push(elm T) {
	s.arr = append(s.arr, elm)
}

// Pop 弹出栈顶元素
func (s *S[T]) Pop() T {
	if s.Empty() {
		panic("empty stack")
	}
	elm := s.arr[s.Len()-1]
	s.arr = s.arr[:s.Len()-1]
	return elm
}

// Peek 获取栈顶元素
func (s *S[T]) Peek() T {
	if s.Empty() {
		panic("empty stack")
	}
	return s.arr[s.Len()-1]
}

// String 返回栈中元素
func (s *S[T]) String() string {
	return fmt.Sprintf("%v", s.arr)
}
