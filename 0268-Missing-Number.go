func missingNumber(nums []int) int {
    var sum int
    for _, n := range nums { sum += n }
    return len(nums) * (len(nums) + 1)/2 - sum
}
