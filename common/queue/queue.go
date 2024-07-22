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
func (q *Q[T]) Empty() bool {
	return q.Len() == 0
}

// Len 获取栈大小
func (q *Q[T]) Len() int {
	return len(q.arr)
}

// Add 元素入队
func (q *Q[T]) Add(elm T) {
	q.arr = append(q.arr, elm)
}

// Poll 队首元素出队
func (q *Q[T]) Poll() T {
	if q.Empty() {
		panic("empty queue")
	}
	elm := q.arr[0]
	q.arr = q.arr[1:]
	return elm
}

// Peek 获取队首元素
func (q *Q[T]) Peek() T {
	if q.Empty() {
		panic("empty queue")
	}
	return q.arr[0]
}

// String 返回栈中元素
func (q *Q[T]) String() string {
	return fmt.Sprintf("%v", q.arr)
}
