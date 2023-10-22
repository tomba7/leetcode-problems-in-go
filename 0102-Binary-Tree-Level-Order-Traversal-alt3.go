/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil { return nil }
    var result [][]int
    q := list.New()
    q.PushBack(root)
    for q.Len() != 0 {
        numLevelNodes := q.Len()
        var levelNodes []int
        for i := 0; i < numLevelNodes; i++ {
            curr := q.Front().Value.(*TreeNode)
            q.Remove(q.Front())
            levelNodes = append(levelNodes, curr.Val)
            if curr.Left != nil { q.PushBack(curr.Left) }
            if curr.Right != nil { q.PushBack(curr.Right) }
        }
        result = append(result, levelNodes)
    }
    return result
}
