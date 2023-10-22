/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
    x, y := robHelper(root)
    return max(x, y)
}

func robHelper(node *TreeNode) (int, int) {
    if node == nil { return 0, 0 }
    leftX, leftY := robHelper(node.Left)
    rightX, rightY := robHelper(node.Right)
    return node.Val + leftY + rightY, max(leftX, leftY) + max(rightX, rightY)
}

func max(x, y int) int { if x > y { return x }; return y }
