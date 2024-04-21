package toyTree_test

import (
	"testing"
	"toyTree/toy_tree"
)

type toysBalancedTest struct {
	root     toyTree.TreeNode
	expected bool
}

var treesTestsBalanced = []toysBalancedTest{
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
	for _, test := range treesTestsBalanced {
		if areBalanced := test.root.AreToysBalanced(); areBalanced != test.expected {
			t.Errorf("Output %t not equal to expected %t", areBalanced, test.expected)
		}
	}
}

type toysUnrollTest struct {
	root     toyTree.TreeNode
	expected []bool
}

var treeTestsUnroll = []toysUnrollTest{
	{
		expected: []bool{false},
		root:     toyTree.TreeNode{},
	},
	{
		expected: []bool{true, true, false, true, true, false, true},
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
		expected: []bool{true, true, false, true, true, false, true, true, false, false, true, true, true},
		root: toyTree.TreeNode{
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
		},
	},
}

func compareBoolSlices(s1, s2 []bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, val := range s1 {
		if val != s2[i] {
			return false
		}
	}

	return true
}

func TestUnroll(t *testing.T) {
	for _, test := range treeTestsUnroll {
		if unrolled := test.root.UnrollGarland(); !compareBoolSlices(unrolled, test.expected) {
			t.Errorf("Output \n%v not equal to expected \n%v", unrolled, test.expected)
		}
	}
}
