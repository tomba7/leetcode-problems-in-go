func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    set := map[int]bool{}
    for i := 0; i < len(to_delete); i++ {
        set[to_delete[i]] = true
    }
    var res []*TreeNode
    if !set[root.Val] {
        res = append(res, root)
    }
    delNodesHelper(root, set, &res)
    return res
}

func delNodesHelper(root *TreeNode, set map[int]bool, res *[]*TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    root.Left = delNodesHelper(root.Left, set, res)
    root.Right = delNodesHelper(root.Right, set, res)
    if set[root.Val] {
        if root.Left != nil {
            *res = append(*res, root.Left)
        }
        if root.Right != nil {
            *res = append(*res, root.Right)
        }
        return nil
    }
    return root
}
