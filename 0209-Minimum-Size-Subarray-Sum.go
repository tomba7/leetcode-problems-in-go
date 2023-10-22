func minSubArrayLen(s int, nums []int) int {
    var i, j, sum int
    result := math.MaxInt32
    for j < len(nums) {
        sum += nums[j]
        if sum >= s {
            result = min(result, j - i + 1)
            sum -= nums[i] + nums[j]
            i++
        } else {
            j++
        }
    }
    if result == math.MaxInt32 { result = 0 }
    return result
}

func min(x, y int) int { if x < y { return x }; return y }

---

func minSubArrayLen(s int, nums []int) int {
    sums := make([]int, len(nums)+1)
    for i := 1; i <= len(nums); i++ {
        sums[i] += sums[i-1] + nums[i-1]
    }
    minLen := math.MaxInt32
    for i := 1; i <= len(nums); i++ {
        toFind := s + sums[i-1]
        j := sort.Search(len(sums), func(i int) bool {
            return sums[i] >= toFind
        })
        if j < len(sums) {
            minLen = min(minLen, j-i+1)
        }
    }
    if minLen == math.MaxInt32 { minLen = 0 }
    return minLen
}

func min(x, y int) int {
    if x < y { return x }
    return y
}

---

/* Brute force approach 1 */
func minSubArrayLen(s int, nums []int) int {
    minLen := math.MaxInt32
    for i := 0; i < len(nums); i++ {
        sum := 0
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            if sum >= s {
                minLen = min(minLen, j-i+1)
                break
            }
        }
    }
    if minLen == math.MaxInt32 { minLen = 0 }
    return minLen
}

func min(x, y int) int { if x < y { return x }; return y }

---

/* Brute force approach 2 */
func minSubArrayLen(s int, nums []int) int {
    if len(nums) == 0 { return 0 }
    sums := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        if i == 0 {
            sums[i] = nums[i]
        } else {
            sums[i] += sums[i-1] + nums[i]
        }
    }
    minLen := math.MaxInt32
    for i := 0; i < len(nums); i++ {
        for j := i; j < len(nums); j++ {
            if sums[j] - sums[i] + nums[i] >= s {
                minLen = min(minLen, j - i + 1)
                break
            }
        }
    }
    if minLen == math.MaxInt32 { minLen = 0 }
    return minLen
}

func min(x, y int) int { if x < y { return x }; return y }
