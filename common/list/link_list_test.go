package list_test

import (
	"coding_interview_guide/common/list"
	"log"
	"testing"
)

func TestInitLinkList(t *testing.T) {
	arr := []int{3, 5, 3, 2, 1, 2, 1, 3, 8, 9}
	head := list.NewLinkFromArray[int](arr)
	for head != nil {
		log.Println(head.Val)
		head = head.Next
	}
}
