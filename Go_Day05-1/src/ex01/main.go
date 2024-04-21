package main

import (
	"fmt"
	"toyTree/toy_tree"
)

func main() {
	root := toyTree.TreeNode{
		HasToy: true,
		Left: &toyTree.TreeNode{
			HasToy: true,
			Left: &toyTree.TreeNode{
				HasToy: true,
				Left: &toyTree.TreeNode{
					HasToy: true,
				},
				Right: &toyTree.TreeNode{},
			},
			Right: &toyTree.TreeNode{
				Left: &toyTree.TreeNode{},
				Right: &toyTree.TreeNode{
					HasToy: true,
				},
			},
		},
		Right: &toyTree.TreeNode{
			Left: &toyTree.TreeNode{
				HasToy: true,
			},
			Right: &toyTree.TreeNode{
				HasToy: true,
				Left: &toyTree.TreeNode{
					HasToy: true,
				},
				Right: &toyTree.TreeNode{
					HasToy: true,
				},
			},
		},
	}

	fmt.Println(root.UnrollGarland())
}
