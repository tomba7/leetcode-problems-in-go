/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    var prev *TreeNode
    flattenHelper(root, &prev)
}

func flattenHelper(root *TreeNode, prev **TreeNode)  {
    if root == nil { return }
    flattenHelper(root.Right, prev)
    flattenHelper(root.Left, prev)
    root.Right = *prev
    root.Left = nil
    *prev = root
}
