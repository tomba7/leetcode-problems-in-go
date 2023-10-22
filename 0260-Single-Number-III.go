func singleNumber(nums []int) []int {
    var bitmask int
    for _, num := range nums {
        bitmask ^= num
    }
    diff := bitmask & -bitmask
    var x int
    for _, num := range nums {
        if num & diff != 0 {
            x ^= num
        }
    }
    return []int{x, bitmask^x}
}
