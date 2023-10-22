/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
    return hasPathSumHelper(root, sum, 0)
}

func hasPathSumHelper(node *TreeNode, sum, total int) bool {
    if node == nil { return false }
    total += node.Val
    if node.Left == nil && node.Right == nil && total == sum { return true }
    return hasPathSumHelper(node.Left, sum, total) || hasPathSumHelper(node.Right, sum, total)
}
