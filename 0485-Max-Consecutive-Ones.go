func findMaxConsecutiveOnes(nums []int) int {
    var i, j, maxLen int
    for ; j <= len(nums); j++ {
        if j == len(nums) || nums[j] == 0 {
            maxLen = max(maxLen, j - i)
            i = j+1
        }
    }
    return maxLen
}

func max(x, y int) int { if x > y { return x }; return y }

---

func findMaxConsecutiveOnes(nums []int) int {
    var maxLen, currLen int
    for _, n := range nums {
        if n == 0 {
            currLen = 0
        } else {
            currLen++
            maxLen = max(maxLen, currLen)
        }
    }
    return maxLen
}

func max(x, y int) int { if x > y { return x }; return y }
