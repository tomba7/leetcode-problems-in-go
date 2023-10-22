func numTrees(n int) int {
    if n == 0 || n == 1 { return 1 }
    var total int
    for left := 0; left < n; left++ {
        right := n - 1 - left
        total += numTrees(left) * numTrees(right)
    }
    return total;
}

---

func numTrees(n int) int {
    table := make([]int, n+1)
    table[0] = 1
    table[1] = 1
    for numNodes := 2; numNodes <= n; numNodes++ {
        for left := 0; left < numNodes; left++ {
            right := numNodes - 1 - left
            table[numNodes] += table[left] * table[right]
        }
    }
    return table[n]
}
