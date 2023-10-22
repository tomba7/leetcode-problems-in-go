func rob(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    } else if n == 1 {
        return nums[0]
    }
    return max(robHelper(nums[1:]), robHelper(nums[:n-1]))
}

func robHelper(nums []int) int {
    n := len(nums)
    dp := make([]int, n+1)
    dp[n-1] = nums[n-1]
    for i := n-2; i >= 0; i-- {
        dp[i] = max(nums[i] + dp[i+2], dp[i+1])
    }
    return dp[0]
}

func max(x, y int) int { if x > y { return x }; return y }

---

func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    } else if len(nums) == 1 {
        return nums[0]
    }
    return max(robRangeOfHouses(nums, 0, len(nums) - 2), robRangeOfHouses(nums, 1, len(nums) - 1))
}

func robRangeOfHouses(nums []int, start, end int) int {
    var x, y int
    for i := start; i <= end; i++ {
        y, x = x, max(x, nums[i] + y)
    }
    return x
}

func max(x, y int) int { if x > y { return x }; return y }
