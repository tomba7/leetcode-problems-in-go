func maximumProduct(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    return max(nums[0] * nums[1] * nums[n-1], nums[n-3] * nums[n-2] * nums[n-1])
}

func max(x, y int) int { if x > y { return x }; return y }
