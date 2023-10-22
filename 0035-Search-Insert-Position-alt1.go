func searchInsert(nums []int, target int) int {
    result := 0
    lo, hi := 0, len(nums) - 1
    for lo <= hi {
        mid := lo + (hi - lo)/2
        if target < nums[mid] {
            result = mid
            hi = mid - 1
        } else if target > nums[mid] {
            result = mid + 1
            lo = mid + 1
        } else {
            return mid
        }
    }
    return result
}
