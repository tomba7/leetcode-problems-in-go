func findPeakElement(nums []int) int {
    lo, hi := 0, len(nums)-1
    for lo < hi {
        mid := lo + (hi-lo)/2
        if nums[mid] > nums[mid+1] {
            hi = mid
        } else {
            lo = mid+1
        }
    }
    return lo
}

---

/* Brute force */
func findPeakElement(nums []int) int {
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] > nums[i+1] {
            return i
        }
    }
    return len(nums)-1
}

---

/* Brute force: Bad, non-elegant solution */
func findPeakElement(nums []int) int {
    for i := 0; i < len(nums); i++ {
        if (i == 0 || nums[i] > nums[i - 1]) && (i + 1 == len(nums) || nums[i] > nums[i + 1]) {
            return i
        }
    }
    return -1
}


