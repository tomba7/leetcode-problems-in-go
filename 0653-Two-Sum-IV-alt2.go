/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    return findTargetHelper(root, root, k)
}

func findTargetHelper(root, node *TreeNode, k int) bool {
    if node == nil { return false }
    if search(root, node, k - node.Val) { return true }
    return findTargetHelper(root, node.Left, k) || findTargetHelper(root, node.Right, k)
}

func search(root, node *TreeNode, target int) bool {
    if root == nil { return false }
    if root.Val == target && root != node { return true }
    return search(root.Left, node, target) || search(root.Right, node, target)
}
