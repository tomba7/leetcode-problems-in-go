/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    freq := map[int]int{0: 1}
    var numPaths int
    pathSumHelper(root, targetSum, 0, freq, &numPaths)
    return numPaths
}

func pathSumHelper(node *TreeNode, targetSum, runningSum int, freq map[int]int, numPaths *int) {
    if node == nil { return }
    runningSum += node.Val
    if count, ok := freq[runningSum - targetSum]; ok { *numPaths += count }
    freq[runningSum]++
    pathSumHelper(node.Left, targetSum, runningSum, freq, numPaths)
    pathSumHelper(node.Right, targetSum, runningSum, freq, numPaths)
    freq[runningSum]--
}
