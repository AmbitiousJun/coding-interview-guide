// 解法1
//
// 从上往下依次遍历矩阵的每一行, 维护好一个 height 数组
// 在每一行维护好 height 数组后, 计算出以当前行为底, 往上扫描能得到的最大矩形大小
// 每一次得到的最大矩形大小都跟全局结果比较取较大值, 最终得到的就是整个矩阵能得到的最大矩形大小
package main

import (
	"coding_interview_guide/common/stack"
	"coding_interview_guide/util/math"
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 0, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 0},
	}
	fmt.Println(MaxRecSize(matrix))
}

// 传入一个 N * M 的矩阵, 返回最大的全是 1 的矩形区域中 1 的数量
func MaxRecSize(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	height := make([]int, len(matrix[0]))
	maxArea := 0
	for _, row := range matrix {
		for j, num := range row {
			if num == 0 {
				height[j] = 0
			} else {
				height[j]++
			}
		}
		maxArea = math.Max(maxArea, MaxRecFromBottom(height))
	}
	return maxArea
}

// 根据以某一行为底计算出来的 height 数组, 计算出往上扫描的符合条件的最大矩形 1 的数量
func MaxRecFromBottom(height []int) int {
	if len(height) == 0 {
		return 0
	}
	s := stack.New[int]()
	maxArea := 0
	for i, num := range height {
		for !s.Empty() && height[s.Peek()] >= num {
			j := s.Pop()
			k := -1
			if !s.Empty() {
				k = s.Peek()
			}
			maxArea = math.Max(maxArea, (i-k-1)*height[j])
		}
		s.Push(i)
	}
	// 清栈
	for !s.Empty() {
		j := s.Pop()
		k := -1
		if !s.Empty() {
			k = s.Peek()
		}
		maxArea = math.Max(maxArea, (len(height)-k-1)*height[j])
	}
	return maxArea
}
