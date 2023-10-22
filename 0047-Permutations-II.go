func permuteUnique(nums []int) [][]int {
    var res [][]int
    permuteUniqueHelper(nums, 0, &res)
    return res
}

func permuteUniqueHelper(nums []int, i int, res *[][]int) {
    if i == len(nums)-1 {
        *res = append(*res, append([]int{}, nums...))
        return
    }
    m := map[int]int{}
    for j := i; j < len(nums); j++ {
        if m[nums[j]] > 0 { continue }
        swap(nums, i, j)
        permuteUniqueHelper(nums, i+1, res)
        swap(nums, i, j)
        m[nums[j]]++
    }
}

func swap(nums []int, i, j int) { nums[i], nums[j] = nums[j], nums[i] }

---

func permuteUnique(nums []int) (result [][]int) {
    permuteUniqueHelper(nums, 0, &result)
    return
}

func permuteUniqueHelper(nums []int, index int, result *[][]int) {
    if index == len(nums) - 1 {
        *result = append(*result, append([]int{}, nums...))
    }
    freq := map[int]struct{}{}
    for i := index; i < len(nums); i++ {
        if _, ok := freq[nums[i]]; ok { continue }
        nums[index], nums[i] = nums[i], nums[index]
        permuteUniqueHelper(nums, index + 1, result)
        nums[index], nums[i] = nums[i], nums[index]
        freq[nums[i]] = struct{}{}
    }
}
