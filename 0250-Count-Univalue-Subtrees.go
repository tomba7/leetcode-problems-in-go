func countUnivalSubtrees(root *TreeNode) int {
    _, count := helper(root)
    return count
}

func helper(root *TreeNode) (bool, int) {
    if root == nil {
        return true, 0
    }
    left, leftCount := helper(root.Left)
    right, rightCount := helper(root.Right)
    unival := left && right &&
                (root.Left == nil || root.Left.Val == root.Val) &&
                (root.Right == nil || root.Right.Val == root.Val)
    count := leftCount + rightCount
    if unival { count++ }
    return unival, count
}

---

func countUnivalSubtrees(root *TreeNode) int {
    _, count := helper(root)
    return count
}

func helper(root *TreeNode) (bool, int) {
    if root == nil {
        return true, 0
    }
    left, leftCount := helper(root.Left)
    right, rightCount := helper(root.Right)
    if left && right &&
        (root.Left == nil || root.Left.Val == root.Val) &&
        (root.Right == nil || root.Right.Val == root.Val) {
        return true, leftCount+rightCount+1
    }
    return false, leftCount+rightCount
}

---

func countUnivalSubtrees(root *TreeNode) int {
    var total int
    countHelper(root, &total)
    return total
}

func countHelper(root *TreeNode, total *int) bool {
    if root == nil { return true }
    left := countHelper(root.Left, total)
    right := countHelper(root.Right, total)
    unival := left && right &&
                (root.Left == nil || root.Left.Val != root.Val) &&
                (root.Right == nil || root.Right.Val != root.Val)
    if unival { *total++ }
    return unival
}

---

func countUnivalSubtrees(root *TreeNode) int {
    var result int
    helper(root, &result)
    return result
}

func helper(root *TreeNode, result *int) bool {
    if root == nil { return true }
    left, right := helper(root.Left, result), helper(root.Right, result)
    if left && right {
        if root.Left != nil && root.Left.Val != root.Val {
            return false
        } else if root.Right != nil && root.Right.Val != root.Val {
            return false
        }
        *result++
        return true
    }
    return false
}

---

func countUnivalSubtrees(root *TreeNode) int {
    var total int
    count(root, &total)
    return total
}

func count(root *TreeNode, total *int) bool {
    if root == nil { return true }
    if root.Left == nil && root.Right == nil {
        *total++
        return true
    } else if root.Right == nil {
        if count(root.Left, total) && root.Left.Val == root.Val {
            *total++
            return true
        }
    } else if root.Left == nil {
        if count(root.Right, total) && root.Right.Val == root.Val {
            *total++
            return true
        }
    } else {
        left, right := count(root.Left, total), count(root.Right, total)
        if left == right && root.Left.Val == root.Right.Val && root.Left.Val == root.Val {
            *total++
            return true
        }
    }
    return false
}
