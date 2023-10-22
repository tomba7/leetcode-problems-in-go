func deleteAndEarn(nums []int) int {
    var n int
    for i := 0; i < len(nums); i++ {
        n = max(n, nums[i])
    }
    arr := make([]int, n+1)
    for _, num := range nums {
        arr[num]++
    }
    dp := make([]int, n+2)
    dp[n] = arr[n] * n
    for i := n-1; i >= 0; i-- {
        dp[i] = max(i*arr[i] + dp[i+2], dp[i+1])
    }
    return dp[0]
}

func max(x, y int) int { if x > y { return x }; return y }
