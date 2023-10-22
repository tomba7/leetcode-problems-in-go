func nextPermutation(nums []int)  {
    for i := len(nums)-2; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            j := len(nums)-1
            for ; j > i && nums[j] <= nums[i]; j-- {}
            swap(nums, i, j)
            reverse(nums, i+1, len(nums)-1)
            return
        }
    }
    reverse(nums, 0, len(nums)-1)
}

func swap(nums []int, i, j int) { nums[i], nums[j] = nums[j], nums[i] }
func reverse(nums []int, i, j int) { for ; i < j; i, j = i+1, j-1 { swap(nums, i, j) } }

---

func nextPermutation(nums []int)  {
    var i int
    for i = len(nums) - 1; i > 0 && nums[i - 1] >= nums[i]; i-- {}
    if i == 0 {
        reverse(nums, 0)
        return
    }
    subStart := i
    for i = len(nums) - 1; i >= subStart && nums[i] <= nums[subStart - 1]; i-- {}
    nums[subStart - 1], nums[i] = nums[i], nums[subStart - 1]
    reverse(nums, subStart)
}

func reverse(nums []int, start int) {
    for i, j := start, len(nums) - 1; i < j; i, j = i + 1, j - 1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
}
