type Solution struct {
    prefix []int
    total int
}


func Constructor(w []int) Solution {
    prefix := make([]int, len(w))
    var sum int
    for i := 0; i < len(w); i++ {
        sum += w[i]
        prefix[i] = sum
    }
    return Solution{
        prefix: prefix, total: sum,
    }
}


func (s *Solution) PickIndex() int {
    target := rand.Intn(s.total) + 1
    lo, hi := 0, len(s.prefix)-1
    for lo < hi {
        mid := lo + (hi-lo)/2
        if target <= s.prefix[mid] {
            hi = mid
        } else {
            lo = mid + 1
        }
    }
    return lo
}

---

type Solution struct {
    prefix []int
    total int
}

func Constructor(w []int) Solution {
    prefix := make([]int, len(w))
    var sum int
    for i := 0; i < len(w); i++ {
        sum += w[i]
        prefix[i] = sum
    }
    return Solution{prefix: prefix, total: sum}
}

func (s *Solution) PickIndex() int {
    target := rand.Intn(s.total) + 1
    var i int
    for ; i < len(s.prefix); i++ {
        if target <= s.prefix[i] {
            return i
        }
    }
    return -1
}
