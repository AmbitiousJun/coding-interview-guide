// 解法1
// 准备一个 qmax 双端队列
// 遍历数组 arr 元素, 对于每一个元素所在位置 i, 判断 qmax 队尾元素 j
//
//	① arr[i] < arr[j]: i 直接入 qmax 队尾
//	② arr[i] >= arr[j]: 不断弹出 qmax 队尾, 直至满足条件 ① 或者 qmax 队空为止
//
// 每次遍历时, 判断 qmax 队首元素 first 是否满足 first == i - w, w 为窗口大小
// 满足则弹出队首
package main

import (
	"coding_interview_guide/common/deque"
	"coding_interview_guide/common/integer"
	"fmt"
)

func main() {
	res := GetMaxWindow([]int{4, 3, 5, 4, 3, 3, 6, 7}, 3)
	fmt.Println("res: ", res)
}

func GetMaxWindow(arr []int, w int) []int {
	if arr == nil || w < 1 || len(arr) < w {
		return nil
	}
	qmax := deque.New[int]()
	res := make([]int, len(arr)-w+1)
	var index integer.I = 0
	for i, num := range arr {
		// 1 qmax 不为空的情况下, 弹出队尾不满足条件的元素
		for !qmax.Empty() && arr[qmax.PeekLast()] <= num {
			qmax.PollLast()
		}
		// 2 当前元素入队
		qmax.AddLast(i)
		// 3 去掉队首的不满足条件元素
		if qmax.PeekFirst() == i-w {
			qmax.PollFirst()
		}
		// 4 当前已经至少遍历了一个窗口大小, 可以开始更新 res
		if i >= w-1 {
			res[index.GetAndIncr()] = arr[qmax.PeekFirst()]
		}
	}
	return res
}
