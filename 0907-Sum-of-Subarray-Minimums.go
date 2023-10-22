func sumSubarrayMins(arr []int) int {
    mod := 7 + int(1e9)
    n := len(arr)
    var res int
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            res = (res + subarrayMin(arr, i, j)) % mod
        }
    }
    return res
}

func subarrayMin(arr []int, i, j int) int {
    res := arr[i]
    for ; i <= j; i++ {
        res = min(res, arr[i])
    }
    return res
}

func min(x, y int) int { if x < y { return x }; return y }
