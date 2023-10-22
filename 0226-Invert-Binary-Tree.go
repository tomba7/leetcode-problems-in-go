/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 *
 * Invert a binary tree.
 * 
 *       4
 *     /   \
 *    2     7
 *   / \   / \
 *  1   3 6   9
 *  
 *  to
 *       4
 *     /   \
 *    7     2
 *   / \   / \
 *  9   6 3   1
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil { return nil }
    left := invertTree(root.Left)
    right := invertTree(root.Right)
    root.Left, root.Right = right, left
    return root
}
