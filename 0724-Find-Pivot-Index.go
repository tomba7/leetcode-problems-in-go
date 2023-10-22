func pivotIndex(nums []int) int {
	var total, leftSum int
	for _, n := range nums {
		total += n
	}
	for i, n := range nums {
		if leftSum == total-leftSum-n {
			return i
		}
		leftSum += n
	}
	return -1
}

---

// Simple approach
func pivotIndex(nums []int) int {
    left := make([]int, len(nums))
    right := make([]int, len(nums))
    var sum int
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        left[i] = sum
    }
    sum = 0
    for i := len(nums)-1; i >= 0; i-- {
        sum += nums[i]
        right[i] = sum
    }
    for i := 0; i < len(nums); i++ {
        if left[i] == right[i] {
            return i
        }
    }
    return -1
}
