func searchBST(root *TreeNode, val int) *TreeNode {
    for root != nil {
        if val < root.Val {
            root = root.Left
        } else if val > root.Val {
            root = root.Right
        } else {
            return root
        }
    }
    return nil
}
