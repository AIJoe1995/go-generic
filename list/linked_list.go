package list

import "go-generic/errs"

type Node[T comparable] struct {
	prev *Node[T]
	next *Node[T]
	val  T
}

func NewNode[T comparable](val T) *Node[T] {
	return &Node[T]{
		val: val,
	}
}

type DoubleLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewDoubleLinkedList[T comparable]() *DoubleLinkedList[T] {
	head := &Node[T]{}
	tail := &Node[T]{}
	head.next = tail
	tail.prev = head
	return &DoubleLinkedList[T]{
		head: head,
		tail: tail,
	}
}

func BuildDoubleLinkedList[T comparable](slice []T) *DoubleLinkedList[T] {
	dll := NewDoubleLinkedList[T]()

	for _, e := range slice {
		err := dll.Append(e)
		if err != nil {
			panic(err)
		}

	}
	return dll
}

func (dll *DoubleLinkedList[T]) ToSlice() []T {
	length := dll.size
	res := make([]T, 0, length)
	node := dll.head.next
	for i := 0; i < length; i++ {
		res = append(res, node.val)
		node = node.next
	}
	return res
}

func (dll *DoubleLinkedList[T]) Append(val T) error {
	node := &Node[T]{val: val}

	node.prev = dll.tail.prev
	dll.tail.prev.next = node
	dll.tail.prev = node
	node.next = dll.tail
	dll.size += 1

	return nil

}

func (dll *DoubleLinkedList[T]) InsertAfter(index int, val T) error {
	if index < 0 || index >= dll.size {
		return errs.NewErrIndexOutOfRange(dll.size, index)
	}
	newnode := &Node[T]{val: val}
	node := dll.head.next
	// insert at tail
	if index == dll.size-1 {
		dll.Append(val)
		return nil
	}
	for cnt := 0; cnt < index; cnt++ {
		node = node.next
	}
	newnode.next = node.next
	node.next.prev = newnode
	node.next = newnode
	newnode.prev = node
	dll.size += 1
	return nil

}

func (dll *DoubleLinkedList[T]) InsertBefore(index int, val T) error {
	if index < 0 || index > dll.size {
		return errs.NewErrIndexOutOfRange(dll.size, index)
	}
	newnode := &Node[T]{val: val}
	node := dll.head.next
	// insert at index 0
	if index == 0 {
		newnode.next = node
		dll.head.next = newnode
		newnode.prev = dll.head
		node.prev = newnode
		dll.size += 1
		return nil
	}
	// insert at tail
	if index == dll.size {
		dll.Append(val)
		return nil
	}
	//
	for cnt := 0; cnt < index; cnt++ {
		node = node.next
	}

	newnode.prev = node.prev
	newnode.next = node
	node.prev.next = newnode
	node.prev = newnode
	dll.size += 1
	return nil

}

func (dll *DoubleLinkedList[T]) Remove(val T) (int, bool) {
	node := dll.head.next
	res := -1

	for i := 0; i < dll.size; i++ {
		if node.val == val {
			res = i
			break
		}
		node = node.next
	}
	// remove head
	if res == 0 {
		dll.head.next = node.next
		node.next.prev = dll.head
		node.next = nil
		dll.size -= 1
		//todo  delete node release memory
	} else if res == dll.size-1 { // remove tail
		dll.tail = node.prev
		node.prev.next = nil
		node.prev = nil
		dll.size -= 1
		// todo delete node release memory
	} else if res == -1 { // not found
		return res, false
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
		dll.size -= 1
		// todo delete node release memory
	}
	return res, true
}

func (dll *DoubleLinkedList[T]) RemoveAt(index int) (T, error) {
	var res T
	if index < 0 || index >= dll.size {
		return res, errs.NewErrIndexOutOfRange(dll.size, index)
	}
	// todo delete node release memory
	node := dll.head.next
	// remove head
	if index == 0 {
		res = node.val
		dll.head.next = node.next
		node.next.prev = dll.head
		node.next.prev = nil
		node.next = nil
		dll.size -= 1
	} else if index == dll.size-1 { // remove tail
		node = dll.tail.prev
		res = node.val
		dll.tail = node.prev
		node.prev.next = nil
		node.prev = nil
		dll.size -= 1
	} else {
		node = dll.head.next
		for cnt := 0; cnt < index; cnt++ {
			node = node.next
		}
		res = node.val
		node.prev.next = node.next
		node.next.prev = node.prev
		dll.size -= 1
	}
	return res, nil

}

func (dll *DoubleLinkedList[T]) Get(index int) (*Node[T], error) {
	if index < 0 || index >= dll.size {
		return nil, errs.NewErrIndexOutOfRange(dll.size, index)
	}
	node := dll.head.next
	for cnt := 0; cnt < dll.size; cnt++ {
		if cnt == index {
			return node, nil
		}
		node = node.next
	}
	return nil, errs.NewSystemError()
}

func (dll *DoubleLinkedList[T]) IsEmpty() bool {
	return dll.size == 0
}

func (dll *DoubleLinkedList[T]) IndexOf(val T) int {
	node := dll.head.next
	for i := 0; i < dll.size; i++ {
		if node.val == val {
			return i
		}
		node = node.next
	}
	return -1
}
