func bstFromPreorder(preorder []int) *TreeNode {
    return bstFromPreorderHelper(preorder, 0, len(preorder)-1)
}

func bstFromPreorderHelper(preorder []int, lo, hi int) *TreeNode {
    if lo > hi {
        return nil
    }
    var mid int
    for mid = lo+1; mid <= hi && preorder[mid] < preorder[lo]; mid++ {}
    return &TreeNode{
        Val: preorder[lo],
        Left: bstFromPreorderHelper(preorder, lo+1, mid-1),
        Right: bstFromPreorderHelper(preorder, mid, hi),
    }
}

---

// Longer logical version
func bstFromPreorder(preorder []int) *TreeNode {
    return bstFromPreorderHelper(preorder, 0, len(preorder)-1)
}

func bstFromPreorderHelper(preorder []int, lo, hi int) *TreeNode {
    if lo > hi {
        return nil
    } else if lo == hi {
        return &TreeNode{Val: preorder[lo]}
    }
    var mid int
    if preorder[lo] < preorder[lo+1] {
        mid = lo + 1
    } else {
        for mid = lo+1; mid <= hi && preorder[mid] < preorder[lo]; mid++ {}
    }
    return &TreeNode{
        Val: preorder[lo],
        Left: bstFromPreorderHelper(preorder, lo+1, mid-1),
        Right: bstFromPreorderHelper(preorder, mid, hi),
    }
}
