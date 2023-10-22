func canPartition(nums []int) bool {
    var totalSum int
    for _, num := range nums { totalSum += num }
    if totalSum & 1 == 1 { return false }
    totalSum /= 2
    dp := make([]bool, totalSum+1)
    dp[0] = true
    for i := 0; i < len(nums); i++ {
        for sum := totalSum; sum > 0; sum-- {
            if nums[i] > sum { continue }
            dp[sum] = dp[sum] || dp[sum-nums[i]]
        }
    }
    return dp[totalSum]
}

---

func canPartition(nums []int) bool {
    var totalSum int
    for _, num := range nums { totalSum += num }
    if totalSum % 2 != 0 {
        return false
    }
    totalSum /= 2
    n := len(nums)
    dp := make([][]bool, n+1)
    for i := range dp {
        dp[i] = make([]bool, totalSum+1)
    }
    for i := 0; i <= n; i++ {
        dp[i][0] = true
    }
    for sum := 1; sum <= totalSum; sum++ {
        dp[n][sum] = false
    }
    for i := n-1; i >= 0; i-- {
        for sum := 1; sum <= totalSum; sum++ {
            var res bool
            if sum >= nums[i] {
                res = dp[i+1][sum-nums[i]]
            }
            dp[i][sum] = res || dp[i+1][sum]
        }
    }
    return dp[0][totalSum]
}

---

func canPartition(nums []int) bool {
    sum := 0
    for _, n := range nums { sum += n }
    if sum % 2 != 0 { return false }
    return helper(nums, 0, sum/2) 
}

func helper(nums []int, i, sum int) bool {
    dp := make([][]bool, len(nums)+1)
    for i := range dp {
        dp[i] = make([]bool, sum+1)
        dp[i][0] = true
    }
    for i := len(nums)-1; i >= 0; i-- {
        for s := 1; s <= sum; s++ {
            res := false
            if s - nums[i] >= 0 {
                res = dp[i+1][s-nums[i]]
            }
            dp[i][s] = res || dp[i+1][s]
        }
    }
    return dp[0][sum]
}
