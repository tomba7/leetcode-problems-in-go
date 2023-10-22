/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    depth := getDepth(root)
    result := make([][]int, depth) 
    levelOrderBottomHelper(root, 0, depth, &result)
    return result
}

func levelOrderBottomHelper(node *TreeNode, level, depth int, result *[][]int) {
    if node == nil { return }
    levelOrderBottomHelper(node.Left, level + 1, depth, result)
    levelOrderBottomHelper(node.Right, level + 1, depth, result)
    (*result)[depth - level - 1] = append((*result)[depth - level - 1], node.Val)
}

func getDepth(root *TreeNode) int {
    if root == nil { return 0 }
    return max(getDepth(root.Left), getDepth(root.Right)) + 1
}

func max(x, y int) int { if x > y { return x }; return y }
