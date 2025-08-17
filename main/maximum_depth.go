package main

// 104. Maximum Depth of Binary Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func depth(node *TreeNode, number int) int {
    if node == nil {
        return 0
    }

    left := depth(node.Left, number+1)
    right := depth(node.Right, number+1)

    return max(left, right) + 1
}

func maxDepth(root *TreeNode) int {
    ans := depth(root, 0)
    return ans
}