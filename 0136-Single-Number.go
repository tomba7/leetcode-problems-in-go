func singleNumber(nums []int) int {
    var result int
    for _, n := range nums {
        result ^= n
    }
    return result
}

---

func singleNumber(nums []int) int {
    set := make(map[int]struct{})
    for _, n := range nums {
        if _, ok := set[n]; ok {
            delete(set, n)
        } else {
            set[n] = struct{}{}
        }
    }
    for k, _ := range set {
        return k
    }
    return -1
}
