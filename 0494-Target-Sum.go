// Brute force approach
func findTargetSumWays(nums []int, target int) int {
    var res int
    helper(nums, 0, 0, target, &res)
    return res
}

func helper(nums []int, i, sum, target int, res *int) {
    if i == len(nums) {
        if sum == target { *res++ }
        return
    }
    helper(nums, i+1, sum+nums[i], target, res)
    helper(nums, i+1, sum-nums[i], target, res)
}
