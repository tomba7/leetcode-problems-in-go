/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) != len(postorder) { return nil }
    indx := len(postorder)-1
    table := map[int]int{}
    for i, n := range inorder { table[n] = i }
    return buildTreeHelper(postorder, &indx, table, 0, len(postorder)-1)
}

func buildTreeHelper(postorder []int, indx *int, table map[int]int, lo, hi int) *TreeNode {
    if lo > hi { return nil }
    val := postorder[*indx]
    *indx--
    mid := table[val] 
    node := &TreeNode{Val: val}
    node.Right = buildTreeHelper(postorder, indx, table, mid+1, hi)
    node.Left = buildTreeHelper(postorder, indx, table, lo, mid-1)
    return node
}
