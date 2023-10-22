func maxSubArrayLen(nums []int, k int) int {
    if len(nums) == 0 {return 0}
    table := map[int]int{0: -1}
    maxLen := 0
    for i := 1; i < len(nums); i++ {nums[i] += nums[i - 1]}
    for i := 0; i < len(nums); i++ {
        if _, ok := table[nums[i] - k]; ok {
            maxLen = max(maxLen, i - table[nums[i] - k])
        }
        if _, ok := table[nums[i]]; !ok {table[nums[i]] = i}
    }
    return maxLen
}

func max(x, y int) int {if x > y {return x}; return y}
