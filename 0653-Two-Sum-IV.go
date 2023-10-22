/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    nums := []int{}
    traverse(root, &nums)
    for i, j := 0, len(nums) - 1; i < j; {
        sum := nums[i] + nums[j]
        if sum == k {
            return true
        } else if sum < k {
            i++
        } else {
            j--
        }
    }
    return false
}

func traverse(root *TreeNode, nums *[]int) {
    if root == nil { return }
    traverse(root.Left, nums)
    *nums = append(*nums, root.Val)
    traverse(root.Right, nums)
}
