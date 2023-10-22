func maxProduct(nums []int) int {
    if len(nums) == 0 { return 0 }
    maxProd, minProd, result := nums[0], nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        maxProdTmp := max(nums[i], max(maxProd * nums[i], minProd * nums[i]))
        minProd = min(nums[i], min(maxProd * nums[i], minProd * nums[i]))
        maxProd = maxProdTmp
        result = max(result, maxProd)
    }
    return result
}

func max(x, y int) int { if x > y { return x }; return y }
func min(x, y int) int { if x < y { return x }; return y }
