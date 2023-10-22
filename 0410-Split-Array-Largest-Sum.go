// Recursion - TLE
func splitArray(nums []int, m int) int {
	res := math.MaxInt32
	splitArrayHelper(nums, 0, m, 0, 0, &res)
	return res
}

func splitArrayHelper(nums []int, i, id, sum, maxSum int, res *int) {
	if i == len(nums) {
		if id == 0 {
			*res = min(*res, maxSum)
		}
		return
	}
	if i > 0 {          // Continue building the same sub-array
		splitArrayHelper(nums, i+1, id, sum+nums[i], max(maxSum, sum+nums[i]), res)
	}
	if id > 0 {         // Branch out and create a new sub-array
		splitArrayHelper(nums, i+1, id-1, nums[i], max(maxSum, nums[i]), res)
	}
}

func min(x, y int) int { if x < y { return x }; return y }
func max(x, y int) int { if x > y { return x }; return y }
