// Divide and Conquer
func maxSubArray(nums []int) int {
    return maxSubArrayHelper(nums, 0, len(nums) - 1)
}

func maxSubArrayHelper(nums []int, lo, hi int) int {
    if lo == hi { return nums[lo] }
    mid := lo + (hi - lo)/2
    return max3(maxSubArrayHelper(nums, lo, mid),
                maxSubArrayHelper(nums, mid + 1, hi),
                maxSubArrayBothSides(nums, lo, mid, hi))
}

func maxSubArrayBothSides(nums []int, lo, mid, hi int) int {
    leftMax, rightMax := int(math.MinInt32), int(math.MinInt32)
    sum := 0
    for i := mid; i >= lo; i-- {
        sum += nums[i]
        leftMax = max2(leftMax, sum)
    }
    sum = 0
    for i := mid + 1; i <= hi; i++ {
        sum += nums[i]
        rightMax = max2(rightMax, sum)
    }
    return leftMax + rightMax
}

func max2(x, y int) int { if x > y { return x }; return y }
func max3(x, y, z int) int { return max2(max2(x, y), z) }
