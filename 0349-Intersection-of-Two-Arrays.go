func intersection(nums1 []int, nums2 []int) []int {
    if len(nums1) > len(nums2) {
        // We hash the smaller array to optimize space complexity
        return intersection(nums2, nums1)
    }
    set := make(map[int]struct{})
    for _, n := range nums1 {
        set[n] = struct{}{}
    }
    var result []int
    for _, n := range nums2 {
        if _, ok := set[n]; ok {
            result = append(result, n)
            delete(set, n)
        }
    }
    return result
}

---

func intersection(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    var result []int
    for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
        if nums1[i] < nums2[j] {
            i++
        } else if nums1[i] > nums2[j] {
            j++
        } else {
            if i == 0 || nums1[i] != nums1[i-1] {
                result = append(result, nums1[i])
            }
            i, j = i+1, j+1
        }
    }
    return result
}

---

func intersection(nums1, nums2 []int) []int {
    set := make(map[int]bool)
    for i := 0; i < len(nums1); i++ {
        set[nums1[i]] = false
    }
    for i := 0; i < len(nums2); i++ {
        if _, ok := set[nums2[i]]; ok {
            set[nums2[i]] = true
        }
    }
    result := []int{}
    for num, intersect := range set {
        if intersect {
            result = append(result, num)
        }
    }
    return result
}
