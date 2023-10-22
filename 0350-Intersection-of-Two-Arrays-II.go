func intersect(nums1 []int, nums2 []int) []int {
    if len(nums1) > len(nums2) { return intersect(nums2, nums1) }
    freq := map[int]int{}
    for _, n := range nums1 {
        freq[n]++
    }
    var result []int
    for _, n := range nums2 {
        if freq[n] > 0 {
            result = append(result, n)
            freq[n]--
        }
    }
    return result
}

---

func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    var result []int
    for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
        if nums1[i] < nums2[j] {
            i++
        } else if nums1[i] > nums2[j] {
            j++
        } else {
            result = append(result, nums1[i])
            i, j = i+1, j+1
        }
    }
    return result
}
