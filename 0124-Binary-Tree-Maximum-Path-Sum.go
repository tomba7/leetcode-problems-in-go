/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    res := math.MinInt32
    maxPathSumHelper(root, &res)
    return res
}

func maxPathSumHelper(root *TreeNode, res *int) int {
    if root == nil { return 0 }
    left := maxPathSumHelper(root.Left, res)
    right := maxPathSumHelper(root.Right, res)
    maxPathEndingHere := max(root.Val, root.Val + max(left, right))
    *res = max3(*res, maxPathEndingHere, root.Val + left + right)
    return maxPath
}

func max(x, y int) int { if x > y { return x }; return y }
func max3(x, y, z int) int { return max(x, max(y, z)) }

---

func maxPathSum(root *TreeNode) int {
    maxSum := math.MinInt32
    maxPathSumHelper(root, &maxSum)
    return maxSum
}

func maxPathSumHelper(node *TreeNode, maxSum *int) int {
    if node == nil { return 0 }
    left := maxPathSumHelper(node.Left, maxSum)
    right := maxPathSumHelper(node.Right, maxSum)
    maxSumEndingAtNode := max(node.Val, node.Val + max(left, right))
    *maxSum = max(*maxSum, max(node.Val + left + right, maxSumEndingAtNode))
    return maxSumEndingAtNode
}

func max(x, y int) int { if x > y { return x }; return y }
