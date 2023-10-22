func subsetsWithDup(nums []int) [][]int {
    var res [][]int
    sort.Ints(nums)
    subset := make([]int, len(nums))
    subsetsWithDupHelper(nums, 0, subset, 0, &res)
    return res
}

func subsetsWithDupHelper(nums []int, i int, subset []int, j int, res *[][]int) {
    if i == len(nums) {
        *res = append(*res, clone(subset[:j]))
        return
    }
    k := nextUniqueIndex(nums, i)
    subsetsWithDupHelper(nums, k, subset, j, res)
    subset[j] = nums[i]
    subsetsWithDupHelper(nums, i+1, subset, j+1, res)
}

func nextUniqueIndex(nums []int, i int) int {
    i++
    for ; i < len(nums) && nums[i] == nums[i-1]; i++ {}
    return i
}

func clone(subset []int) []int { return append([]int{}, subset...) }

---

func subsetsWithDup(nums []int) [][]int {
    var res [][]int
    sort.Ints(nums)
    helper(nums, 0, make([]int, len(nums)), 0, &res)
    return res
}

func helper(nums []int, i int, sub []int, j int, res *[][]int) {
    if i == len(nums) {
        dup := append([]int{}, sub[:j]...)
        *res = append(*res, dup)
        return
    }
    next := i + 1
    for next < len(nums) && nums[i] == nums[next] { next++ }
    helper(nums, next, sub, j, res)
    sub[j] = nums[i]
    helper(nums, i+1, sub, j+1, res)
}

---

func subsetsWithDup(nums []int) [][]int {
    var res [][]int
    sort.Ints(nums)
    helper(nums, 0, []int{}, &res)
    return res
}

func helper(nums []int, i int, tmp []int, res *[][]int) {
    if i == len(nums) {
        *res = append(*res, append([]int{}, tmp...))
        return
    }
    j := i+1
    for j < len(nums) && nums[i] == nums[j]  { j++ }
    helper(nums, j, tmp, res)
    helper(nums, i+1, append(tmp, nums[i]), res)
}

---

func subsetsWithDup(nums []int) [][]int {
    res := [][]int{{}}
    sort.Ints(nums)
    var start int
    for i := 0; i < len(nums); i++ {
        j := 0
        if i != 0 && nums[i] == nums[i-1] {
            j = start
        }
        n := len(res)
        start = n
        for ; j < n; j++ {
            tmp := append(
                append([]int{}, res[j]...), nums[i],
            )
            res = append(res, tmp)
        }
    }
    return res
}

---

func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    var res [][]int
    subsetsWithDupHelper(nums, 0, make([]int, len(nums)), 0, &res)
    return res
}

func subsetsWithDupHelper(nums []int, i int, sub []int, j int, res *[][]int) {
    if i == len(nums) {
        *res = append(*res, append([]int{}, sub[:j]...))
        return
    }
    skip := i
    // For efficiency we should compute the skip index in advance in main
    for ; skip < len(nums) && nums[skip] == nums[i]; skip++ {} 
    subsetsWithDupHelper(nums, skip, sub, j, res)
    sub[j] = nums[i]
    subsetsWithDupHelper(nums, i+1, sub, j+1, res)
}

---

func subsetsWithDup(nums []int) (res [][]int) {
    sort.Ints(nums)
    subsetsHelper(nums, 0, nil, &res)
    return res
}

func subsetsHelper(nums []int, i int, sub []int, res *[][]int) {
    *res = append(*res, append([]int{}, sub...))
    if i == len(nums) { return }
    for j := i; j < len(nums); j++ {
        if j != i && nums[j] == nums[j - 1] { continue }
        subsetsHelper(nums, j+1, append(sub, nums[j]), res)
    }
}

---

func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    var res [][]int
    subsetsWithDupHelper(nums, 0, make([]int, len(nums)), 0, &res)
    return res
}

func subsetsWithDupHelper(nums []int, i int, sub []int, j int, res *[][]int) {
    if i == len(nums) {
        *res = append(*res, append([]int{}, sub[:j]...))
        return
    }
    skip := i
    for ; skip < len(nums) && nums[skip] == nums[i]; skip++ {}
    subsetsWithDupHelper(nums, skip, sub, j, res)
    for k := i; k < skip; k++ {
        sub[j] = nums[i]
        j++
        subsetsWithDupHelper(nums, skip, sub, j, res)
    }
}

---

func subsetsWithDup(nums []int) [][]int {
    var res [][]int
    sort.Ints(nums)
    subset := make([]int, len(nums))
    subsetsWithDupHelper(nums, 0, subset, 0, &res)
    return res
}

func subsetsWithDupHelper(nums []int, i int, subset []int, j int, res *[][]int) {
    if i == len(nums) {
        *res = append(*res, clone(subset[:j]))
        return
    }
    k := nextUniqueIndex(nums, i 
    subsetsWithDupHelper(nums, k, subset, j, res)
    for ; i < k; i++ {
        subset[j] = nums[i]
        j++
        subsetsWithDupHelper(nums, k, subset, j, res)
    }
}

func nextUniqueIndex(nums []int, i int) int {
    i++
    for ; i < len(nums) && nums[i] == nums[i-1]; i++ {}
    return i
}

func clone(subset []int) []int {
    return append([]int{}, subset...)
}
