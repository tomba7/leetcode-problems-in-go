func pathSum(root *TreeNode, target int) int {
    var res int
    freq := map[int]int{0:1}
    pathSumHelper(root, 0, target, freq, &res)
    return res
}

func pathSumHelper(root *TreeNode, sum, target int, freq map[int]int, res *int) {
    if root == nil { return }
    sum += root.Val
    if count, contains := freq[sum-target]; contains {
        *res += count
    }
    freq[sum]++
    pathSumHelper(root.Left, sum, target, freq, res)
    pathSumHelper(root.Right, sum, target, freq, res)
    freq[sum]--
    if freq[sum] == 0 {
        delete(freq, sum)
    }
}

---

func pathSum(root *TreeNode, sum int) int {
    freq := map[int]int{0: 1}
    return pathSumHelper(root, sum, 0, freq)
}

func pathSumHelper(node *TreeNode, sum, total int, freq map[int]int) int {
    if node == nil { return 0 }
    total += node.Val
    result := freq[total - sum]
    freq[total]++
    result += pathSumHelper(node.Left, sum, total, freq) +
              pathSumHelper(node.Right, sum, total, freq)
    freq[total]--
    return result
}
