package main

// https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/928/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	ans := []int{}

	ans = helper(root, ans)

	return ans
}

func helper(root *TreeNode, result []int) []int {
	if root == nil {
		return make([]int, 0)
	}
	ans := []int{}
	ans = append(ans, root.Val)

	ans1 := helper(root.Left, ans)
	ans = append(ans, ans1...)

	ans2 := helper(root.Right, ans)
	ans = append(ans, ans2...)

	return ans
}
