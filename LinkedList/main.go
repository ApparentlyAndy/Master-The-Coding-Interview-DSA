package main

import (
	"errors"
	"fmt"
)

type ListNode struct {
	value interface{}
	next  *ListNode
}

type LinkedList struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func (list *LinkedList) traverseBefore(position int) (*ListNode, error) {
	if position < 0 {
		return nil, errors.New("position cannot be lower than 0")
	} else if position+1 > list.length {
		return nil, errors.New("position exceeds list")
	}

	if position == 0 {
		return list.head, nil
	}

	var previousNode *ListNode

	for i := 0; i < position; i++ {
		if i == 0 {
			previousNode = list.head
		} else {
			previousNode = previousNode.next
		}
	}

	return previousNode, nil
}

func (list *LinkedList) append(value interface{}) {
	node := ListNode{
		value: value,
	}

	if list.length == 0 {
		list.head = &node
		list.tail = &node
	} else {
		list.tail.next = &node
		list.tail = &node
	}
	list.length++
}

func (list *LinkedList) prepend(value interface{}) {
	node := ListNode{
		value: value,
	}

	if list.length == 0 {
		list.head = &node
		list.tail = &node
	} else {
		node.next = list.head
		list.head = &node
	}

	list.length++
}

func (list *LinkedList) remove(position int) {
	if position < 0 {
		return
	} else if position == 0 {
		list.head = list.head.next
		return
	}

	previousNode, err := list.traverseBefore(position)

	if err == nil {
		previousNode.next = previousNode.next.next
	}
}

func (list *LinkedList) insert(position int, value interface{}) {
	newNode := ListNode{
		value: value,
	}

	if position < 0 {
		return
	} else if position == 0 {
		list.prepend(value)
		return
	} else if position >= list.length {
		list.append(value)
		return
	}

	previousNode, err := list.traverseBefore(position)

	if err == nil {
		newNode.next = previousNode.next
		previousNode.next = &newNode
		list.length++
	}

}

func (list *LinkedList) displayAll() {
	if list.head != nil {
		node := list.head
		for node.next != nil {
			fmt.Println(node.value)
			node = node.next
		}

		if node.value != nil {
			fmt.Println(node.value)
		}
	}
}

func main() {
	linkedList := LinkedList{}

	linkedList.append("0")

	linkedList.append("1")

	linkedList.append("2")

	linkedList.append("3")

	linkedList.insert(1, "1")

	linkedList.displayAll()
}
