func combine(n, k int) (result [][]int) {
    combineHelper(n, k, 1, 0, make([]int, k), &result)
    return result
}

func combineHelper(n, k, indx, depth int, combo []int, result *[][]int) {
    if depth == k {
        *result = append(*result, append([]int{}, combo...))
        return
    }
    for i := indx; i <= n; i++ {
        combo[depth] = i
        combineHelper(n, k, i + 1, depth + 1, combo, result)
    }
}
