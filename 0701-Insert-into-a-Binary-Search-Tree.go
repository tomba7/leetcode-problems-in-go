func insertIntoBST(root *TreeNode, val int) *TreeNode {
    node := &TreeNode{Val: val}
    if root == nil { return node }
    var prev, curr *TreeNode = nil, root
    for curr != nil {
        prev = curr
        if val < curr.Val {
            curr = curr.Left
        } else {
            curr = curr.Right
        }
    }
    if val < prev.Val {
        prev.Left = node
    } else {
        prev.Right = node
    }
    return root
}

---

func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
         return &TreeNode{Val: val}
    }
    if val < root.Val {
        root.Left = insertIntoBST(root.Left, val)
    } else {
        root.Right = insertIntoBST(root.Right, val)
    }
    return root
}
