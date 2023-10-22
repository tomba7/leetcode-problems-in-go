/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    return sumNumbersHelper(root, 0)
}

func sumNumbersHelper(root *TreeNode, value int) int {
    if root == nil {
        return 0
    }
    value = value * 10 + root.Val
    if root.Left == nil && root.Right == nil {
        return value
    }
    return sumNumbersHelper(root.Left, value) +
           sumNumbersHelper(root.Right, value)
}
