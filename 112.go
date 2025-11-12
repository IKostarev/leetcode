package main

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	sumSlice, sum := []int{}, 0

	calculateSum(root, &sum, &sumSlice)

	for _, v := range sumSlice {
		if v == targetSum {
			return true
		}
	}

	return false
}

func calculateSum(root *TreeNode, sum *int, sumSlice *[]int) {
	if root == nil {
		return
	}

	*sum += root.Val

	if root.Left == nil && root.Right == nil {
		*sumSlice = append(*sumSlice, *sum)
		*sum -= root.Val
		return
	}

	calculateSum(root.Left, sum, sumSlice)
	calculateSum(root.Right, sum, sumSlice)
	*sum -= root.Val
}
