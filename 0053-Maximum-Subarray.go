func maxSubArray(nums []int) int {
    if len(nums) <= 0 { return 0 }
    maxSum, sum := nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        sum = max(sum + nums[i], nums[i])
        maxSum = max(maxSum, sum)
    }
    return maxSum
}

func max(x, y int) int { if x > y { return x }; return y }

---

func maxSubArray(nums []int) int {
    return maxSubArrayHelper(nums, 0, len(nums)-1)
}

func maxSubArrayHelper(nums []int, i, j int) int {
    if i > j { return math.MinInt32 }
    mid := i + (j-i)/2
    left := maxSubArrayHelper(nums, i, mid-1)
    right := maxSubArrayHelper(nums, mid+1, j)
    both := crossHelper(nums, i, mid, j)
    return max(left, max(both, right))
}

func crossHelper(nums []int, lo, mid, hi int) int {
    leftSum, leftMax := 0, math.MinInt32
    for i := mid; i >= lo; i-- {
        leftSum += nums[i]
        leftMax = max(leftMax, leftSum)
    }
    rightSum, rightMax := 0, 0
    for i := mid+1; i <= hi; i++ {
        rightSum += nums[i]
        rightMax = max(rightMax, rightSum)
    }
    return leftMax + rightMax
}

func max(x, y int) int { if x > y { return x }; return y }

---

func maxSubArray(nums []int) int {
    return maxSubArrayHelper(nums, 0, len(nums)-1)
}

func maxSubArrayHelper(nums []int, i, j int) int {
    if i > j { return math.MinInt32 }
    if i == j { return nums[i] }
    mid := i + (j-i)/2
    left := maxSubArrayHelper(nums, i, mid)
    right := maxSubArrayHelper(nums, mid+1, j)
    both := crossHelper(nums, i, mid, j)
    return max(left, max(both, right))
}

func crossHelper(nums []int, lo, mid, hi int) int {
    leftSum, leftMax := 0, math.MinInt32
    for i := mid; i >= lo; i-- {
        leftSum += nums[i]
        leftMax = max(leftMax, leftSum)
    }
    rightSum, rightMax := 0, 0
    for i := mid+1; i <= hi; i++ {
        rightSum += nums[i]
        rightMax = max(rightMax, rightSum)
    }
    return leftMax + rightMax
}

func max(x, y int) int { if x > y { return x }; return y }
