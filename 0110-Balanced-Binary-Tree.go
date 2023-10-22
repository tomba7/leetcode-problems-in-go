/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return isBalancedHelper(root) >= 0
}

func isBalancedHelper(node *TreeNode) int {
    if (node == nil) { return 0 }

    left := isBalancedHelper(node.Left)
    if left < 0 { return left }

    right := isBalancedHelper(node.Right)
    if right < 0 { return right }

    if math.Abs(float64(left -right)) > 1 { return -1 }
    return max(left, right) + 1
}

func max(x, y int) int { if x > y { return x }; return y }
