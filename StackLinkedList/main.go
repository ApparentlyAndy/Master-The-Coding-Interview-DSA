package main

import "fmt"

func main() {
	stack := Stack{}
	stack.push("hello")
	stack.push("hello1")
	stack.push("hello2")
	stack.push("hello3")
	stack.push("hello4")
	fmt.Println(stack.peek())
	fmt.Println(stack.pop(), stack.length)
}
