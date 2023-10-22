// 3:45 PM
//  0. 1. 2. 3. 4. 
// [3, 2, 1, 3, 3]
//
//  1, 2, 2,   3, 3, 3, 3, 3, 3, 3,  3,  3
//. 1, 2, 3,   4, 5, 6, 7, 8, 9, 10, 11, 12

type Solution struct {
    index map[int][]int
}

func Constructor(nums []int) Solution {
    index := map[int][]int{}
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        index[num] = append(index[num], i)
    }
    return Solution{index: index}
}

func (s *Solution) Pick(target int) int {
    n := len(s.index[target])
    return s.index[target][rand.Intn(n)]
}

---

// Random sampling
type Solution []int

func Constructor(nums []int) Solution {
    return Solution(nums)
}

func (s *Solution) Pick(target int) int {
    res := -1
    var count int
    for i := 0; i < len(*s); i++ {
        if (*s)[i] == target {
            count++
            if rand.Intn(count) == 0 {
                res = i
            }
        }
    }
    return res
}
