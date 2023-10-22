func rob(nums []int) int {
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

// Memoization
func rob(nums []int) int {
    cache := make([]int, len(nums))
    for i := range cache {
        cache[i] = -1
    }
    return robHelper(nums, 0, cache)
}

func robHelper(nums []int, i int, cache []int) int {
    if i >= len(nums) {
        return 0
    }
    if cache[i] == -1 {
        cache[i] = max(
            nums[i] + robHelper(nums, i+2, cache), robHelper(nums, i+1, cache),
        )
    }
    return cache[i]
}

func max(x, y int) int { if x > y { return x }; return y }

---

// Recursion (TLE)
func rob(nums []int) int {
    return robHelper(nums, 0)
}

func robHelper(nums []int, i int) int {
    if i >= len(nums) {
        return 0
    }
    return max(
        nums[i] + robHelper(nums, i+2), robHelper(nums, i+1),
    )
}

func max(x, y int) int { if x > y { return x }; return y }
