// 解法1
// 使用递归思想
package main

import "fmt"

func main() {
	n := 15
	fmt.Println("移动", n, "层汉诺塔使用总步数：", HanoiProblem(n, "左", "中", "右"))
}

// HanoiProblem 汉诺塔问题
// num 层数
// 返回值: num 层汉诺塔从 left 移动到 right 所需要移动的总步数
func HanoiProblem(num int, left, mid, right string) int {
	if num < 1 {
		return 0
	}
	return process(num, left, mid, right, left, right)
}

// process 核心递归函数
// 将 num 层塔从 from 移动到 to, 返回需要移动的总步数
func process(num int, left, mid, right, from, to string) int {
	// base case: 只需移动一层塔
	if num == 1 {
		if from == mid || to == mid {
			fmt.Printf("Move 1 from %s to %s\n", from, to)
			return 1
		}
		fmt.Printf("Move 1 from %s to %s\n", from, mid)
		fmt.Printf("Move 1 from %s to %s\n", mid, to)
		return 2
	}
	if from == mid || to == mid {
		// 另一个可以用来暂存的位置
		another := right
		if from == left || right == left {
			another = left
		}
		part1 := process(num-1, left, mid, right, from, another)
		fmt.Printf("Move %d from %s to %s\n", num, from, to)
		part3 := process(num-1, left, mid, right, another, to)
		return part1 + 1 + part3
	}
	part1 := process(num-1, left, mid, right, from, to)
	fmt.Printf("Move %d from %s to %s\n", num, from, mid)
	part3 := process(num-1, left, mid, right, to, from)
	fmt.Printf("Move %d from %s to %s\n", num, mid, to)
	part5 := process(num-1, left, mid, right, from, to)
	return part1 + 1 + part3 + 1 + part5
}
