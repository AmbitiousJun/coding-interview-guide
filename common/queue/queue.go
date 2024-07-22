// 通用队列结构实现
package queue

import "fmt"

type Q[T any] struct {
	arr []T
}

// New 初始化一个通用队列结构
// 传入指定类型的任意值确定泛型
func New[T any](typeElm T) *Q[T] {
	return &Q[T]{arr: []T{}}
}

// Empty 检查是否栈空
func (s *Q[T]) Empty() bool {
	return s.Len() == 0
}

// Len 获取栈大小
func (s *Q[T]) Len() int {
	return len(s.arr)
}

// Add 元素入队
func (s *Q[T]) Add(elm T) {
	s.arr = append(s.arr, elm)
}

// Poll 队首元素出队
func (s *Q[T]) Poll() T {
	if s.Empty() {
		panic("empty queue")
	}
	elm := s.arr[0]
	s.arr = s.arr[1:]
	return elm
}

// Peek 获取队首元素
func (s *Q[T]) Peek() T {
	if s.Empty() {
		panic("empty queue")
	}
	return s.arr[0]
}

// String 返回栈中元素
func (s *Q[T]) String() string {
	return fmt.Sprintf("%v", s.arr)
}
