/*
 Time: O(n)
 Space: O(n) + O(log n) = O(n)
 */
func balanceBST(root *TreeNode) *TreeNode {
    var sortedNodes []*TreeNode
    inorderBST(root, &sortedNodes)
    return buildBST(sortedNodes, 0, len(sortedNodes)-1)
}

func inorderBST(root *TreeNode, sortedNodes *[]*TreeNode) {
    if root == nil { return }
    inorderBST(root.Left, sortedNodes)
    *sortedNodes = append(*sortedNodes, root)
    inorderBST(root.Right, sortedNodes)
}

func buildBST(sortedNodes []*TreeNode, lo, hi int) *TreeNode {
    if lo > hi {
        return nil
    }
    mid := lo + (hi-lo)/2
    sortedNodes[mid].Left = buildBST(sortedNodes, lo, mid-1)
    sortedNodes[mid].Right = buildBST(sortedNodes, mid+1, hi)
    return sortedNodes[mid]
}
