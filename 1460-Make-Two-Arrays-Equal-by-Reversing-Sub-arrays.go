func canBeEqual(target, arr []int) bool {
    freq := map[int]int{}
    for i := 0; i < len(arr); i++ {
        freq[target[i]]++
        freq[arr[i]]--
    }
    for _, count := range freq {
        if count != 0 { return false }
    }
    return true
}

---

func canBeEqual(target, arr []int) bool {
    sort.Ints(target)
    sort.Ints(arr)
    return reflect.DeepEqual(target, arr)
}

---

func canBeEqual(target, arr []int) bool {
    targetFreq := map[int]int{}
    arrFreq := map[int]int{}
    for i := 0; i < len(arr); i++ {
        targetFreq[target[i]]++
        arrFreq[arr[i]]++
    }
    return reflect.DeepEqual(targetFreq, arrFreq)
}
