func findDuplicate(nums []int) int {
    lo, hi := 0, len(nums)-1
    for lo < hi {
        mid := lo + (hi-lo)/2
        // Count the number of elements less than or equal to mid
        // If count is greater than mid, the duplicate number
        // falls between [lo ~ mid]. If count is less than or equal
        // to mid, the dupicate number falls between [mid+1 ~ hi]
        count := 0
        for _, n := range nums {
            if n <= mid { count++ }
        }
        if count > mid {
            hi = mid
        } else {
            lo = mid+1
        }
    }
    return lo
}

---

func findDuplicate(nums []int) int {
    sort.Ints(nums)
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            return nums[i]
        }
    }
    return -1
}

---

func findDuplicate(nums []int) int {
    set := make(map[int]struct{})
    for _, n := range nums {
        if _, ok := set[n]; ok {
            return n
        }
        set[n] = struct{}{}
    }
    return -1
}
