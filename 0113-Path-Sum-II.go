/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
    result := [][]int{}
    pathSumHelper(root, sum, 0, nil, &result)
    return result
}

func pathSumHelper(node *TreeNode, sum, total int, path []int, result *[][]int) {
    if node == nil { return }
    total += node.Val
    path = append(path, node.Val)
    if node.Left == nil && node.Right == nil && sum == total {
        *result = append(*result, append([]int{}, path...))
        return
    }
    pathSumHelper(node.Left, sum, total, path, result)
    pathSumHelper(node.Right, sum, total, path, result)
}
