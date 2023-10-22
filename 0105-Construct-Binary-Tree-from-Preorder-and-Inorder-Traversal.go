/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(inorder) != len(preorder) { return nil }
    var indx int
    table := map[int]int{}
    for i, n := range inorder { table[n] = i }
    return buildTreeHelper(preorder, &indx, table, 0, len(preorder)-1)
}

func buildTreeHelper(preorder []int, indx *int, table map[int]int, lo, hi int) *TreeNode {
    if lo > hi { return nil }
    val := preorder[*indx]
    *indx++
    mid := table[val]
    return &TreeNode{
        Val: val,
        Left: buildTreeHelper(preorder, indx, table, lo, mid-1),
        Right: buildTreeHelper(preorder, indx, table, mid+1, hi),
    }
}
