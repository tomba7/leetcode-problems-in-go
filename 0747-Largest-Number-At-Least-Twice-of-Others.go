func dominantIndex(nums []int) int {
    var max, maxIndex int
    for i, n := range nums {
        if n > max {
            max, maxIndex = n, i
        }
    }
    for _, n := range nums {
        if n != max && 2*n > max {
            return -1
        }
    }
    return maxIndex
}


--------------------
ALTERNATIVE SOLUTION
--------------------
func dominantIndex(nums []int) int {
    var maxIndex int
    for i := range nums {
        if nums[i] > nums[maxIndex] { maxIndex = i }
    }
    for i := range nums {
        if maxIndex != i && nums[maxIndex] < 2 * nums[i] {
            return -1
        }
    }
    return maxIndex
}
