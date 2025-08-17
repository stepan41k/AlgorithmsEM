package main

// 102. Binary Tree Level Order Traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traversal(node *TreeNode, level int, result *[][]int) {
	if node == nil {
		return
	}

	if len(*result) < level+1 {
		*result = append(*result, []int{node.Val})
	} else {
		(*result)[level] = append((*result)[level], node.Val)
	}

	
	traversal(node.Left, level+1, result)
	traversal(node.Right, level+1, result)
}

func levelOrder(root *TreeNode) [][]int {
    level, result := 0, [][]int{}
    traversal(root, level, &result)
    return result
}