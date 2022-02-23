package main

import "fmt"

func main() {
	myStack := Stack{}
	myStack.push("hello")
	myStack.push("hello1")
	myStack.push("hello2")
	fmt.Println(myStack.peek())
	fmt.Println(myStack.pop())
	fmt.Println(myStack.peek())
	fmt.Println(myStack.pop())
	fmt.Println(myStack.pop())
	fmt.Println(myStack.pop())
	fmt.Println(myStack.peek())
}
