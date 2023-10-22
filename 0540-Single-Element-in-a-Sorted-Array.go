func singleNonDuplicate(nums []int) int {
    // The size of 'nums' should always be odd. Validate it here.
    if len(nums) % 2 == 0 {panic("invalid input")}
    lo, hi := 0, len(nums) - 1
    for lo < hi {
        mid := lo + (hi - lo)/2
        // Inside this loop, the size of 'nums' is always guaranteed to be
        // equal to or greater than 3.
        if nums[mid] != nums[mid - 1] && nums[mid] != nums[mid + 1] {
            return nums[mid]
        } else if (mid % 2 != 0 && nums[mid] == nums[mid + 1]) || (mid % 2 == 0 && nums[mid] != nums[mid + 1]) {
            hi = mid - 1
        } else {
            // (mid % 2 != 0 && nums[mid] != nums[mid + 1]) || (mid % 2 == 0 && nums[mid] == nums[mid + 1)
            lo = mid + 1
        }
    }
    return nums[lo]
}
