package toyTree

import (
	"toyTree/queue"
)

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func reverseSlice(toys []bool) {
	for i, j := 0, len(toys)-1; i < j; i, j = i+1, j-1 {
		toys[i], toys[j] = toys[j], toys[i]
	}
}

func (root *TreeNode) UnrollGarland() []bool {
	var toys []bool
	var que queue.Queue
	layer := 2

	toys = append(toys, root.HasToy)
	que.Enque(root)

	for len(que.Queue) != 0 {
		var nodes []*TreeNode
		for len(que.Queue) != 0 {
			elem := que.Deque().(*TreeNode)
			if elem.Left != nil {
				nodes = append(nodes, elem.Left)
			}
			if elem.Right != nil {
				nodes = append(nodes, elem.Right)
			}
		}
		var nodeToys []bool
		for _, val := range nodes {
			nodeToys = append(nodeToys, val.HasToy)
		}
		if layer%2 != 0 {
			reverseSlice(nodeToys)
		}
		toys = append(toys, nodeToys...)
		for _, node := range nodes {
			que.Enque(node)
		}
		layer += 1
	}

	return toys
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
