/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return isValidBSTHelper(root, int(math.MinInt64), int(math.MaxInt64))
}

func isValidBSTHelper(node *TreeNode, lo, hi int) bool {
    if node == nil { return true }
    return lo < node.Val && node.Val < hi &&
            isValidBSTHelper(node.Left, lo, node.Val) &&
            isValidBSTHelper(node.Right, node.Val, hi)
}

---

func isValidBST(root *TreeNode) bool {
    return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lo, hi int) bool {
    if root == nil { return true }
    if root.Val <= lo || root.Val >= hi { return false }
    return helper(root.Left, lo, root.Val) && helper(root.Right, root.Val, hi)
}

---

func isValidBST(root *TreeNode) bool {
    valid, _, _ := isValidBSTHelper(root)
    return valid
}

func isValidBSTHelper(root *TreeNode) (bool, int, int) {
    if root == nil {
        return true, math.MaxInt64, math.MinInt64
    }
    left, leftMin, leftMax := isValidBSTHelper(root.Left)
    if !left {
        return left, -1, -1
    }
    right, rightMin, rightMax := isValidBSTHelper(root.Right)
    if !right {
        return right, -1, -1
    }
    return leftMax < root.Val && root.Val < rightMin,
            min(leftMin, root.Val), max(root.Val, rightMax)
}

func min(x, y int) int { if x < y { return x }; return y }
func max(x, y int) int { if x > y { return x }; return y }

---

func isValidBST(root *TreeNode) bool {
    var prev *TreeNode = nil
    return isValidBSTHelper(root, &prev)
}

func isValidBSTHelper(node *TreeNode, prev **TreeNode) bool {
    if node == nil {return true}
    if !isValidBSTHelper(node.Left, prev) {return false}
    if *prev != nil && (*prev).Val >= node.Val {return false}
    *prev = node
    return isValidBSTHelper(node.Right, prev)
}
