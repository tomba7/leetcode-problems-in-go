func containsDuplicate(nums []int) bool {
    set := make(map[int]struct{})
    for _, n := range nums {
        if _, ok := set[n]; ok {
            return true
        }
        set[n] = struct{}{}
    }
    return false
}

---

func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
    for i := 0; i < len(nums); i++ {
        if i != 0 && nums[i] == nums[i-1] {
            return true
        }
    }
    return false
}

---

func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] == nums[i + 1] { return true }
    }
    return false
}
