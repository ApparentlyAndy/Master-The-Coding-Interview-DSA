package main

type Node struct {
	value interface{}
	next  *Node
}

type Stack struct {
	top    *Node
	bottom *Node
	length int
}

func (stack *Stack) push(value interface{}) {
	newNode := Node{
		value: value,
	}

	if stack.length == 0 {
		stack.top = &newNode
		stack.bottom = &newNode
	} else {
		stack.bottom.next = &newNode
		stack.bottom = &newNode
	}

	stack.length++
}

func (stack *Stack) peek() interface{} {
	return stack.bottom.value
}

func (stack *Stack) pop() interface{} {
	var node *Node
	var result interface{}

	if stack.length == 0 {
		return nil
	} else if stack.length == 1 {
		result = stack.top.value
		stack.top = nil
		stack.bottom = nil
	} else {
		node = stack.top
		for node.next.next != nil {
			node = node.next
		}

		node.next = nil
		result = stack.bottom.value
		stack.bottom = node
	}
	stack.length--
	return result
}
