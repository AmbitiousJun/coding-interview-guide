// 解法1
// 直接在遍历 arr 的时候分别向左和向右遍历即可确定答案, 时间复杂度是 O(n²)
package main

import "fmt"

func main() {
	arr := []int{3, 4, 1, 5, 6, 2, 7}
	res := RightWay(arr)
	fmt.Println(res)
}

func RightWay(arr []int) [][]int {
	res := make([][]int, len(arr))
	for i, num := range arr {
		leftIdx, rightIdx := -1, -1
		cur := i - 1
		for cur >= 0 {
			if arr[cur] < num {
				leftIdx = cur
				break
			}
			cur--
		}
		cur = i + 1
		for cur < len(arr) {
			if arr[cur] < num {
				rightIdx = cur
				break
			}
			cur++
		}
		res[i] = []int{leftIdx, rightIdx}
	}
	return res
}
