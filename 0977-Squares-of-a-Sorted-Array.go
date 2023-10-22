/*
- find the first +ve value or zero. point j to it and i to j-1
- for i >= 0 && j < len(nums)
- compare both s[i]*s[i] and s[j]*s[j]
- pick the max and increment either i or j
- When out of this loop, exhaust any of the remaining values
 */
func sortedSquares(nums []int) []int {
    res := make([]int, len(nums))
    lo, hi := 0, len(nums)-1
    for i := len(nums)-1; i >= 0; i-- {
        x, y := nums[lo] * nums[lo], nums[hi] * nums[hi]
        if x > y {
            res[i] = x
            lo++
        } else {
            res[i] = y
            hi--
        }
    }
    return res
}

---

/*
  Non optimal approach:

- find the first +ve value or zero. point j to it and i to j-1
- for i >= 0 && j < len(nums)
- compare both s[i]*s[i] and s[j]*s[j]
- pick the max and increment either i or j
- When out of this loop, exhaust any of the remaining values
 */
func sortedSquares(nums []int) []int {
    var j int
    for ; j < len(nums) && nums[j] < 0; j++ {}
    i := j-1
    var res []int
    for i >= 0 && j < len(nums) {
        x, y := nums[i] * nums[i], nums[j] * nums[j]
        if x < y {
            res = append(res, x)
            i--
        } else {
            res = append(res, y)
            j++
        }
    }
    for ; i >= 0; i-- {
        res = append(res, nums[i] * nums[i])
    }
    for ; j < len(nums); j++ {
        res = append(res, nums[j] * nums[j])
    }
    return res
}
