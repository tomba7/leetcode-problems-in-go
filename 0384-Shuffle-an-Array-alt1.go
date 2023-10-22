type Solution struct {
    nums []int
}

func Constructor(nums []int) Solution {
    return Solution{nums: nums}
}

/** Resets the array to its original configuration and return it. */
func (s *Solution) Reset() []int {
    return s.nums
}

/** Returns a random shuffling of the array. */
func (s *Solution) Shuffle() []int {
    nums := append([]int{}, s.nums...)
    for i := len(nums) - 1; i >= 1; i-- {
        n := rand.Intn(i + 1)
        nums[i], nums[n] = nums[n], nums[i]
    }
    return nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
