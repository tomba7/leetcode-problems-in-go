type NumArray []int

func Constructor(nums []int) NumArray {
    for i := 1; i < len(nums); i++ {
        nums[i] += nums[i - 1]
    }
    return nums
}

func (na *NumArray) SumRange(i, j int) int {
    if i != 0 {
        return (*na)[j] - (*na)[i - 1]
    }
    return (*na)[j]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
