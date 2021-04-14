package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

var (
	keyFound bool
)

func (tree *node) isNil() bool {
	if tree.data == 0 && tree.left == nil && tree.right == nil {
		return true
	} else {
		return false
	}
}
func (tree *node) search(val int) bool {
	keyFound = false
	if tree == nil {
	} else if tree.data == val {
		keyFound = true
	} else if val < tree.data {
		tree = tree.left
		tree.search(val)
	} else if val > tree.data {
		tree = tree.right
		tree.search(val)
	}
	return keyFound
}

func (tree *node) findMin() node {
	if tree.left != nil {
		tree = tree.left
	}
	return *tree.left
}
func (tree *node) findMax() node {
	if tree.right != nil {
		tree = tree.right
	}
	return *tree.right
}

func (tree *node) insert(val int) {

	if tree.isNil() {
		*tree = node{val, nil, nil}
	} else if val < tree.data {
		if tree.left == nil {
			childNode := node{val, nil, nil}
			tree.left = &childNode
		} else {
			tree = tree.left
			tree.insert(val)
		}
	} else if val > tree.data {
		if tree.right == nil {
			childNode := node{val, nil, nil}
			tree.right = &childNode
		} else {
			tree = tree.right
			tree.insert(val)
		}
	}
}

func (tree *node) inOrderTreeWalk() {

	if tree != nil {
		tree.left.inOrderTreeWalk()
		fmt.Println(*tree)
		tree.right.inOrderTreeWalk()
	}
}

func (tree *node) preOrderTreeWalk() {

	if tree != nil {
		fmt.Println(*tree)
		tree.left.preOrderTreeWalk()
		tree.right.preOrderTreeWalk()
	}
}

func (tree *node) postOrderTreeWalk() {

	if tree != nil {
		tree.left.postOrderTreeWalk()
		tree.right.postOrderTreeWalk()
		fmt.Println(*tree)
	}
}

func main() {

	// var (
	// 	childH = node{9, nil, nil}
	// 	childG = node{13, &childH, nil}
	// 	childE = node{17, nil, nil}
	// 	childF = node{20, nil, nil}
	// 	childC = node{3, nil, nil}
	// 	childD = node{7, nil, &childG}
	// 	childB = node{18, &childE, &childF}
	// 	childA = node{6, &childC, &childD}
	// 	root   = node{15, &childA, &childB}
	// )

	root := node{}
	root.insert(15)
	root.insert(6)
	root.insert(3)
	root.insert(18)
	root.insert(7)
	root.insert(20)
	root.insert(17)
	root.insert(13)
	root.insert(9)

	fmt.Println("\\\\inOrderTreeWalk\\\\")
	root.inOrderTreeWalk()
	fmt.Println("\\\\preOrderTreeWalk\\\\")
	root.preOrderTreeWalk()
	fmt.Println("\\\\postOrderTreeWalk\\\\")
	root.postOrderTreeWalk()

}
