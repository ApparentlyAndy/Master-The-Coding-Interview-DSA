package main

type Stack struct {
	data   []interface{}
	length int
}

func (stack *Stack) push(data interface{}) {
	stack.data = append(stack.data, data)
	stack.length++
}

func (stack *Stack) peek() interface{} {
	if stack.length == 0 {
		return nil
	}
	return stack.data[stack.length-1]
}

func (stack *Stack) pop() interface{} {
	if stack.length == 0 {
		return nil
	}
	result := stack.data[stack.length-1]
	stack.data = stack.data[:stack.length-1]
	stack.length--
	return result
}
