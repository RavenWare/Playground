package main

import "errors"

var (
	errValueNotFound = errors.New("value not found")
)

func (tree *node) search(value int) (*node, error) {
	if tree == nil {
		return nil, errValueNotFound
	} else if tree.data == value {
		return tree, nil
	} else if value < tree.data {
		return tree.left.search(value)
	} else {
		return tree.right.search(value)
	}
}
