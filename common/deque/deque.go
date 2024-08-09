// 通用双端队列结构实现
package deque

import "fmt"

type D[T any] struct {
	arr []T
}

// New 初始化一个通用双端队列结构
func New[T any]() *D[T] {
	return &D[T]{arr: []T{}}
}

// Empty 检查是否栈空
func (d *D[T]) Empty() bool {
	return d.Len() == 0
}

// Len 获取栈大小
func (d *D[T]) Len() int {
	return len(d.arr)
}

// AddLast 元素入队尾
func (d *D[T]) AddLast(elm T) {
	d.arr = append(d.arr, elm)
}

// AddFirst 元素入队首
func (d *D[T]) AddFirst(elm T) {
	d.arr = append([]T{elm}, d.arr...)
}

// PollLast 队首元素出队尾
func (d *D[T]) PollLast() T {
	if d.Empty() {
		panic("empty deque")
	}
	elm := d.arr[d.Len()-1]
	d.arr = d.arr[:d.Len()-1]
	return elm
}

// PollFirst 队首元素出队首
func (d *D[T]) PollFirst() T {
	if d.Empty() {
		panic("empty deque")
	}
	elm := d.arr[0]
	d.arr = d.arr[1:]
	return elm
}

// PeekLast 获取队尾元素
func (d *D[T]) PeekLast() T {
	if d.Empty() {
		panic("empty deque")
	}
	return d.arr[d.Len()-1]
}

// PeekFirst 获取队首元素
func (d *D[T]) PeekFirst() T {
	if d.Empty() {
		panic("empty deque")
	}
	return d.arr[0]
}

// String 返回栈中元素
func (d *D[T]) String() string {
	return fmt.Sprintf("%v", d.arr)
}
