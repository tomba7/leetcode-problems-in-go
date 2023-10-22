func searchInsert(nums []int, target int) int {
    lo, hi := 0, len(nums) - 1
    for lo < hi {
        mid := lo + (hi - lo)/2
        if nums[mid] < target {
            lo = mid + 1
        } else {
            hi = mid
        }
    }
    if target > nums[lo] { return lo + 1 }
    return lo
}
