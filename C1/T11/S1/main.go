// 解法1, 进阶问题, 环中可能存在重复值
// 利用单调栈, 通过 "小找大" 的形式依次找出可见山峰对
package main

import (
	"coding_interview_guide/common/stack"
	"fmt"
)

func main() {
	arr := []int{5, 4, 3, 5, 4, 2, 4, 4, 5, 3, 2}
	fmt.Println(GetVisibleNum(arr))
}

// Record 单调栈中存放的实体
type Record struct {
	Val   int
	Times int // 表示同一个 Val 值出现了多少次（不一定是连续的, 但是它们中间夹着的必定是比它们小的数)
}

func NewRecord(val int) *Record {
	return &Record{Val: val, Times: 1}
}

func GetVisibleNum(arr []int) int {
	// 1 数组安全校验
	size := len(arr)
	if size < 2 {
		return 0
	}
	// 2 找出数组中任意一个最大值的索引
	maxIdx := 0
	for i, num := range arr {
		if num > arr[maxIdx] {
			maxIdx = i
		}
	}
	// 3 初始化单调栈, 并压入找到的最大值
	s := stack.New[*Record]()
	s.Push(NewRecord(arr[maxIdx]))
	// 4 准备一个全局索引值, 开始遍历环
	idx := NextIdx(maxIdx, size)
	res := 0
	for idx != maxIdx {
		// 当前值不满足单调栈的压栈规则, 不断弹出进行结算
		for arr[idx] > s.Peek().Val {
			k := s.Pop().Times
			res += 2*k + GetInternalSum(k)
		}
		// 将当前值压入栈中, 需要判断是否与栈顶相等, 并进行合并
		if arr[idx] == s.Peek().Val {
			s.Peek().Times++
		} else {
			s.Push(NewRecord(arr[idx]))
		}
		idx = NextIdx(idx, size)
	}
	// 5 清栈: 第 1 阶段
	for s.Len() > 2 {
		k := s.Pop().Times
		res += 2*k + GetInternalSum(k)
	}
	// 6 清栈: 第 2 阶段
	if s.Len() == 2 {
		k := s.Pop().Times
		if s.Peek().Times <= 1 {
			res += k + GetInternalSum(k)
		} else {
			res += 2*k + GetInternalSum(k)
		}
	}
	// 7 清栈: 第 3 阶段
	res += GetInternalSum(s.Pop().Times)
	return res
}

// GetInternalSum 计算组合数 C(2, K) 的值
func GetInternalSum(k int) int {
	if k <= 1 {
		return 0
	}
	return k * (k - 1) / 2
}

// NextIdx 计算在环中遍历时, 当前位置的下一个位置的索引
func NextIdx(cur, size int) int {
	return (cur + 1) % size
}
