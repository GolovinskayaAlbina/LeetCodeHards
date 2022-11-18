package medium

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lLevel := getLevel(root, func(node *TreeNode) *TreeNode { return node.Left })
	rLevel := getLevel(root, func(node *TreeNode) *TreeNode { return node.Right })
	if lLevel == rLevel {
		return int(math.Pow(2, float64(lLevel))) - 1
	}
	return 1 + countNodes(root.Right) + countNodes(root.Left)
}

func getLevel(node *TreeNode, move func(node *TreeNode) *TreeNode) int {
	level := 0
	for node != nil {
		level++
		node = move(node)
	}
	return level
}
