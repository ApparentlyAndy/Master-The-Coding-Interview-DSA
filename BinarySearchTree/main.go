package main

// func randInt() int {
// 	rand.Seed(time.Now().UnixNano())
// 	min := 1
// 	max := 100
// 	return (rand.Intn(max-min+1) + min)
// }

func main() {
	bst := BST{}
	bst.insert(80)
	bst.insert(95)
	bst.insert(51)
	bst.insert(33)
	bst.insert(49)
	bst.insert(60)
	bst.insert(65)
	bst.insert(64)
	bst.insert(63)
	bst.insert(61)
	bst.insert(70)
	bst.insert(69)
	bst.insert(68)
	bst.insert(72)
	bst.insert(71)

	bst.remove(70)
	bst.printNode(bst.root)
	// counter := 0

	// for counter < 20 {
	// 	bst.insert(randInt())
	// 	counter++
	// }
}
