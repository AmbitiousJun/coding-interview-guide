// 解法2
// 处理双链表
package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	arr := []int{1, 2, 1, 3, 8}
	head := Arr2DoubleList(arr)
	log.Println("移除倒数第 2 个元素之前: ", head)
	head = RemoveLastKthNode(head, 2)
	log.Println("移除倒数第 2 个元素之后: ", head)
}

// RemoveLastKthNode 移除倒数第 lastKth 个元素, 返回移除完成后的双链表
func RemoveLastKthNode(head *DoubleNode, lastKth int) *DoubleNode {
	if head == nil || lastKth < 1 {
		return head
	}
	cur := head
	for cur != nil {
		lastKth--
		cur = cur.Next
	}
	if lastKth == 0 {
		head = head.Next
		head.Last = nil
	}
	if lastKth < 0 {
		cur = head
		lastKth++
		for lastKth != 0 {
			cur = cur.Next
			lastKth++
		}
		newNext := cur.Next.Next
		cur.Next = newNext
		if newNext != nil {
			newNext.Last = cur
		}
	}
	return head
}

// DoubleNode 双链表节点
type DoubleNode struct {

	// Val 节点值
	Val int

	// Next 下一个节点
	Next *DoubleNode

	// Last 上一个节点
	Last *DoubleNode
}

func (dn *DoubleNode) String() string {
	cur := dn
	sb := strings.Builder{}
	// 输出正方向
	sb.WriteString("[\n正: " + strconv.Itoa(cur.Val))
	for cur.Next != nil {
		cur = cur.Next
		sb.WriteString(" -> ")
		sb.WriteString(strconv.Itoa(cur.Val))
	}
	sb.WriteString("\n反: " + strconv.Itoa(cur.Val))
	for cur.Last != nil {
		cur = cur.Last
		sb.WriteString(" -> ")
		sb.WriteString(strconv.Itoa(cur.Val))
	}
	sb.WriteString("\n]")
	return sb.String()
}

// Arr2DoubleList 数组转双链表
func Arr2DoubleList(arr []int) *DoubleNode {
	if len(arr) == 0 {
		return nil
	}
	head := &DoubleNode{Val: arr[0]}
	// 队尾节点
	tail := head
	for i := 1; i < len(arr); i++ {
		cur := &DoubleNode{Val: arr[i], Last: tail}
		tail.Next = cur
		tail = cur
	}
	return head
}
