func findPairs(nums []int, k int) int {
    numPairs := 0
    sort.Ints(nums)
    for i := 0; i < len(nums) - 1; i++ {
        if i != 0 && nums[i] == nums[i - 1] {
            continue
        }
        for j := i + 1; j < len(nums) && nums[j] - nums[i] <= k; j++ {
            if nums[j] - nums[i] == k {
                numPairs++
                break
            }
        }
    }
    return numPairs
}
