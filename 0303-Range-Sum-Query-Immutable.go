type NumArray struct {
    nums []int
}

func Constructor(nums []int) NumArray {
    na := NumArray{nums: append([]int{}, nums...)}
    for i := 1; i < len(na.nums); i++ {
        na.nums[i] += na.nums[i - 1]
    }
    return na
}

func (na *NumArray) SumRange(i int, j int) int {
    if i != 0 {
        return na.nums[j] - na.nums[i - 1]
    }
    return na.nums[j]
}
