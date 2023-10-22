// Time: O(n), Space: O(n)
func findDuplicates(nums []int) []int {
    freq := make([]int, len(nums))
    var res []int
    for i := 0; i < len(nums); i++ {
        freq[nums[i]-1]++
        if freq[nums[i]-1] > 1 {
            res = append(res, nums[i])
        }
    }
    return res
}

---

// Time: O(n logn), Space: O(1)
func findDuplicates(nums []int) []int {
    var res []int
    sort.Ints(nums)
    for i := 0; i < len(nums); i++ {
        if i != 0 && nums[i] == nums[i-1] {
            res = append(res, nums[i])
        }
    }
    return res
}
