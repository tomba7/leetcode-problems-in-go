func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    var subtree *TreeNode
    var maxDepth int
    subtreeHelper(root, 0, &maxDepth, &subtree)
    return subtree
}

func subtreeHelper(root *TreeNode, i int, maxDepth *int, subtree **TreeNode) int {
    if root == nil {
        return i
    }
    *maxDepth = max(*maxDepth, i)
    left := subtreeHelper(root.Left, i + 1, maxDepth, subtree)
    right := subtreeHelper(root.Right, i + 1, maxDepth, subtree)
    if left == right && left > *maxDepth {
        *subtree = root
    }
    return max(left, right)
}

func max(x, y int) int { if x > y { return x }; return y }
