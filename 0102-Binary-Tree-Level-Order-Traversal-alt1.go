/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    result := [][]int{}
    levelOrderHelper(root, 0, &result)
    return result
}

func levelOrderHelper(node *TreeNode, level int, result *[][]int) {
    if node == nil { return }
    if level == len(*result) { *result = append(*result, []int{}) }
    (*result)[level] = append((*result)[level], node.Val)
    levelOrderHelper(node.Left, level + 1, result)
    levelOrderHelper(node.Right, level + 1, result)
}
