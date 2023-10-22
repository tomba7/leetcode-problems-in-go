func threeSumClosest(nums []int, target int) int {
    // Input size is guaranteed to be at least three.
    sort.Ints(nums)
    closest := nums[0] + nums[1] + nums[2]
    for i := 0; i < len(nums) - 2; i++ {
        j, k := i + 1, len(nums) - 1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            if sum < target {
                j++
            } else if sum > target {
                k--
            } else {
                return sum
            }
            if abs(target - sum) < abs(target - closest) { closest = sum }
        }
    }
    return closest
}

func abs(x int) int { if x < 0 { return -x }; return x }
