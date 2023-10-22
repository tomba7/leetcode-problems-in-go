func permuteUnique(nums []int) (result [][]int) {
    permuteUniqueHelper(nums, 0, &result)
    return
}

func permuteUniqueHelper(nums []int, index int, result *[][]int) {
    if index == len(nums) - 1 {
        *result = append(*result, append([]int{}, nums...))
    }
    freq := [20]bool{}
    for i := index; i < len(nums); i++ {
        if freq[nums[i] + 10] { continue }
        nums[index], nums[i] = nums[i], nums[index]
        permuteUniqueHelper(nums, index + 1, result)
        nums[index], nums[i] = nums[i], nums[index]
        freq[nums[i] + 10] = true
    }
}
