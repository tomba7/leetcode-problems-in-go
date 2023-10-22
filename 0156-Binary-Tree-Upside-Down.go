/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
    return upsideDownBinaryTreeHelper(root, nil)
}

func upsideDownBinaryTreeHelper(node, parent *TreeNode) *TreeNode {
    if node == nil { return parent }
    newRoot := upsideDownBinaryTreeHelper(node.Left, node)
    if parent == nil {
        node.Left = nil
    } else {
        node.Left = parent.Right
    }
    node.Right = parent
    return newRoot
}
