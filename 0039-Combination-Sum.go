func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    var res [][]int
    combinationSumHelper(candidates, target, 0, 0, nil, &res)
    return res 
}

func combinationSumHelper(candidates []int, target, i, sum int, comb []int, res *[][]int) {
    if sum >= target {
        if sum == target {
            *res = append(*res, append([]int{}, comb...))
        }
        return
    }
    for j := i; j < len(candidates); j++ {
        combinationSumHelper(candidates, target, j, sum + candidates[j], append(comb, candidates[j]), res)
    }
}

---

func combinationSum(nums []int, target int) (result [][]int) {
    sort.Ints(nums)
    combinationSumHelper(nums, []int{}, target, 0, 0, &result)
    return result
}

func combinationSumHelper(nums, subset []int, target, startIndex, sum int, result *[][]int) {
    if sum == target {
        *result = append(*result, append([]int{}, subset...))
        return
    }
    for i := startIndex; i < len(nums) && sum + nums[i] <= target; i++ {
        combinationSumHelper(nums, append(subset, nums[i]), target, i, sum + nums[i], result)
    }
}
