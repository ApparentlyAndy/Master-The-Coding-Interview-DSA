package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (bst *BST) printNode(node *Node) interface{} {
	if node != nil {
		var left interface{}
		var right interface{}

		if node.left != nil {
			left = node.left.value
		}

		if node.right != nil {
			right = node.right.value
		}

		fmt.Printf("Value: %v, Left: %v, Right: %v\n", node.value, left, right)
	}

	if node.left != nil {
		bst.printNode(node.left)
	}

	if node.right != nil {
		bst.printNode(node.right)
	}

	return nil
}

func (bst *BST) traverse(value int) (*Node, error) {
	if bst.root == nil {
		return nil, errors.New("no root in BST")
	}

	found := false

	node := bst.root

	for !found {
		if value == node.value {
			return nil, errors.New("duplicate key")
		}

		if value < node.value && node.left != nil {
			node = node.left
			continue
		}

		if value > node.value && node.right != nil {
			node = node.right
			continue
		}

		found = true

	}
	return node, nil
}

func (bst *BST) insert(value int) {
	newNode := Node{
		value: value,
	}

	if bst.root == nil {
		bst.root = &newNode
	} else {
		if node, _ := bst.traverse(value); node != nil {
			if newNode.value < node.value {
				node.left = &newNode
			} else if newNode.value > node.value {
				node.right = &newNode
			}
		}

	}
}

func (bst *BST) lookup(value int) bool {
	if node, _ := bst.search(value); node != nil {
		return true
	}
	return false
}

func (bst *BST) search(value int) (*Node, error) {
	currentNode := bst.root
	found := false

	for !found {
		if currentNode == nil {
			return nil, errors.New("no root in tree")
		} else {
			if value == currentNode.value {
				return currentNode, nil
			} else if value < currentNode.value && currentNode.left != nil {
				currentNode = currentNode.left
				continue
			} else if value > currentNode.value && currentNode.right != nil {
				currentNode = currentNode.right
				continue
			}
			found = true
		}
	}
	return nil, errors.New("cannot find value")
}

func (bst *BST) remove(value int) {
	if bst.root == nil {
		return
	}

	var parent *Node
	current := bst.root
	found := false

	for !found {
		if value == current.value {
			break
		}
		if value < current.value && current.left != nil {
			parent = current
			current = current.left
			continue
		} else if value > current.value && current.right != nil {
			parent = current
			current = current.right
			continue
		}
		found = true
	}

	if current.value == value {
		if current.right == nil {
			if parent == nil {
				bst.root = current.left
			} else {
				if current.value < parent.value {
					parent.left = current.left
				} else if current.value > parent.value {
					parent.right = current.left
				}
			}
		} else if current.right.left == nil {
			current.right.left = current.left
			if parent == nil {
				bst.root = current.right
			} else {
				if current.value < parent.value {
					parent.left = current.right
				} else if current.value > parent.value {
					parent.right = current.right
				}
			}
		} else {
			leftMost := current.right.left
			leftMostParent := current.right

			for leftMost.left != nil {
				leftMostParent = leftMost
				leftMost = leftMost.left
			}

			leftMostParent.left = leftMost.right
			leftMost.left = current.left
			leftMost.right = current.right

			if parent == nil {
				bst.root = leftMost
			} else {
				if current.value < parent.value {
					parent.left = leftMost
				} else if current.value > parent.value {
					parent.right = leftMost
				}
			}
		}
	}

}
