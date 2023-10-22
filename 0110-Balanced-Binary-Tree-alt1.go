/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) (balanced bool) {
    _, balanced = isBalancedHelper(root)
    return
}

func isBalancedHelper(node *TreeNode) (int, bool) {
    if node == nil{ return 0, true }
    leftHeight, leftBalanced := isBalancedHelper(node.Left)
    if !leftBalanced { return 0, false }

    rightHeight, rightBalanced := isBalancedHelper(node.Right)
    if !rightBalanced { return 0, false }

    if math.Abs(float64(leftHeight - rightHeight)) > 1 {
        return 0, false
    }
    return max(leftHeight, rightHeight) + 1, true
}

func max(x, y int) int { if x > y { return x }; return y }
