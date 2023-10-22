func combinationSum2(nums []int, target int) (result [][]int) {
    sort.Ints(nums)
    combinationSum2Helper(nums, nil, target, 0, 0, &result)
    return result
}

func combinationSum2Helper(nums, combo []int, target, sum, startIndex int, result *[][]int) {
    if sum == target {
        *result = append(*result, append([]int{}, combo...))
        return
    }
    for i := startIndex; i < len(nums) && (sum + nums[i]) <= target; i++ {
        if i != startIndex && nums[i] == nums[i - 1] { continue }
        combinationSum2Helper(nums, append(combo, nums[i]), target, sum + nums[i], i + 1, result)
    }
}

---

func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    var res [][]int 
    combinationSum2Helper(candidates, target, 0, 0, nil, &res)
    return res
}

func combinationSum2Helper(candidates []int, target, i, sum int, comb []int, res *[][]int) {
    if i == len(candidates) || sum >= target {
        if sum == target {
            *res = append(*res, append([]int{}, comb...))
        }
        return
    }
    var j = i+1
    for ; j < len(candidates) && candidates[j] == candidates[j-1]; j++ {}
    combinationSum2Helper(candidates, target, j, sum, comb, res)
    combinationSum2Helper(
        candidates, target, i+1, sum + candidates[i], append(comb, candidates[i]), res,
    )
}
