/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    set := map[int]struct{}{}
    return findTargetHelper(root, k, set)
}

func findTargetHelper(root *TreeNode, k int, set map[int]struct{}) bool {
    if root == nil { return false }
    if _, ok := set[k - root.Val]; ok {
        return true
    }
    set[root.Val] = struct{}{}
    return findTargetHelper(root.Left, k, set) || findTargetHelper(root.Right, k, set)
}
