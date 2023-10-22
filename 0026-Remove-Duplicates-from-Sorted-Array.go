func removeDuplicates(nums []int) int {
    var i int
    for j := 0; j < len(nums); j++ {
        if j == 0 || nums[j] != nums[j-1] {
            nums[i] = nums[j]
            i++
        }
    }
    return i
}

---

func removeDuplicates(nums []int) int {
    var i int
    for j := 0; j < len(nums); j++ {
        if j != 0 && nums[j] == nums[j-1] { continue }
        nums[i] = nums[j]
        i++
    }
    return i
}
