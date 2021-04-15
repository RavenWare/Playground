package main

type node struct {
	data  int
	left  *node
	right *node
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

func (tree *node) delete(value int) {

}

func main() {

	// We declare root of the tree
	root := node{15, nil, nil}
	root.insert(10)
	root.insert(20)
	root.insert(8)
	root.insert(12)
	root.insert(18)
	root.insert(25)

	// root.inOrderTreeWalk()

	// node, err := root.search(9)
	// fmt.Println(*node, err)

	root.delete(8)
	// root.delete(25)
	// fmt.Println(node)
	root.inOrderTreeWalk()
	// node, err = root.search(15)
	// fmt.Println(node, err)

	// root.inOrderTreeWalk()

	// node, err = root.search(9)
	// fmt.Println(*node, err)

	// root.postOrderTreeWalk()
	// root.preOrderTreeWalk()

}
