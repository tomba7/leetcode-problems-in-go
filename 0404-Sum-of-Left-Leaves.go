func sumOfLeftLeaves(root *TreeNode) int {
    return sumOfLeftLeavesHelper(root, false)
}

func sumOfLeftLeavesHelper(root *TreeNode, left bool) int {
    if root == nil {
        return 0
    } else if root.Left == nil && root.Right == nil && left {
        return root.Val
    }
    return sumOfLeftLeavesHelper(root.Left, true) + sumOfLeftLeavesHelper(root.Right, false)
}
