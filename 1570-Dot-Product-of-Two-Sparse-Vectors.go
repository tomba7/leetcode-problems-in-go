type SparseVector struct {
    mapping map[int]int
}

func Constructor(nums []int) SparseVector {
    mapping := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            mapping[i] = nums[i]
        }
    }
    return SparseVector{mapping: mapping}
}

func (v *SparseVector) dotProduct(vec SparseVector) int {
    var res int
    for i, val := range v.mapping {
        res += val * vec.mapping[i]
    }
    return res
}

---

type SparseVector struct {
    pairs [][]int
}

func Constructor(nums []int) SparseVector {
    var pairs [][]int
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            pairs = append(pairs, []int{i, nums[i]})
        }
    }
    return SparseVector{pairs: pairs}
}

func (v *SparseVector) dotProduct(vec SparseVector) int {
    var res, i, j int
    for i < len(v.pairs) && j < len(vec.pairs) {
        if v.pairs[i][0] == vec.pairs[j][0] {
            res += v.pairs[i][1] * vec.pairs[j][1]
            i++;
            j++
        } else if v.pairs[i][0] < vec.pairs[j][0] {
            i++
        } else {
            j++
        }
    }
    return res
}

---

type SparseVector struct {
    nums []int
}

func Constructor(nums []int) SparseVector {
    return SparseVector{nums: nums}
}

func (v *SparseVector) dotProduct(vec SparseVector) int {
    if len(v.nums) != len(vec.nums) { return 0 }
    var res int
    for i := 0; i < len(v.nums); i++ {
        res += v.nums[i] * vec.nums[i]
    }
    return res
}

---

type SparseVector []int

func Constructor(nums []int) SparseVector {
    return SparseVector(nums)
}

func (v *SparseVector) dotProduct(vec SparseVector) int {
    if len(*v) != len(vec) { return 0 }
    var res int
    for i := 0; i < len(*v); i++ {
        res += (*v)[i] * vec[i]
    }
    return res
}
