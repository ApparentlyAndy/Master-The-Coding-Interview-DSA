package main

import "fmt"

func main() {
	queue := Queue{}
	queue.enqueue("hello")
	fmt.Println(queue.peek())
	fmt.Println(queue.dequeue())
	fmt.Println(queue.isEmpty())
	queue.printAll()
}
