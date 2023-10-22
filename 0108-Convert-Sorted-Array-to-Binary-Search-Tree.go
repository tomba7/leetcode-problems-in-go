/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    return sortedArrayToBSTHelper(nums, 0, len(nums) - 1)
}

func sortedArrayToBSTHelper(nums []int, low, high int) *TreeNode {
    if low > high { return nil }
    mid := low + (high - low)/2
    leftNode := sortedArrayToBSTHelper(nums, low, mid - 1)
    rightNode := sortedArrayToBSTHelper(nums, mid + 1, high)
    return &TreeNode{nums[mid], leftNode, rightNode}
}

---

func sortedArrayToBST(nums []int) *TreeNode {
    return bstHelper(nums, 0, len(nums)-1)
}

func bstHelper(nums []int, lo, hi int) *TreeNode {
    if lo > hi { return nil }
    mid := lo + (hi-lo)/2
    return &TreeNode{
        Val: nums[mid],
        Left: bstHelper(nums, lo, mid-1),
        Right: bstHelper(nums, mid+1, hi),
    }
}
