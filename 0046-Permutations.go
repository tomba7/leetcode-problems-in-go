func permute(nums []int) (result [][]int) {
    permuteHelper(nums, 0, &result)
    return result
}

func permuteHelper(nums []int, index int, result *[][]int) {
    if index == len(nums) - 1 {
        *result = append(*result, append([]int{}, nums...))
        return
    }
    for i := index; i < len(nums); i++ {
        nums[index], nums[i] = nums[i], nums[index]
        permuteHelper(nums, index + 1, result)
        nums[index], nums[i] = nums[i], nums[index]
    }
}
