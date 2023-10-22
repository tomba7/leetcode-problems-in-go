func findPairs(nums []int, k int) int {
    if len(nums) == 0 || k < 0 {return 0}
    numPairs := 0
    numMap := make(map[int]int)
    for _, num := range nums {
        numMap[num]++
    }
    for num, count := range numMap {
        if k == 0 {
            if count > 1 {
                numPairs++
            }
        } else {
            if _, ok := numMap[num + k]; ok {
                numPairs++
            }
        }
    }
    return numPairs
}
