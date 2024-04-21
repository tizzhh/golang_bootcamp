package toyTree_test

import (
	"testing"
	"toyTree/toy_tree"
)

type toysBalancedTest struct {
	root     toyTree.TreeNode
	expected bool
}

var treesTests = []toysBalancedTest{
	{
		expected: true,
		root: toyTree.TreeNode{
			Left: &toyTree.TreeNode{
				Left: &toyTree.TreeNode{},
				Right: &toyTree.TreeNode{
					HasToy: true,
				},
			},
			Right: &toyTree.TreeNode{
				HasToy: true,
			},
		},
	},
	{
		expected: true,
		root: toyTree.TreeNode{
			HasToy: true,
			Left: &toyTree.TreeNode{
				HasToy: true,
				Left: &toyTree.TreeNode{
					HasToy: true,
				},
				Right: &toyTree.TreeNode{},
			},
			Right: &toyTree.TreeNode{
				Left: &toyTree.TreeNode{
					HasToy: true,
				},
				Right: &toyTree.TreeNode{
					HasToy: true,
				},
			},
		},
	},
	{
		expected: false,
		root: toyTree.TreeNode{
			HasToy: true,
			Left: &toyTree.TreeNode{
				HasToy: true,
			},
			Right: &toyTree.TreeNode{
				HasToy: false,
			},
		},
	},
	{
		expected: false,
		root: toyTree.TreeNode{
			Left: &toyTree.TreeNode{
				HasToy: true,
				Right: &toyTree.TreeNode{
					HasToy: true,
				},
			},
			Right: &toyTree.TreeNode{
				Right: &toyTree.TreeNode{
					HasToy: true,
				},
			},
		},
	},
}

func TestAreToysBalanced(t *testing.T) {
	for _, test := range treesTests {
		if areBalanced := test.root.AreToysBalanced(); areBalanced != test.expected {
			t.Errorf("Output %t not equal to expected %t", areBalanced, test.expected)
		}
	}
}
