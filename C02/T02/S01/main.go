// 解法1
// 处理单链表
package main

import (
	"coding_interview_guide/common/list"
	"log"
)

func main() {
	arr := []int{1, 2, 1, 3, 8}
	head := list.NewLinkFromArray(arr)
	log.Println("current head: ", head)
	head = RemoveLastKthNode(head, 3)
	log.Println("after remove last 3th node: ", head)
}

// RemoveLastKthNode 移除链表倒数第 K 个节点, 并返回新的链表
func RemoveLastKthNode(head *list.Node[int], lastKth int) *list.Node[int] {
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
	}
	if lastKth < 0 {
		cur = head
		lastKth++
		for lastKth != 0 {
			cur = cur.Next
			lastKth++
		}
		cur.Next = cur.Next.Next
	}
	return head
}
