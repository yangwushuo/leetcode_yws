/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {

	/*
		解题思路:
			使用广度优先 一层一层遍历节点
			1创建一个用于存放节点的列表 压入根节点进入列表中
			2遍历列表，如果列表中含有节点则进行循环，判断列表中的子节点是否还有后代，如果有则继续压入列表
	*/

	if root == nil {
		return 0
	}

	queen := []*TreeNode{}
	queen = append(queen, root)
	anns := 0

	for len(queen) > 0 {
		size := len(queen)
		for size > 0 {
			node := queen[0]
			queen = queen[1:]
			if node.Left != nil {
				queen = append(queen, node.Left)
			}
			if node.Right != nil {
				queen = append(queen, node.Right)
			}
			size--
		}
		anns++
	}

	return anns
}

// @lc code=end

