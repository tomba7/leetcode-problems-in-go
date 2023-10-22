type Solution struct {
    nums []int
}

func Constructor(nums []int) Solution {
    return Solution{nums: nums}
}

func (s *Solution) Reset() []int {
    return s.nums
}

func (s *Solution) Shuffle() []int {
    nums := append([]int{}, s.nums...)
    for i := 0; i < len(nums); i++ {
        j := i + rand.Intn(len(nums)-i)
        nums[i], nums[j] = nums[j], nums[i]
    }
    return nums
}

---

type Solution []int

func Constructor(nums []int) Solution {
    return Solution(nums)
}

func (s *Solution) Reset() []int {
    return *s
}

func (s *Solution) Shuffle() []int {
    nums := make([]int, len(*s))
    copy(nums, *s)
    for i := 0; i < len(nums)-1; i++ {
        j := i + rand.Intn(len(nums)-i)
        nums[i], nums[j] = nums[j], nums[i]
    }
    return nums
}

---

// Brute force
type Solution struct {
    nums []int
    array []int
}

func Constructor(nums []int) Solution {
    return Solution{
        nums: nums,
        array: append([]int{}, nums...),
    }
}

func (s *Solution) Reset() []int {
    s.array = append([]int{}, s.nums...)
    return s.array
}

func (s *Solution) Shuffle() []int {
    aux := append([]int{}, s.array...)
    for i := 0; i < len(s.array); i++ {
        removeIndex := rand.Intn(len(aux))
        s.array[i] = aux[removeIndex]
        aux = append(aux[:removeIndex], aux[removeIndex+1:]...)
    }
    return s.array
}
