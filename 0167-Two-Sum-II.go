func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums) - 1; i++ {
        indx := binarySearch(nums, target - nums[i], i + 1)
        if indx >= 0 {
            return []int{i + 1, indx + 1}
        }
    }
    return nil
}

func binarySearch(nums []int, target, lo int) int {
    hi := len(nums) - 1
    for lo < hi {
        mid := lo + (hi - lo)/2
        if nums[mid] < target {
            lo = mid + 1
        } else {
            hi = mid
        }
    }
    if lo == hi && nums[lo] == target { return lo }
    return -1
}
