func totalHammingDistance(nums []int) int {
    total := 0
    for i := 0; i < 32; i++ {
        numOnes := 0
        for j := range nums {
            numOnes += nums[j] & 1
            nums[j] >>= 1
        }
        total += numOnes * (len(nums) - numOnes)
    }
    return total
}
