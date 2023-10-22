func findKthPositive(arr []int, k int) int {
    if k < arr[0] { return k }
    k -= arr[0] - 1
    for i := 0; i < len(arr)-1; i++ {
        missing := arr[i+1] - arr[i] - 1
        if k <= missing {
            return arr[i] + k
        }
        k -= missing
    }
    return arr[len(arr)-1] + k
}

---

/* Brute force
 [2,3,4,7,11]
 1,5,6,8,9,10,12....
 */
func findKthPositive(arr []int, k int) int {
    table := map[int]bool{}
    for _, num := range arr {
        table[num] = true
    }
    var res int
    for i := 0; i != k; {
        res++
        if table[res] { continue }
        i++
    }
    return res
}
