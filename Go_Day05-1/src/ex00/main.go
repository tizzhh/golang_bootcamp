package main

import (
	"fmt"
	"toyTree/toy_tree"
)

func main() {
	var root toyTree.TreeNode
	root.Left = &toyTree.TreeNode{}
	root.Left.Left = &toyTree.TreeNode{}
	root.Left.Right = &toyTree.TreeNode{HasToy: true}
	root.Right = &toyTree.TreeNode{HasToy: true}
	fmt.Println(root.AreToysBalanced())

	var root2 toyTree.TreeNode
	root2.Left = &toyTree.TreeNode{HasToy: true}
	root2.Left.Right = &toyTree.TreeNode{HasToy: true}
	root2.Right = &toyTree.TreeNode{}
	root2.Right.Right = &toyTree.TreeNode{HasToy: true}
	fmt.Println(root2.AreToysBalanced())
}
