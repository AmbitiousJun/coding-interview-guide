// 通用的列表结构
package list

import "fmt"

type L[T any] struct {
	arr []T
}

// New 初始化一个通用列表结构
// 传入指定类型的任意值确定泛型
func New[T any](typeElm T) *L[T] {
	return &L[T]{arr: []T{}}
}

// Empty 检查列表是否空
func (l *L[T]) Empty() bool {
	return l.Len() == 0
}

// Len 获取列表大小
func (l *L[T]) Len() int {
	return len(l.arr)
}

// Add 元素加入列表
func (l *L[T]) Add(elm T) {
	l.arr = append(l.arr, elm)
}

// Get 获取列表中指定位置的元素
func (l *L[T]) Get(idx int) T {
	if idx < 0 || idx >= l.Len() {
		panic("idx out of range")
	}
	return l.arr[idx]
}

// Range 遍历列表元素
func (l *L[T]) Range(handler func(elm T)) {
	for i := 0; i < l.Len(); i++ {
		handler(l.arr[i])
	}
}

// Remove 移除指定索引下的元素
func (l *L[T]) Remove(idx int) {
	if idx < 0 || idx >= l.Len() {
		panic("idx out of range")
	}
	l.arr = append(l.arr[:idx], l.arr[idx+1:]...)
}

// String 返回栈中元素
func (l *L[T]) String() string {
	return fmt.Sprintf("%v", l.arr)
}
