func subsets(nums []int) [][]int {
    var res [][]int
    pref := make([]int, len(nums))
    subsetsHelper(nums, 0, pref, 0, &res)
    return res
}

func subsetsHelper(nums []int, i int, pref []int, j int, res *[][]int) {
    if i == len(nums) {
        *res = append(*res, append([]int{}, pref[:j]...))
        return
    }
    subsetsHelper(nums, i+1, pref, j, res)
    pref[j] = nums[i]
    subsetsHelper(nums, i+1, pref, j+1, res)
}

---

func subsets(nums []int) [][]int {
    res := [][]int{{}}
    for i := 0; i < len(nums); i++ {
        size := len(res)
        for j := 0; j < size; j++ {
            clone := append([]int{}, res[j]...)
            res = append(res, append(clone, nums[i]))
        }
    }
    return res
}

---

func subsets(nums []int) (result [][]int) {
    subsetsHelper(nums, nil, 0, &result)
    return result
}

func subsetsHelper(nums, sub []int, index int, result *[][]int) {
    *result = append(*result, append([]int{}, sub...))
    for i := index; i < len(nums); i++ {
        subsetsHelper(nums, append(sub, nums[i]), i + 1, result)
    }
}

---

func subsets(nums []int) [][]int {
    return subsetsHelper(nums, 0)
}

func subsetsHelper(nums []int, index int) [][]int {
    if len(nums) == index {
        return [][]int{nil}
    }
    result := subsetsHelper(nums, index + 1)
    var tmp [][]int
    for _, subset := range result {
        new := []int{nums[index]}
        new = append(new, subset...)
        tmp = append(tmp, new)
    }
    return append(result, tmp...)
}
