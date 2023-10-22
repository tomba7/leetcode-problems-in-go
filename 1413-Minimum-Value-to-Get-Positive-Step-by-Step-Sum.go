func minStartValue(nums []int) int {
    sum := 0
    minVal := 0
    for i := 0; i < len(nums); i++ {
       sum += nums[i] 
       minVal = min(minVal, sum)
    }
    return abs(minVal-1)
}

func min(x, y int) int { if x < y { return x }; return y }
func abs(x int) int { if x < 0 { return -x }; return x }
