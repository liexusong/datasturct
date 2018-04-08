package datastruct

import (
	"fmt"
)

type ListNode struct {
	prev  *ListNode
	next  *ListNode
	value interface{}
}

type List struct {
	head   *ListNode
	tail   *ListNode
	length uint64
}

func NewList() *List {
	return &List{nil, nil, 0}
}

func (this *List) Add(value interface{}) {

	node := &ListNode{nil, nil, value}

	node.prev = this.tail

	if this.tail != nil {
		this.tail.next = node
	}

	this.tail = node

	if this.head == nil {
		this.head = this.tail
	}

	this.length++
}

func (this *List) Remove(index int) (interface{}, error) {

	node := this.head

	for i := 0; i < index; i++ {
		if node != nil {
			node = node.next
		} else {
			break
		}
	}

	if node == nil {
		return nil, fmt.Errorf("Index out of range")
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		this.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		this.tail = node.prev
	}

	this.length--

	return node.value, nil
}

func (thsis *List) Get(index int) (interface{}, error) {

	node := this.head

	for i := 0; i < index; i++ {
		if node != nil {
			node = node.next
		} else {
			break
		}
	}

	if node == nil {
		return nil, fmt.Errorf("Index out of range")
	}

	return node.value, nil
}

func (this *List) Length() uint64 {
	return this.length
}
