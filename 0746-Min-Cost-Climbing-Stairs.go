func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    dp := make([]int, n+1)
    for i := 2; i <= n; i++ {
        dp[i] = min(
            cost[i-1] + dp[i-1],
            cost[i-2] + dp[i-2],
        )
    }
    return dp[n]
}

func min(x, y int) int { if x < y { return x }; return y }

---

// Recursion (TLE)
func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    return helper(cost, n)
}

func helper(cost []int, i int) int {
    if i <= 1 { return 0 }
    return min(
        cost[i-1] + helper(cost, i-1),
        cost[i-2] + helper(cost, i-2),
    )
}

func min(x, y int) int { if x < y { return x }; return y }
