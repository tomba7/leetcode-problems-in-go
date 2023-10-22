func arrayPairSum(nums []int) int {
    var maxSum int
    sort.Ints(nums)
    for i := 0; i < len(nums); i += 2 {
        maxSum += nums[i]
    }
    return maxSum
}

--

func arrayPairSum(nums []int) int {
    const maxRange = 10000
    var table [20001]int
    for _, n := range nums { table[n+maxRange]++ }
    var sum, first int
    for n := -maxRange; n <= maxRange; {
        if table[n+maxRange] > 0 && first == 0 {
            sum += n
            table[n+maxRange]--
            first = 1
        } else if table[n+maxRange] > 0 && first == 1 {
            table[n+maxRange]--
            first = 0
        } else {
            n++
        }
    }
    return sum
}
