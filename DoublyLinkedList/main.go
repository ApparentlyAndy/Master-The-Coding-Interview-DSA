package main

import "fmt"

func main() {
	linkedList := DoublyLinkedList{}

	linkedList.append(0)
	linkedList.append(1)

	linkedList.reverse()
	fmt.Println(linkedList.printAll())
	fmt.Println(linkedList.length)
}
