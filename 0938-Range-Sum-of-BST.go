/*
 - Use binary search to find the 'low' node.
 - From 'low', start an in
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    var sum int
    helper(root, low, high, &sum)
    return sum
}

func helper(root *TreeNode, low, high int, sum *int) {
    if root == nil { return }
    if root.Val >= low && root.Val <= high {
        *sum += root.Val
    }
    if root.Val > low {
        helper(root.Left, low, high, sum)
    }
    if root.Val < high {
        helper(root.Right, low, high, sum)
    }
}

---

func rangeSumBST(root *TreeNode, low int, high int) int {
    if root == nil {
        return 0
    }
    var res int
    if low <= root.Val && root.Val <= high {
        res = root.Val
    }
    if low < root.Val {
        res += rangeSumBST(root.Left, low, high)
    }
    if high > root.Val {
        res += rangeSumBST(root.Right, low, high)
    }
    return res
}

---

func rangeSumBST(root *TreeNode, low int, high int) int {
    if root == nil {
        return 0
    }
    val := rangeSumBST(root.Left, low, high)
    if low <= root.Val && root.Val <= high {
        val += root.Val
    }
    return val + rangeSumBST(root.Right, low, high)
}
