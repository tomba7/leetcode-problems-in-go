func getPermutation(n int, k int) string {
    nums, result := make([]byte, n), make([]byte, n)
    for i := 0; i < n; i++ { nums[i] = byte(i) + 1 }
    k--
    for i := 1; i <= n; i++ {
        index := k/fact(n - i)
        result[i - 1] = '0' + nums[index]
        nums = append(nums[:index], nums[index + 1:]...)
        k -= index * fact(n - i)
    }
    return string(result)
}

func fact(n int) int { if n <= 1 { return 1 }; return n * fact(n - 1) }
