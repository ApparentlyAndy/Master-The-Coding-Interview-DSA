package main

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

type Queue struct {
	first  *Node
	last   *Node
	length int
}

func (q *Queue) peek() interface{} {
	if q.length == 0 {
		return nil
	}
	return q.first.value
}

func (q *Queue) enqueue(value interface{}) {
	newNode := Node{
		value: value,
	}
	if q.length == 0 {
		q.first = &newNode
		q.last = &newNode
		q.length++
		return
	}
	newNode.next = q.first
	q.first = &newNode
	q.length++
}

func (q *Queue) dequeue() interface{} {
	if q.length == 0 {
		return nil
	}

	result := q.first.value

	if q.first == q.last {
		q.last = nil
	}

	q.first = q.first.next
	q.length--

	return result
}

func (q *Queue) isEmpty() bool {
	return q.length == 0
}

func (q *Queue) printAll() {
	var node *Node
	var results []interface{}
	for i := 0; i < q.length; i++ {
		if i == 0 {
			node = q.first
		} else {
			node = node.next
		}
		results = append(results, node.value)
	}
	fmt.Println(results)
}
