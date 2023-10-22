func reverse(nums []int, begin, end int) {
    for begin < end {
        nums[begin], nums[end] = nums[end], nums[begin]
        begin++
        end--
    }
}

func rotate(nums []int, k int)  {
    if k <= 0 { return }
    k = k % len(nums)
    reverse(nums, 0, len(nums) - 1)
    reverse(nums, 0, k - 1)
    reverse(nums, k, len(nums) - 1)
}

---

/* Brute force approach 1 */
func rotate(nums []int, k int)  {
    k %= len(nums)
    for i := 0; i < k; i++ {
        prev := nums[len(nums)-1]
        for j := 0; j < len(nums); j++ {
            temp := nums[j]
            nums[j] = prev
            prev = temp
        }
    }
}

---

/* Brute force approach 2 */
func rotate(nums []int, k int)  {
    dup := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        dup[(i+k) % len(nums)] = nums[i]
    }
    for i := 0; i < len(nums); i++ {
        nums[i] = dup[i]
    }
}
