func permuteUnique(nums []int) (result [][]int) {
    sort.Ints(nums)
    permuteUniqueHelper(nums, 0, &result)
    return result
}

func permuteUniqueHelper(nums []int, start int, result *[][]int) {
    if len(nums) - 1 == start {
        *result = append(*result, append([]int{}, nums...))
        return
    }
    for i := start; i < len(nums); i++ {
        if i != start && nums[i] == nums[start] {continue}
        nums[i], nums[start] = nums[start], nums[i]
        permuteUniqueHelper(append([]int{}, nums...), start + 1, result)
    }
}
