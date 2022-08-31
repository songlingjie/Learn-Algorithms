package main

import (
	"fmt"
)

type tree struct {
	data int
	left *tree
	right *tree
}

func main()  {
	arr := []int{1,3,5,6,9,10}

	var node *tree
	fmt.Printf("node is %v\n", node)

	for i := 0; i < len(arr); i++ {
		node = insert(node, arr[i])
		fmt.Printf("node is %v\n", node)
	}
	preOrder(node)
}

func insert(node *tree, nodeVal int, ) *tree {

	if node == nil {
		return &tree{data: nodeVal}
	}

	if node.data == nodeVal {
		return node
	}
	if node.data > nodeVal {
		node.left = insert(node.left, nodeVal)
		return node
	}

	node.right = insert(node.right, nodeVal)
	return  node
}

// 前序遍历
func preOrder(node *tree)  {
	if node == nil {
		return
	}

	fmt.Println(node.data, " ")
	preOrder(node.left)
	preOrder(node.right)
}

// 中序遍历
func inOrder(node *tree)  {
	if node == nil {
		return
	}

	preOrder(node.left)
	fmt.Println(node.data, " ")
	preOrder(node.right)
}

// 后序遍历
func postOrder(node *tree)  {
	if node == nil {
		return
	}

	preOrder(node.left)
	preOrder(node.right)
	fmt.Println(node.data, " ")
}



