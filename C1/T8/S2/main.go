// 解法2
//
// 使用单调栈, 解决原始问题 (数组 arr 中元素不重复)
//
// 核心思路, 初始化一个栈, 并将其维护成一个从顶到底递减的序列
//
// 具体的, 在遍历 arr 的时候, 对于每个 i, 和栈顶 x 比较
//
//	① arr[i] > arr[x]: i 入栈
//	② arr[i] < arr[x]: x 弹出, 并且 i 为 x 在 res[x] 中的右值, 此时的栈顶为 x 在 res[x] 中的左值
//
// arr 遍历完成的时候, s 栈不为空, 需要进行清栈
// 此时只需要依次弹出栈顶元素 x, x 在 res[x] 中的右值必定是 -1, res[x] 的左值为 s 当前的栈顶 (栈空时为 -1)
package main

import (
	"coding_interview_guide/common/stack"
	"fmt"
)

func main() {
	arr := []int{3, 4, 1, 5, 6, 2, 7}
	res := GetNearLessNoRepeat(arr)
	fmt.Println(res)
}

func GetNearLessNoRepeat(arr []int) [][]int {
	res := make([][]int, len(arr))
	s := stack.New[int]()
	for i, num := range arr {
		for !s.Empty() && arr[s.Peek()] > num {
			UpdateResult(res, s, i)
		}
		s.Push(i)
	}
	for !s.Empty() {
		UpdateResult(res, s, -1)
	}
	return res
}

// UpdateResult 更新 res 数组, x 值为 s 栈顶
func UpdateResult(res [][]int, s *stack.S[int], rightIdx int) {
	if s.Empty() {
		return
	}
	cur := s.Pop()
	leftIdx := -1
	if !s.Empty() {
		leftIdx = s.Peek()
	}
	res[cur] = []int{leftIdx, rightIdx}
}
