// 解法1
//
// 类似于 T7 题目，维护 qmax 和 qmin 两个双端队列
// 在右边界往右扩时，判断如果 qmax 和 qmin 的左端点不符合题目条件，就进行子数组结果结算
package main

import (
	"coding_interview_guide/common/deque"
	"fmt"
)

func main() {
	arr := []int{3, 5, 9, 6, 2, 4}
	fmt.Println(GetNum(arr, 2))
}

// GetNum 输入一个数组 arr 和一个目标值 num
// 返回符合条件的子数组区间个数
func GetNum(arr []int, num int) int {
	if len(arr) == 0 || num < 0 {
		return 0
	}
	res := 0
	qmin, qmax := deque.New[int](), deque.New[int]()
	i, j := 0, 0
	for i < len(arr) {
		for j < len(arr) {
			if qmin.Empty() || qmin.PeekLast() != j {
				for !qmin.Empty() && arr[qmin.PeekLast()] >= arr[j] {
					qmin.PollLast()
				}
				qmin.AddLast(j)
				for !qmax.Empty() && arr[qmax.PeekLast()] <= arr[j] {
					qmax.PollLast()
				}
				qmax.AddLast(j)
			}
			if arr[qmax.PeekFirst()]-arr[qmin.PeekFirst()] > num {
				// 右边界扩到第一个不合法的位置, 退出循环进行结算
				break
			}
			j++
		}
		res += j - i
		// 维护 qmax 和 qmin 的左端点
		if qmax.PeekFirst() == i {
			qmax.PollFirst()
		}
		if qmin.PeekFirst() == i {
			qmin.PollFirst()
		}
		i++
	}
	return res
}
