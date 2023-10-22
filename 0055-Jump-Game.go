/*  Greedy
    2,3,1,1,4
    i
    j

    [3,2,1,0,4]
     i
             j
 */
func canJump(nums []int) bool {
    j := len(nums)-1
    for i := j; i >= 0; i-- {
        if i + nums[i] >= j {
            j = i
        }
    }
    return j == 0
}

---

func canJump(nums []int) bool {
    furthest := 0
    for i := 0; i <= furthest && furthest < len(nums); i++ {
        furthest = max(furthest, i + nums[i])
    }
    return furthest >= len(nums) - 1
}

func max(x, y int) int { if x > y { return x }; return y }

---

// DP
func canJump(nums []int) bool {
    n := len(nums)
    dp := make([]bool, n)
    dp[n-1] = true
    for i := n-2; i >= 0; i-- {
        for j := 1; j <= nums[i]; j++ {
            if i+j < n {
                dp[i] = dp[i] || dp[i+j]
                if dp[i] { break }
            }
        }
    }
    return dp[0]
}

---

// Recursion (TLE)
func canJump(nums []int) bool {
    return canJumpHelper(nums, 0)
}

func canJumpHelper(nums []int, i int) bool {
    if i == len(nums)-1 {
        return true
    }
    for j := 1; j <= nums[i]; j++ {
        if canJumpHelper(nums, i+j) {
            return true
        }
    }
    return false
}
