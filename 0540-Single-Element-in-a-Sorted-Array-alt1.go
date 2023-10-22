func singleNonDuplicate(nums []int) int {
    if len(nums) == 0 || len(nums) % 2 == 0 {
        return -1
    } else if len(nums) == 1 {
        return nums[0]
    }
    lo, hi := 0, len(nums) - 1
    for lo < hi {
        mid := lo + (hi - lo)/2
        if mid % 2 != 0 {
            if nums[mid] == nums[mid + 1] {
                hi = mid - 1
            } else {
                lo = mid + 1
            }
        } else {
            if nums[mid - 1] != nums[mid] && nums[mid] != nums[mid + 1] {
                return mid
            } else if nums[mid] == nums[mid + 1] {
                lo = mid + 2
            } else {
                hi = mid - 2
            }
        }
    }
    return nums[lo]
}
