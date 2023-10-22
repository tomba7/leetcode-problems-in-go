/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    var result int
    maxDepthHelper(root, 1, &result) 
    return result
}

func maxDepthHelper(root *TreeNode, depth int, result *int) {
    if root == nil {
        return
    } else if root.Left == nil && root.Right == nil {
        if depth > *result { *result = depth }
        return
    }
    maxDepthHelper(root.Left, depth+1, result) 
    maxDepthHelper(root.Right, depth+1, result) 
}
