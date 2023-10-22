func isSubtree(s, t *TreeNode) bool {
    seq1, seq2 := traversalSeq(s), traversalSeq(t)
    return bytes.Contains(seq1, seq2)
}

func traversalSeq(root *TreeNode) (result []byte) {
    traversalSeqHelper(root, &result)
    return result
}

func traversalSeqHelper(node *TreeNode, result *[]byte) {
    if node == nil {
        *result = append(*result, ",nil"...)
        return
    }
    *result = append(*result, fmt.Sprintf(",%d", node.Val)...)
    traversalSeqHelper(node.Left, result)
    traversalSeqHelper(node.Right, result)
}

---

func isSubtree(s, t *TreeNode) bool {
    if t == nil { return false }
    return isSubtreeHelper(s, t)
}

func isSubtreeHelper(s, t *TreeNode) bool {
    if s == nil {
        return false
    } else if s.Val == t.Val && isIdentical(s, t) {
        return true
    }
    return isSubtreeHelper(s.Left, t) || isSubtreeHelper(s.Right, t)
}

func isIdentical(s, t *TreeNode) bool {
    if s == nil && t == nil {
        return true
    } else if s == nil || t == nil || s.Val != t.Val {
        return false
    }
    return isIdentical(s.Left, t.Left) && isIdentical(s.Right, t.Right)
}
