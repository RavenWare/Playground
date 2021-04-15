package main

func (tree *node) insert(value int) {
	if value < tree.data {
		if tree.left == nil {
			tree.left = &node{value, nil, nil}
		} else {
			tree.left.insert(value)
		}
	} else if value > tree.data {
		if tree.right == nil {
			tree.right = &node{value, nil, nil}
		} else {
			tree.right.insert(value)
		}
	}
}
