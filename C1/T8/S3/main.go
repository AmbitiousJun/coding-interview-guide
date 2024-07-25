// 解法3
//
// 解决进阶问题, arr 中可能存在重复元素
//
// 思路与解法2是一致的, 只不过栈中存放的元素不再是一个整型数, 而是一个 list 结构, 所有重复的值都看作一个整体存放在同一个 list 下
package main

import (
	"coding_interview_guide/common/list"
	"coding_interview_guide/common/stack"
	"fmt"
)

func main() {
	arr := []int{3, 1, 3, 4, 3, 5, 3, 2, 2}
	res := GetNearLess(arr)
	fmt.Println(res)
}

func GetNearLess(arr []int) [][]int {
	res := make([][]int, len(arr))
	s := stack.New((*list.L[int])(nil))
	for i, num := range arr {
		for !s.Empty() && arr[s.Peek().Get(0)] > num {
			UpdateResult(res, s, i)
		}
		// 当前索引入栈
		if !s.Empty() && arr[s.Peek().Get(0)] == num {
			// 之前已经有相同元素入过栈, 复用之前的 list
			s.Peek().Add(i)
		} else {
			// 新建 list
			l := list.New(0)
			l.Add(i)
			s.Push(l)
		}
	}
	// 清栈
	for !s.Empty() {
		UpdateResult(res, s, -1)
	}
	return res
}

// UpdateResult 更新 res 数组, x 值为 s 栈顶
func UpdateResult(res [][]int, s *stack.S[*list.L[int]], rightIdx int) {
	if s.Empty() {
		return
	}
	l := s.Pop()
	// 确定 leftIdx, 使用 list 最右侧的值, 才是最靠近 x 的
	leftIdx := -1
	if !s.Empty() {
		leftIdx = s.Peek().Get(s.Peek().Len() - 1)
	}
	l.Range(func(elm int) {
		res[elm] = []int{leftIdx, rightIdx}
	})
}
