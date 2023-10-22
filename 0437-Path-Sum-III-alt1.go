/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// This is a brute force approach 
func pathSum(root *TreeNode, sum int) (count int) {
    traverse(root, sum, &count)
    return count
}

func traverse(node *TreeNode, sum int, count *int) {
    if node == nil { return }
    pathSumHelper(node, sum, 0, count)
    traverse(node.Left, sum, count)
    traverse(node.Right, sum, count)
}

func pathSumHelper(node *TreeNode, sum, total int, count *int) {
    if node == nil { return }
    if total + node.Val == sum { *count++ }
    pathSumHelper(node.Left, sum, total + node.Val, count)
    pathSumHelper(node.Right, sum, total + node.Val, count)
}
