package main

import (
	"errors"
	"math"
)

type Node struct {
	value    interface{}
	next     *Node
	previous *Node
}

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (list *DoublyLinkedList) traverseTo(index int) (*Node, error) {
	backwards := false
	if list.length == 0 {
		return nil, errors.New("list is empty")
	}

	if index+1 > list.length {
		return nil, errors.New("index out of range")
	}

	var node *Node
	length := index

	if index > int(math.Ceil(float64(list.length/2))) {
		backwards = true
	}

	if backwards {
		node = list.tail
		length = list.length - index - 1
	} else {
		node = list.head
	}

	for i := 0; i < length; i++ {
		if backwards {
			node = node.previous
		} else {
			node = node.next
		}
	}

	return node, nil
}

func (list *DoublyLinkedList) remove(position int) {
	if position < 0 || position > list.length {
		return
	} else if position == 0 {
		list.head = list.head.next
	} else {
		if node, _ := list.traverseTo(position); node != nil {
			previous := node.previous
			next := node.next

			if previous != nil {
				previous.next = next
			}

			if next != nil {
				next.previous = previous
			}
		}
	}
	list.length--
}

func (list *DoublyLinkedList) insert(position int, value interface{}) {
	if position <= 0 {
		list.prepend(value)
		return
	} else if position >= list.length {
		list.append(value)
		return
	}

	newNode := Node{
		value: value,
	}

	if node, _ := list.traverseTo(position); node != nil {
		newNode.previous = node.previous
		node.previous.next = &newNode
		node.previous = &newNode
		newNode.next = node
		list.length++
	}
}

func (list *DoublyLinkedList) append(value interface{}) {
	newNode := Node{
		value: value,
	}

	if list.length == 0 {
		list.head = &newNode
		list.tail = &newNode
		list.length++
		return
	}

	newNode.previous = list.tail
	list.tail.next = &newNode
	list.tail = &newNode
	list.length++
}

func (list *DoublyLinkedList) prepend(value interface{}) {
	if list.length == 0 {
		list.append(value)
		return
	}

	newNode := Node{
		value: value,
		next:  list.head,
	}

	list.head.previous = &newNode
	list.head = &newNode
	list.length++
}

func (list *DoublyLinkedList) reverse() {
	if list.length == 1 {
		return
	}

	firstNode := list.head

	for i := 0; i < list.length; i++ {
		if i == 0 {
			list.append(firstNode.value)
		} else {
			list.insert(list.length-i, firstNode.value)
		}
		list.remove(0)
		firstNode = firstNode.next
	}
}

func (list *DoublyLinkedList) printAll() []interface{} {
	values := []interface{}{}

	if list.length == 0 {
		return values
	}

	node := list.head

	for node != nil {
		values = append(values, node.value)
		node = node.next
	}

	return values
}
