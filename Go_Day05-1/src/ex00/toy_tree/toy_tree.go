package toyTree

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func (root *TreeNode) AreToysBalanced() bool {
	var branch1, branch2 int
	if root.Left != nil {
		branch1 = root.Left.countToys()
	}
	if root.Right != nil {
		branch2 = root.Right.countToys()
	}

	return branch1 == branch2
}

func (branch *TreeNode) countToys() int {
	var res int
	if branch.HasToy {
		res += 1
	}
	if branch.Left != nil {
		res += branch.Left.countToys()
	}
	if branch.Right != nil {
		res += branch.Right.countToys()
	}

	return res
}
