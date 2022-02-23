package main

import "fmt"

// func randInt() int {
// 	rand.Seed(time.Now().UnixNano())
// 	min := 1
// 	max := 100
// 	return (rand.Intn(max-min+1) + min)
// }

func main() {
	bst := BST{}
	bst.insert(9)
	bst.insert(4)
	bst.insert(6)
	bst.insert(20)
	bst.insert(170)
	bst.insert(15)
	bst.insert(1)
	bst.printNode(bst.root)

	fmt.Println(bst.breadthFirstSearch())
	// counter := 0

	// for counter < 20 {
	// 	bst.insert(randInt())
	// 	counter++
	// }
}
