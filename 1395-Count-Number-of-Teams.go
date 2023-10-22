func numTeams(rating []int) int {
    var res int
    n := len(rating)
    for i := 0; i < n; i++ {
        var leftSmaller, leftGreater, rightSmaller, rightGreater int
        for j := 0; j < i; j++ {
            if rating[j] < rating[i] { leftSmaller++ }
            if rating[j] > rating[i] { leftGreater++ }
        }
        for k := i+1; k < n; k++ {
            if rating[k] < rating[i] { rightSmaller++ }
            if rating[k] > rating[i] { rightGreater++ }
        }
        res += leftSmaller * rightGreater + leftGreater * rightSmaller
    }
    return res
}

---

// Brute Force
func numTeams(rating []int) int {
    var res int
    n := len(rating)
    for i := 0; i < n-2; i++ {
        for j := i+1; j < n-1; j++ {
            for k := j+1; k < n; k++ {
                if rating[i] < rating[j] && rating[j] < rating[k] {
                    res++
                } else if rating[i] > rating[j] && rating[j] > rating[k] {
                    res++
                }
            }
        }
    }
    return res
}
