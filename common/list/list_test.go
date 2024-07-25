package list_test

import (
	"coding_interview_guide/common/list"
	"fmt"
	"log"
	"testing"
)

func TestRemove(t *testing.T) {
	l := list.New(0)
	l.Add(1)
	l.Add(2)
	l.Add(1)
	l.Add(3)
	l.Add(8)
	fmt.Println(l, l.Len())
	l.Remove(2)
	fmt.Println(l, l.Len())
	l.Range(func(elm int) {
		log.Println(elm)
	})
}
