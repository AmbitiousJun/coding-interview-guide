package list

// Node 表示链表的每一个节点
type Node[T any] struct {
	Val  T
	Next *Node[T]
}

// NewLinkNode 创建一个链表节点
func NewLinkNode[T any](val T) *Node[T] {
	return &Node[T]{Val: val}
}

// NewLinkFromArray 使用数组初始化一个链表结构, 返回头节点
func NewLinkFromArray[T any](arr []T) *Node[T] {
	if len(arr) == 0 {
		return nil
	}
	var head, tail *Node[T] = nil, nil

	for _, elm := range arr {
		node := &Node[T]{Val: elm}
		if head == nil {
			head, tail = node, node
			continue
		}
		tail.Next = node
		tail = node
	}

	return head
}
