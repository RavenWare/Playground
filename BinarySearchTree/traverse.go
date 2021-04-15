package main

import "fmt"

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
