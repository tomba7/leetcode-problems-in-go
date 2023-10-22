/*
    [23, 2, 4, 6, 7]
     23,25,29,35,42
    
    [ 3, 1, 4, 7, 3 ]
      3, 4, 8,15,18
 */
func checkSubarraySum(nums []int, k int) bool {
    index := map[int]int{0: -1}
    var sum int
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        sum %= k
        if j, exist := index[sum]; exist {
            if i - j > 1 { return true }
        } else {
            index[sum] = i
        }
    }
    return false
}

---

/*
    k = 6
    subarray(2 elem or more) with sum = n * k
    
    23, 2,  4,  6,  7
    23, 25, 29, 35, 42
 */
func checkSubarraySum(nums []int, k int) bool {
    prefix := map[int]int{0:-1}
    var sum int
    for j, n := range nums {
        sum += n
        if i, exist := prefix[sum%k]; exist {
            if j - i > 1 { return true }
        } else {
            prefix[sum%k] = j
        }
    }
    return false
}
