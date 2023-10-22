// Time: O(n), Space: O(n)
func productExceptSelf(nums []int) []int {
    result := make([]int, len(nums))
    product := 1
    for i := range nums {
        result[i] = product
        product *= nums[i]
    }
    product = 1
    for i := len(nums) - 1; i >= 0; i-- {
        result[i] *= product
        product *= nums[i]
    }
    return result
}

---
// Time: O(n), Space: O(1)
func productExceptSelf(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    for i := range res {
        res[i] = 1
    }
    leftProd, rightProd := 1, 1
    for i := 0; i < n; i++ {
        j := n-1-i
        res[i] *= leftProd
        res[j] *= rightProd
        leftProd *= nums[i]
        rightProd *= nums[j]
    }
    return res
}
