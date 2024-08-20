// 解法2
// 使用 3 个栈进行模拟
package main

import (
	"coding_interview_guide/common/stack"
	"fmt"
	"math"
)

const (
	ActionNo = iota
	ActionLToM
	ActionMToL
	ActionMToR
	ActionRToM
)

func main() {
	n := 14
	fmt.Println("移动", n, "层汉诺塔使用总步数：", HanoiProblem(n, "左", "中", "右"))
}

// HanoiProblem 汉诺塔问题
// num 层数
// 返回值: num 层汉诺塔从 left 移动到 right 所需要移动的总步数
func HanoiProblem(num int, left, mid, right string) int {
	if num < 1 {
		return 0
	}
	// 初始化 3 个栈并加入三个绝对最大值, 方便代码运行
	lS, mS, rS := stack.New[int](), stack.New[int](), stack.New[int]()
	lS.Push(math.MaxInt32)
	mS.Push(math.MaxInt32)
	rS.Push(math.MaxInt32)
	for i := num; i > 0; i-- {
		lS.Push(i)
	}
	// 初始时, 没有进行任何动作
	record := new(int)
	*record = ActionNo
	step := 0
	for rS.Len() < num+1 {
		step += fStackToTStack(record, ActionMToL, ActionLToM, lS, mS, left, mid)
		step += fStackToTStack(record, ActionLToM, ActionMToL, mS, lS, mid, left)
		step += fStackToTStack(record, ActionRToM, ActionMToR, mS, rS, mid, right)
		step += fStackToTStack(record, ActionMToR, ActionRToM, rS, mS, right, mid)
	}
	return step
}

// 尝试将塔从 from 栈移动到 to 栈
//
// actRecord 记录上一步的动作
//
// preNotAct, nowAct 上一步不能是什么动作（遵循相邻不可逆原则）以及当前要进行什么动作
//
// fStack, tStack from 和 to 两个地方代表的栈
func fStackToTStack(actRecord *int, preNotAct, nowAct int, fStack, tStack *stack.S[int], from, to string) int {
	if preNotAct == *actRecord || fStack.Peek() > tStack.Peek() {
		// 违反相邻不可逆, 或者违反小压大原则, 这一步不能走
		return 0
	}
	tStack.Push(fStack.Pop())
	fmt.Printf("Move %d from %s to %s\n", tStack.Peek(), from, to)
	*actRecord = nowAct
	return 1
}
