package main

type Tree struct {
	root *binaryNode
}

type binaryNode struct {
	left  *binaryNode
	right *binaryNode
	val   int
}

func (t *Tree) PreOrderSearch() []int {
	return preWalk(t.root, []int{})

}

func preWalk(curr *binaryNode, path []int) []int {
	if curr == nil {
		return path
	}
	path = append(path, curr.val)
	preWalk(curr.left, path)
	preWalk(curr.right, path)
	return path
}

func (t *Tree) InOrderSearch() []int {
	return inWalk(t.root, []int{})
}

func inWalk(curr *binaryNode, path []int) []int {
	if curr == nil {
		return path
	}
	inWalk(curr.left, path)
	path = append(path, curr.val)
	inWalk(curr.left, path)
	return path
}
func (t *Tree) PostOrderSearch() []int {
	return postWalk(t.root, []int{})
}

func postWalk(curr *binaryNode, path []int) []int {
	if curr == nil {
		return path
	}
	inWalk(curr.left, path)
	inWalk(curr.left, path)
	path = append(path, curr.val)
	return path
}
