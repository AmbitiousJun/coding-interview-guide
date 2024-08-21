// 解法1
package main

import (
	"coding_interview_guide/common/list"
	"fmt"
)

func main() {
	arr1, arr2 := []int{1, 3, 5, 7, 8, 9}, []int{2, 3, 6, 8}
	head1, head2 := list.NewLinkFromArray(arr1), list.NewLinkFromArray(arr2)
	PrintCommonPart(head1, head2)
}

func PrintCommonPart(head1, head2 *list.Node[int]) {
	if head1 == nil || head2 == nil {
		return
	}
	fmt.Print("Common part: ")
	for head1 != nil && head2 != nil {
		if head1.Val > head2.Val {
			head2 = head2.Next
		} else if head1.Val < head2.Val {
			head1 = head1.Next
		} else {
			fmt.Print(head1.Val, " ")
			head1, head2 = head1.Next, head2.Next
		}
	}
	fmt.Println()
}
