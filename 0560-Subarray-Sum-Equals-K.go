func subarraySum(nums []int, k int) int {
    freq := map[int]int{0: 1}
    var sum, count int
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if _, exist := freq[sum-k]; exist {
            count += freq[sum-k]
        }
        freq[sum]++
    }
    return count
}

---

func subarraySum(nums []int, k int) (n int) {
    freq := map[int]int{0: 1}
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if count, ok := freq[sum - k]; ok { n += count }
        freq[sum]++
    }
    return n
}

---

func subarraySum(nums []int, k int) int {
    freq := map[int]int{0: 1}
    var res int
    var prefix int
    for i := 0; i < len(nums); i++ {
        prefix += nums[i]
        count, exist := freq[prefix-k]
        if exist {
            res += count
        }
        freq[prefix]++
    }
    return res
}
