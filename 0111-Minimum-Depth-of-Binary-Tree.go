func minDepth(root *TreeNode) int {
    if root == nil { return 0 }
    if root.Left == nil { return minDepth(root.Right) + 1 }
    if root.Right == nil { return minDepth(root.Left) + 1 }
    return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(x, y int) int { if x < y { return x }; return y }

---

func minDepth(root *TreeNode) int {
    if root == nil { return 0 }
    return minDepthHelper(root)
}

func minDepthHelper(root *TreeNode) int {
    if root == nil { return math.MaxInt32 }
    if root.Left == nil && root.Right == nil { return 1 }
    return min(minDepthHelper(root.Left), minDepthHelper(root.Right)) + 1
}

func min(x, y int) int { if x < y { return x }; return y }
